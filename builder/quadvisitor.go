package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/semantic"
	"fmt"
)

// ExpressionResult representa el resultado de evaluar una expresión
type ExpressionResult struct {
	Address interface{}
	Type    string
}

// QuadrupleVisitor maneja toda la lógica de generación de cuáruplos
// ENFOQUE LIMPIO: Sin pilas, usando visitor directo con ANTLR
type QuadrupleVisitor struct {
	builder   *DirectoryBuilder
	quadQueue *quads.QuadrupleQueue

	// Para manejo de condicionales
	pendingJumps []int

	// Para manejo de funciones
	currentFunction string
	functionStarts  map[string]int

	// Manejar parametros y llamadas
	parameterStack []interface{}
	returnStack    []int

	// Debug
	debugExpressions bool
}

// NewQuadrupleVisitor crea un nuevo visitor para cuádruplos
func NewQuadrupleVisitor(builder *DirectoryBuilder) *QuadrupleVisitor {
	return &QuadrupleVisitor{
		builder:          builder,
		quadQueue:        quads.NewQuadrupleQueue(),
		pendingJumps:     []int{},
		functionStarts:   make(map[string]int),
		currentFunction:  "program",
		parameterStack:   make([]interface{}, 0),
		returnStack:      make([]int, 0),
		debugExpressions: true,
	}
}

// ========== MÉTODOS DE DEBUG ==========

func (qv *QuadrupleVisitor) debugLog(msg string, args ...interface{}) {
	if qv.debugExpressions {
		fmt.Printf("[EXPR DEBUG] "+msg+"\n", args...)
	}
}

// ========== EXPRESIONES - ENFOQUE VISITOR LIMPIO ==========

func (qv *QuadrupleVisitor) VisitExpression(ctx grammar.IExpressionContext) *ExpressionResult {
	qv.debugLog("=== VisitExpression: %s ===", ctx.GetText())
	return qv.VisitRelExpr(ctx.Rel_expr())
}

func (qv *QuadrupleVisitor) VisitRelExpr(ctx grammar.IRel_exprContext) *ExpressionResult {
	qv.debugLog("=== VisitRelExpr: %s ===", ctx.GetText())

	// Visitar primera expresión de suma
	leftResult := qv.VisitAddExpr(ctx.Add_expr(0))

	// Si hay operador relacional
	if ctx.Relop() != nil && len(ctx.AllAdd_expr()) > 1 {
		operator := ctx.Relop().GetText()
		rightResult := qv.VisitAddExpr(ctx.Add_expr(1))

		qv.debugLog("Processing relational: %v %s %v", leftResult.Address, operator, rightResult.Address)
		return qv.generateBinaryOperation(operator, leftResult, rightResult)
	}

	return leftResult
}

func (qv *QuadrupleVisitor) VisitAddExpr(ctx grammar.IAdd_exprContext) *ExpressionResult {
	qv.debugLog("=== VisitAddExpr: %s ===", ctx.GetText())

	// Visitar primer término
	result := qv.VisitTerm(ctx.Term(0))

	// Procesar cada operación de suma/resta de izquierda a derecha
	terms := ctx.AllTerm()
	addops := ctx.AllAddop()

	for i := 0; i < len(addops); i++ {
		operator := addops[i].GetText()
		rightTerm := qv.VisitTerm(terms[i+1])

		qv.debugLog("Processing add/sub: %v %s %v", result.Address, operator, rightTerm.Address)
		result = qv.generateBinaryOperation(operator, result, rightTerm)
	}

	return result
}

func (qv *QuadrupleVisitor) VisitTerm(ctx grammar.ITermContext) *ExpressionResult {
	qv.debugLog("=== VisitTerm: %s ===", ctx.GetText())

	// Visitar primer factor
	result := qv.VisitFactor(ctx.Factor(0))

	// Procesar cada operación de multiplicación/división de izquierda a derecha
	factors := ctx.AllFactor()
	mulops := ctx.AllMulop()

	for i := 0; i < len(mulops); i++ {
		operator := mulops[i].GetText()
		rightFactor := qv.VisitFactor(factors[i+1])

		qv.debugLog("Processing mul/div: %v %s %v", result.Address, operator, rightFactor.Address)
		result = qv.generateBinaryOperation(operator, result, rightFactor)
	}

	return result
}

func (qv *QuadrupleVisitor) VisitFactor(ctx grammar.IFactorContext) *ExpressionResult {
	qv.debugLog("=== VisitFactor: %s ===", ctx.GetText())

	// Manejar paréntesis
	if ctx.Expression() != nil {
		return qv.VisitExpression(ctx.Expression())
	}

	// Manejar signo unario
	if ctx.Addop() != nil {
		valueResult := qv.VisitValue(ctx.Value())
		sign := ctx.Addop().GetText()

		if sign == "-" {
			// Generar cuádruple para negación: -1 * valor
			tempAddress := qv.getTemporaryAddress(valueResult.Type)
			constNegOne := qv.builder.ConstantTable.AddConstant("-1", valueResult.Type)
			qv.quadQueue.Add("*", constNegOne, valueResult.Address, tempAddress)
			return &ExpressionResult{Address: tempAddress, Type: valueResult.Type}
		}
	}

	// Visitar valor normal
	return qv.VisitValue(ctx.Value())
}

func (qv *QuadrupleVisitor) VisitValue(ctx grammar.IValueContext) *ExpressionResult {
	qv.debugLog("=== VisitValue: %s ===", ctx.GetText())

	if ctx.ID() != nil {
		// Variable
		varName := ctx.ID().GetText()
		variable, found := qv.builder.Directory.FindVariableDeep(varName)
		if !found {
			qv.builder.addError(fmt.Sprintf("Variable '%s' no definida", varName))
			return &ExpressionResult{Address: 0, Type: "error"}
		}

		var address interface{}
		if variable.MemoryAddress == 0 {
			address = varName // Usar string cuando no hay dirección de memoria asignada
		} else {
			address = variable.MemoryAddress // Usar int cuando sí hay dirección
		}

		qv.debugLog("Variable %s -> address: %v, type: %s", varName, address, variable.Type)
		return &ExpressionResult{Address: address, Type: variable.Type}
	}

	if ctx.Constant() != nil {
		// Constante
		constValue := ctx.Constant().GetText()
		constType := semantic.InferTypeFromConstant(constValue)
		address := qv.builder.ConstantTable.AddConstant(constValue, constType)

		qv.debugLog("Constant %s -> address: %v, type: %s", constValue, address, constType)
		return &ExpressionResult{Address: address, Type: constType}
	}

	return &ExpressionResult{Address: 0, Type: "error"}
}

// ========== MÉTODO AUXILIAR PARA GENERAR OPERACIONES BINARIAS ==========

func (qv *QuadrupleVisitor) generateBinaryOperation(operator string, left, right *ExpressionResult) *ExpressionResult {
	qv.debugLog("Generating binary operation: %v %s %v", left.Address, operator, right.Address)

	// Verificar compatibilidad de tipos
	resultType, ok := semantic.Cube.GetResultType(left.Type, right.Type, operator)
	if !ok || resultType == "error" {
		qv.builder.addError(fmt.Sprintf("Tipo incompatible: %s %s %s", left.Type, operator, right.Type))
		return &ExpressionResult{Address: 0, Type: "error"}
	}

	// Generar dirección temporal
	tempAddress := qv.getTemporaryAddress(resultType)

	// Generar cuádruple
	quadIndex := qv.quadQueue.Add(operator, left.Address, right.Address, tempAddress)
	qv.debugLog("Generated quadruple %d: (%s, %v, %v, %v)", quadIndex, operator, left.Address, right.Address, tempAddress)

	return &ExpressionResult{Address: tempAddress, Type: resultType}
}

func (qv *QuadrupleVisitor) getTemporaryAddress(resultType string) interface{} {
	tempVar := qv.builder.Directory.NewTempVar(resultType)
	return tempVar.MemoryAddress
}

// ========== DECLARACIONES Y ASIGNACIONES ==========

func (qv *QuadrupleVisitor) VisitAssignment(ctx *grammar.AssignContext) {
	variableName := ctx.ID().GetText()
	qv.debugLog("=== Assignment: %s ===", variableName)

	// Validar existencia de variable
	if err := qv.builder.Directory.ValidateVariable(variableName); err != nil {
		qv.builder.addError(err.Error())
		return
	}

	// Evaluar expresión usando visitor directo
	exprResult := qv.VisitExpression(ctx.Expression())
	if exprResult.Type == "error" {
		return
	}

	// Buscar variable destino
	variable, exists := qv.builder.Directory.FindVariableDeep(variableName)
	if !exists {
		qv.builder.addError(fmt.Sprintf("Variable '%s' no encontrada", variableName))
		return
	}

	// Validar compatibilidad de tipos
	if exprResult.Type != variable.Type {
		qv.builder.addError(fmt.Sprintf("Tipo incompatible en asignación: %s = %s", variable.Type, exprResult.Type))
		return
	}

	// Generar cuádruple
	address := variable.MemoryAddress
	if address == 0 {
		qv.quadQueue.Add("=", exprResult.Address, "", variableName)
	} else {
		qv.quadQueue.Add("=", exprResult.Address, "", address)
	}

	qv.debugLog("Assignment complete: %s = %v", variableName, exprResult.Address)
}

// ========== DECLARACIONES PRINT ==========

func (qv *QuadrupleVisitor) VisitPrintStatement(ctx *grammar.Print_stmtContext) {
	qv.debugLog("=== Print Statement ===")

	for _, pexpr := range ctx.AllPrintexpr() {
		switch {
		case pexpr.STRING() != nil:
			stringValue := pexpr.STRING().GetText()
			qv.quadQueue.Add("print", stringValue, "", "")
			qv.debugLog("Print string: %s", stringValue)

		case pexpr.Expression() != nil:
			// Usar visitor directo
			exprResult := qv.VisitExpression(pexpr.Expression())
			if exprResult.Type != "error" {
				qv.quadQueue.Add("print", exprResult.Address, "", "")
				qv.debugLog("Print expression: %v", exprResult.Address)
			}
		}
	}
}

// ========== CONDICIONALES ==========

func (qv *QuadrupleVisitor) VisitCondition(ctx *grammar.ConditionContext) {
	qv.debugLog("=== Condition ===")

	condResult := qv.VisitExpression(ctx.Expression())
	if condResult.Type == "error" {
		return
	}

	if condResult.Type != "bool" {
		qv.builder.addError(fmt.Sprintf("Condición no booleana: %v", condResult.Type))
		return
	}

	gotoFIndex := qv.quadQueue.Add("GOTOF", condResult.Address, nil, nil)
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
	startQuad := qv.quadQueue.GenerateFUNC(functionName, "void")
	qv.functionStarts[functionName] = startQuad
	qv.currentFunction = functionName

	qv.builder.Directory.SetFunctionQuadruples(functionName, startQuad, -1)
	qv.builder.debugLog("Función '%s' inicia en cuádruple %d", functionName, startQuad)
}

func (qv *QuadrupleVisitor) ExitFunction(functionName string) {
	qv.quadQueue.GenerateENDFUNC()

	endQuad := qv.quadQueue.Size() - 1
	startQuad := qv.functionStarts[functionName]

	qv.builder.Directory.SetFunctionQuadruples(functionName, startQuad, endQuad)

	localCount := qv.builder.Directory.CountLocalVariables(functionName)
	tempCount := qv.builder.Directory.CountTempVariables(functionName)

	qv.builder.debugLog("Función '%s' termina en cuádruple %d (locales: %d, temps: %d)",
		functionName, endQuad, localCount, tempCount)

	qv.currentFunction = "program"
}

func (qv *QuadrupleVisitor) HandleParameterDeclaration(paramType, paramName string, address int) {
	qv.quadQueue.GeneratePARAM(paramType, paramName, address)

	if qv.builder != nil && qv.builder.Debug {
		fmt.Printf("[QuadVisitor] PARAM: %s %s -> %d\n", paramType, paramName, address)
	}
}

func (qv *QuadrupleVisitor) VisitFunctionCall(ctx *grammar.F_callContext) {
	funcName := ctx.ID().GetText()
	numArgs := qv.countArguments(ctx)

	if err := qv.builder.Directory.ValidateFunctionCall(funcName, numArgs); err != nil {
		qv.builder.addError(err.Error())
		return
	}

	// Recopilar argumentos evaluando expresiones directamente
	paramValues := make([]interface{}, numArgs)
	if ctx.Arg_list() != nil {
		argListCtx := ctx.Arg_list().(*grammar.Arg_listContext)
		for i, argExpr := range argListCtx.AllExpression() {
			argResult := qv.VisitExpression(argExpr)
			if argResult.Type == "error" {
				qv.builder.addError(fmt.Sprintf("Error en argumento %d de función %s", i+1, funcName))
				return
			}
			paramValues[i] = argResult.Address
		}
	}

	// Generar cuádruplos para llamada a función
	funcInfo, err := qv.builder.Directory.GetFunctionInfo(funcName)
	if err != nil {
		qv.builder.addError(err.Error())
		return
	}

	// 1. ERA - Crear espacio de registro de activación
	localVars := qv.builder.Directory.CountLocalVariables(funcName)
	tempVars := qv.builder.Directory.CountTempVariables(funcName)
	totalSize := localVars + tempVars + numArgs
	qv.quadQueue.GenerateERA(funcName, totalSize)

	// 2. PARAMETER - Pasar cada parámetro
	for i, paramValue := range paramValues {
		qv.quadQueue.GeneratePARAMETER(paramValue, i)
	}

	// 3. GOSUB - Llamar a la función
	startAddress := funcInfo.StartQuadruple
	qv.quadQueue.GenerateGOSUB(funcName, startAddress)

	qv.builder.debugLog("Generada llamada completa a función: %s", funcName)
}

func (qv *QuadrupleVisitor) VisitReturnStatement(returnValue interface{}) {
	if returnValue != nil {
		qv.quadQueue.GenerateRETURN(returnValue)
	} else {
		qv.quadQueue.GenerateRETURN(nil)
	}
}

// ========== MÉTODOS AUXILIARES ==========

func (qv *QuadrupleVisitor) hasElseClause(condCtx *grammar.ConditionContext) bool {
	if condCtx.Else_part() == nil {
		return false
	}

	elsePartCtx := condCtx.Else_part().(*grammar.Else_partContext)
	return elsePartCtx.Body() != nil
}

func (qv *QuadrupleVisitor) handleBodyWithElse() {
	elseStart := qv.quadQueue.Size()
	gotoEndIndex := qv.quadQueue.Add("EEE", "", "", "")

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
