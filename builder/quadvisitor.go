package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/semantic"
	"fmt"
)

// QuadrupleVisitor maneja toda la lógica de generación de cuádruplos
type QuadrupleVisitor struct {
	builder       *DirectoryBuilder
	operatorStack *quads.OperatorStack
	operandStack  *quads.Stack
	typeStack     *quads.TypeStack
	quadQueue     *quads.QuadrupleQueue
	pendingJumps  []int

	// Para manejo de funciones
	currentFunction string
	functionStarts  map[string]int // Mapeo función -> cuádruplo de inicio
}

// NewQuadrupleVisitor crea un nuevo visitor para cuádruplos
func NewQuadrupleVisitor(builder *DirectoryBuilder) *QuadrupleVisitor {
	return &QuadrupleVisitor{
		builder:         builder,
		operatorStack:   quads.NewOperatorStack(),
		operandStack:    quads.NewOperandStack(),
		typeStack:       quads.NewTypeStack(),
		quadQueue:       quads.NewQuadrupleQueue(),
		pendingJumps:    []int{},
		functionStarts:  make(map[string]int),
		currentFunction: "program",
	}
}

// ========== ASIGNACIONES ==========

func (qv *QuadrupleVisitor) VisitAssignment(ctx *grammar.AssignContext) {
	variableName := ctx.ID().GetText()

	// Validar existencia de variable
	if err := qv.builder.Directory.ValidateVariable(variableName); err != nil {
		qv.builder.addError(err.Error())
		return
	}

	// Obtener operandos de la expresión
	rightOperand, rightType, ok := qv.popOperandAndType()
	if !ok {
		qv.builder.addError("Error al obtener el valor de la expresión en la asignación")
		return
	}

	// Buscar variable destino
	variable, exists := qv.builder.Directory.FindVariableDeep(variableName)
	if !exists {
		qv.builder.addError(fmt.Sprintf("Variable '%s' no encontrada", variableName))
		return
	}

	// Validar compatibilidad de tipos
	if rightType != variable.Type {
		qv.builder.addError(fmt.Sprintf("Tipo incompatible en asignación: %s = %s", variable.Type, rightType))
		return
	}

	// Generar cuádruplo
	address := variable.MemoryAddress
	if address == 0 {
		qv.quadQueue.Add("=", rightOperand, "", variableName)
	} else {
		qv.quadQueue.Add("=", rightOperand, "", address)
	}
}

// ========== DECLARACIONES PRINT ==========

func (qv *QuadrupleVisitor) VisitPrintStatement(ctx *grammar.Print_stmtContext) {
	for _, pexpr := range ctx.AllPrintexpr() {
		switch {
		case pexpr.STRING() != nil:
			stringValue := pexpr.STRING().GetText()
			qv.quadQueue.Add("print", stringValue, "", "")

		case pexpr.Expression() != nil:
			qv.builder.validateExpression(pexpr.Expression())
			value, _ := qv.operandStack.Pop()
			qv.typeStack.Pop()
			qv.quadQueue.Add("print", value, "", "")
		}
	}
}

// ========== CONDICIONALES ==========

func (qv *QuadrupleVisitor) VisitCondition(ctx *grammar.ConditionContext) {
	qv.builder.validateExpression(ctx.Expression())

	condResult, condType, ok := qv.popOperandAndType()
	if !ok {
		qv.builder.addError("Condición sin resultado")
		return
	}

	if condType != "bool" {
		qv.builder.addError(fmt.Sprintf("Condición no booleana: %v", condType))
		return
	}

	gotoFIndex := qv.quadQueue.Add("GOTOF", condResult, nil, nil)
	qv.pendingJumps = append(qv.pendingJumps, gotoFIndex)
}

func (qv *QuadrupleVisitor) VisitBody(ctx *grammar.BodyContext) {
	parentCtx := ctx.GetParent()
	condCtx, ok := parentCtx.(*grammar.ConditionContext)
	if !ok {
		return
	}

	hasElse := qv.hasElseClause(condCtx)

	if hasElse {
		qv.handleBodyWithElse()
	} else {
		qv.handleBodyWithoutElse()
	}
}

func (qv *QuadrupleVisitor) VisitElsePart(ctx *grammar.Else_partContext) {
	if len(qv.pendingJumps) == 0 {
		return
	}

	gotoEndIndex := qv.pendingJumps[len(qv.pendingJumps)-1]
	qv.pendingJumps = qv.pendingJumps[:len(qv.pendingJumps)-1]

	end := qv.quadQueue.Size()
	qv.quadQueue.FillJump(gotoEndIndex, end)
}

// ========== MANEJO DE FUNCIONES ==========

func (qv *QuadrupleVisitor) EnterFunction(functionName string) {
	startQuad := qv.quadQueue.Size()
	qv.functionStarts[functionName] = startQuad
	qv.currentFunction = functionName

	qv.builder.Directory.SetFunctionQuadruples(functionName, startQuad, -1)
	qv.builder.debugLog("Función '%s' inicia en cuádruplo %d", functionName, startQuad)
}

func (qv *QuadrupleVisitor) ExitFunction(functionName string) {
	qv.quadQueue.Add("ENDFUNC", "", "", "")

	endQuad := qv.quadQueue.Size() - 1
	startQuad := qv.functionStarts[functionName]

	qv.builder.Directory.SetFunctionQuadruples(functionName, startQuad, endQuad)

	localCount := qv.builder.Directory.CountLocalVariables(functionName)
	tempCount := qv.builder.Directory.CountTempVariables(functionName)

	qv.builder.debugLog("Función '%s' termina en cuádruplo %d (locales: %d, temps: %d)",
		functionName, endQuad, localCount, tempCount)

	qv.currentFunction = "program"
}

func (qv *QuadrupleVisitor) VisitFunctionCall(ctx *grammar.F_callContext) {
	funcName := ctx.ID().GetText()
	numArgs := qv.countArguments(ctx)

	if err := qv.builder.Directory.ValidateFunctionCall(funcName, numArgs); err != nil {
		qv.builder.addError(err.Error())
		return
	}

	paramValues := make([]interface{}, numArgs)
	for i := numArgs - 1; i >= 0; i-- {
		argValue, ok := qv.operandStack.Pop()
		if !ok {
			qv.builder.addError(fmt.Sprintf("Error: argumento %d no encontrado para función %s", i+1, funcName))
			return
		}
		qv.typeStack.Pop()
		paramValues[i] = argValue
	}

	for i, paramValue := range paramValues {
		qv.quadQueue.Add("PARAM", paramValue, "", i)
	}

	funcInfo, err := qv.builder.Directory.GetFunctionInfo(funcName)
	if err != nil {
		qv.builder.addError(err.Error())
		return
	}

	localVars := qv.builder.Directory.CountLocalVariables(funcName)
	tempVars := qv.builder.Directory.CountTempVariables(funcName)
	qv.quadQueue.Add("ERA", funcName, localVars, tempVars)

	startAddress := funcInfo.StartQuadruple
	returnAddress := qv.quadQueue.Size() + 2
	qv.quadQueue.Add("GOSUB", funcName, startAddress, returnAddress)

	qv.builder.debugLog("Generada llamada completa a función: %s", funcName)
}

// ========== EXPRESIONES ARITMÉTICAS ==========

func (qv *QuadrupleVisitor) VisitFactor(ctx *grammar.FactorContext) {
	if ctx.Expression() != nil {
		return
	}

	valCtx := ctx.Value()
	if valCtx == nil {
		qv.builder.addError("Error: factor sin valor")
		return
	}

	if ctx.Addop() != nil {
		qv.handleUnaryOperator(ctx.Addop().GetText())
	}
}

func (qv *QuadrupleVisitor) VisitTerm(ctx *grammar.TermContext) {
	mulops := ctx.AllMulop()
	for range mulops {
		qv.processBinaryOperation("mulop")
	}
}

func (qv *QuadrupleVisitor) VisitAddExpression(ctx *grammar.Add_exprContext) {
	addops := ctx.AllAddop()
	for range addops {
		qv.processBinaryOperation("addop")
	}
}

func (qv *QuadrupleVisitor) VisitRelationalExpression(ctx *grammar.Rel_exprContext) {
	if ctx.Relop() == nil {
		return
	}
	qv.processBinaryOperation("relop", ctx.Relop().GetText())
}

func (qv *QuadrupleVisitor) VisitValue(ctx *grammar.ValueContext) {
	switch {
	case ctx.ID() != nil:
		qv.handleVariableValue(ctx.ID().GetText())
	case ctx.Constant() != nil:
		qv.handleConstantValue(ctx.Constant().GetText())
	}
}

// ========== OPERADORES ==========

func (qv *QuadrupleVisitor) EnterMulop(ctx *grammar.MulopContext) {
	qv.operatorStack.Push(ctx.GetText())
}

func (qv *QuadrupleVisitor) EnterAddop(ctx *grammar.AddopContext) {
	qv.operatorStack.Push(ctx.GetText())
}

// ========== MÉTODOS AUXILIARES ==========

func (qv *QuadrupleVisitor) popOperandAndType() (interface{}, string, bool) {
	operand, ok1 := qv.operandStack.Pop()
	typ, ok2 := qv.typeStack.Pop()
	return operand, typ, ok1 && ok2
}

func (qv *QuadrupleVisitor) getTemporaryAddress(resultType string) interface{} {
	tempVar := qv.builder.Directory.NewTempVar(resultType)
	return tempVar.MemoryAddress
}

func (qv *QuadrupleVisitor) processBinaryOperation(opType string, specificOp ...string) {
	var op string
	if len(specificOp) > 0 {
		op = specificOp[0]
	} else {
		var ok bool
		op, ok = qv.operatorStack.Pop()
		if !ok {
			qv.builder.addError(fmt.Sprintf("Error: operador faltante en pila (%s)", opType))
			return
		}
	}

	right, rightType, ok1 := qv.popOperandAndType()
	left, leftType, ok2 := qv.popOperandAndType()
	if !ok1 || !ok2 {
		qv.builder.addError(fmt.Sprintf("Error: operandos insuficientes para operación %s", opType))
		return
	}

	resultType, ok := semantic.Cube.GetResultType(leftType, rightType, op)
	if !ok || resultType == "error" {
		qv.builder.addError(fmt.Sprintf("Tipo incompatible: %s %s %s", leftType, op, rightType))
		return
	}

	tempAddress := qv.getTemporaryAddress(resultType)
	qv.quadQueue.Add(op, left, right, tempAddress)

	qv.operandStack.Push(tempAddress)
	qv.typeStack.Push(resultType)
}

func (qv *QuadrupleVisitor) handleUnaryOperator(sign string) {
	val, ok := qv.operandStack.Pop()
	if !ok {
		qv.builder.addError("Error: pila de operandos vacía al aplicar signo")
		return
	}
	qv.operandStack.Push(sign + fmt.Sprint(val))
}

func (qv *QuadrupleVisitor) handleVariableValue(name string) {
	variable, found := qv.builder.Directory.FindVariableDeep(name)
	if !found {
		qv.builder.addError(fmt.Sprintf("Variable '%s' no definida", name))
		return
	}

	address := variable.MemoryAddress
	if address == 0 {
		qv.operandStack.Push(name)
	} else {
		qv.operandStack.Push(address)
	}

	qv.typeStack.Push(variable.Type)
}

func (qv *QuadrupleVisitor) handleConstantValue(val string) {
	typ := semantic.InferTypeFromConstant(val)
	address := qv.builder.ConstantTable.AddConstant(val, typ)

	qv.operandStack.Push(address)
	qv.typeStack.Push(typ)
}

func (qv *QuadrupleVisitor) hasElseClause(condCtx *grammar.ConditionContext) bool {
	if condCtx.Else_part() == nil {
		return false
	}

	elsePartCtx := condCtx.Else_part().(*grammar.Else_partContext)
	return elsePartCtx.Body() != nil
}

func (qv *QuadrupleVisitor) handleBodyWithElse() {
	elseStart := qv.quadQueue.Size()
	gotoEndIndex := qv.quadQueue.Add("GOTO", "", "", "")

	if len(qv.pendingJumps) > 0 {
		lastGotoFIndex := qv.pendingJumps[len(qv.pendingJumps)-1]
		qv.quadQueue.UpdateResult(lastGotoFIndex, elseStart)
		qv.pendingJumps = qv.pendingJumps[:len(qv.pendingJumps)-1]
		qv.pendingJumps = append(qv.pendingJumps, gotoEndIndex)
	}
}

func (qv *QuadrupleVisitor) handleBodyWithoutElse() {
	if len(qv.pendingJumps) > 0 {
		lastGotoFIndex := qv.pendingJumps[len(qv.pendingJumps)-1]
		end := qv.quadQueue.Size()
		qv.quadQueue.UpdateResult(lastGotoFIndex, end)
		qv.pendingJumps = qv.pendingJumps[:len(qv.pendingJumps)-1]
	}
}

func (qv *QuadrupleVisitor) countArguments(ctx *grammar.F_callContext) int {
	if ctx.Arg_list() == nil {
		return 0
	}

	argListCtx := ctx.Arg_list().(*grammar.Arg_listContext)
	return len(argListCtx.AllExpression())
}

func (qv *QuadrupleVisitor) PrintQuadruples() {
	fmt.Println("=== Cuádruplos generados ===")
	for i, quad := range qv.quadQueue.GetAll() {
		fmt.Printf("%d: (%s, %v, %v, %v)\n", i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result)
	}
}

func (qv *QuadrupleVisitor) GetQuadruples() []quads.Quadruple {
	if qv.quadQueue == nil {
		return []quads.Quadruple{}
	}
	return qv.quadQueue.GetAll()
}
