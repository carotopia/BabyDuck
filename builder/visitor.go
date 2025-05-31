package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/semantic"
	"BabyDuckCompiler/symbols"
	"fmt"
)

type PureQuadrupleVisitor struct {
	Directory     *symbols.FunctionDirectory
	ConstantTable *symbols.ConstantTable
	quadQueue     *quads.QuadrupleQueue
	Errors        []string
	Debug         bool

	jumpStack []int

	// Para funciones
	currentFunction string
	functionStarts  map[string]int

	// Debug
	debugExpressions bool
}

type ExpressionResult struct {
	Address interface{}
	Type    string
}

// NewPureQuadrupleVisitor crea un visitor que SOLO maneja cuádruplos
func NewPureQuadrupleVisitor(funcDir *symbols.FunctionDirectory, constTable *symbols.ConstantTable, debug bool) *PureQuadrupleVisitor {
	v := &PureQuadrupleVisitor{
		Directory:        funcDir,
		ConstantTable:    constTable,
		quadQueue:        quads.NewQuadrupleQueue(),
		Errors:           []string{},
		Debug:            debug,
		jumpStack:        []int{},
		functionStarts:   make(map[string]int),
		currentFunction:  "program",
		debugExpressions: debug,
	}

	// Generar cuádruple inicial GOTO main
	mainGotoIndex := v.quadQueue.Add("GOTO", "", "", nil)
	v.jumpStack = append(v.jumpStack, mainGotoIndex)
	v.debugLog("Generated initial GOTO to main")

	return v
}

// ========== MÉTODO PRINCIPAL ==========

func (v *PureQuadrupleVisitor) VisitProgram(ctx *grammar.ProgramContext) {
	v.debugLog("=== Starting Program Quadruple Generation ===")

	// Procesar funciones - CORRECCIÓN: usar Func_() con guión bajo
	if ctx.AllFuncs() != nil {
		for _, funcsCtx := range ctx.AllFuncs() {
			// Hacer casting a FuncsContext primero, luego acceder a Func_()
			if funcsContext, ok := funcsCtx.(*grammar.FuncsContext); ok && funcsContext != nil {
				if funcCtx, ok := funcsContext.Func_().(*grammar.FuncContext); ok && funcCtx != nil {
					v.VisitFunction(funcCtx)
				}
			}
		}
	}

	if len(v.jumpStack) > 0 {
		mainStart := v.quadQueue.Size()
		mainGotoIndex := v.jumpStack[0]
		v.quadQueue.FillJump(mainGotoIndex, mainStart)
		v.jumpStack = v.jumpStack[1:]
		v.debugLog("Main starts at quadruple: %d", mainStart)
	}

	if bodyCtx, ok := ctx.Body().(*grammar.BodyContext); ok && bodyCtx != nil {
		v.VisitBody(bodyCtx)
	}

	v.debugLog("=== Program Quadruple Generation Complete ===")
}

// ========== CUERPO DE CÓDIGO ==========

func (v *PureQuadrupleVisitor) VisitBody(ctx *grammar.BodyContext) {
	v.debugLog("=== Processing Body ===")

	if ctx.AllStatement() != nil {
		for _, stmt := range ctx.AllStatement() {

			if stmtCtx, ok := stmt.(*grammar.StatementContext); ok && stmtCtx != nil {
				v.VisitStatement(stmtCtx)
			}
		}
	}
}

func (v *PureQuadrupleVisitor) VisitStatement(ctx *grammar.StatementContext) {
	v.debugLog("Processing Statement: %s", ctx.GetText())

	switch {
	case ctx.Assign() != nil:
		if assignCtx, ok := ctx.Assign().(*grammar.AssignContext); ok {
			v.VisitAssignment(assignCtx)
		}
	case ctx.Print_stmt() != nil:
		if printCtx, ok := ctx.Print_stmt().(*grammar.Print_stmtContext); ok {
			v.VisitPrintStatement(printCtx)
		}
	case ctx.Condition() != nil:
		if condCtx, ok := ctx.Condition().(*grammar.ConditionContext); ok {
			v.VisitCondition(condCtx)
		}
	case ctx.Cycle() != nil:
		if cycleCtx, ok := ctx.Cycle().(*grammar.CycleContext); ok {
			v.VisitCycle(cycleCtx) // ← AQUÍ SE PROCESA EL WHILE
		}
	case ctx.F_call() != nil:
		if callCtx, ok := ctx.F_call().(*grammar.F_callContext); ok {
			v.VisitFunctionCall(callCtx)
		}
	}
}

func (v *PureQuadrupleVisitor) VisitCycle(ctx *grammar.CycleContext) {
	v.debugLog("=== Processing While Loop ===")

	// 1. Marcar inicio del loop (donde estará la condición)
	loopStartPosition := v.quadQueue.Size()
	v.debugLog("Loop start position: %d", loopStartPosition)

	// 2. Evaluar condición del while - CORRECCIÓN: casting
	if exprCtx, ok := ctx.Expression().(*grammar.ExpressionContext); ok && exprCtx != nil {
		conditionResult := v.VisitExpression(exprCtx)
		if conditionResult.Type == "error" {
			v.addError("Error en condición del while")
			return
		}

		// Verificar que la condición sea booleana
		if conditionResult.Type != "bool" {
			v.addError(fmt.Sprintf("Condición de while no booleana: %v", conditionResult.Type))
			return
		}

		// 3. Generar GOTOF
		conditionalJumpPosition := v.quadQueue.Add("GOTOF", conditionResult.Address, "", nil)
		v.debugLog("Generated GOTOF at position: %d", conditionalJumpPosition)

		// 4. Procesar el body del while - CORRECCIÓN: casting
		if bodyCtx, ok := ctx.Body().(*grammar.BodyContext); ok && bodyCtx != nil {
			v.VisitBody(bodyCtx)
		}

		// 5. Generar GOTO de regreso al inicio del loop
		v.quadQueue.Add("GOTO", "", "", loopStartPosition)
		v.debugLog("Generated GOTO back to position: %d", loopStartPosition)

		// 6. Llenar el GOTOF con la posición de salida
		exitPosition := v.quadQueue.Size()
		v.quadQueue.FillJump(conditionalJumpPosition, exitPosition)
		v.debugLog("Filled GOTOF to exit at position: %d", exitPosition)
	}

	v.debugLog("=== While Loop Complete ===")
}

// ========== EXPRESIONES ==========

func (v *PureQuadrupleVisitor) VisitExpression(ctx *grammar.ExpressionContext) *ExpressionResult {
	v.debugLog("VisitExpression: %s", ctx.GetText())

	if relExprCtx, ok := ctx.Rel_expr().(*grammar.Rel_exprContext); ok && relExprCtx != nil {
		return v.VisitRelExpr(relExprCtx)
	}

	return &ExpressionResult{Address: 0, Type: "error"}
}

func (v *PureQuadrupleVisitor) VisitRelExpr(ctx *grammar.Rel_exprContext) *ExpressionResult {
	v.debugLog("VisitRelExpr: %s", ctx.GetText())

	var leftResult *ExpressionResult
	if addExprCtx, ok := ctx.Add_expr(0).(*grammar.Add_exprContext); ok && addExprCtx != nil {
		leftResult = v.VisitAddExpr(addExprCtx)
	}

	if ctx.Relop() != nil && len(ctx.AllAdd_expr()) > 1 {
		operator := ctx.Relop().GetText()

		var rightResult *ExpressionResult
		if addExprCtx, ok := ctx.Add_expr(1).(*grammar.Add_exprContext); ok && addExprCtx != nil {
			rightResult = v.VisitAddExpr(addExprCtx)
		}

		if leftResult != nil && rightResult != nil {
			v.debugLog("Processing relational: %v %s %v", leftResult.Address, operator, rightResult.Address)
			return v.generateBinaryOperation(operator, leftResult, rightResult)
		}
	}

	return leftResult
}

func (v *PureQuadrupleVisitor) VisitAddExpr(ctx *grammar.Add_exprContext) *ExpressionResult {
	v.debugLog("VisitAddExpr: %s", ctx.GetText())

	// CORRECCIÓN: casting para term
	var result *ExpressionResult
	if termCtx, ok := ctx.Term(0).(*grammar.TermContext); ok && termCtx != nil {
		result = v.VisitTerm(termCtx)
	}

	terms := ctx.AllTerm()
	addops := ctx.AllAddop()

	for i := 0; i < len(addops); i++ {
		operator := addops[i].GetText()

		var rightTerm *ExpressionResult
		if termCtx, ok := terms[i+1].(*grammar.TermContext); ok && termCtx != nil {
			rightTerm = v.VisitTerm(termCtx)
		}

		if result != nil && rightTerm != nil {
			v.debugLog("Processing add/sub: %v %s %v", result.Address, operator, rightTerm.Address)
			result = v.generateBinaryOperation(operator, result, rightTerm)
		}
	}

	return result
}

func (v *PureQuadrupleVisitor) VisitTerm(ctx *grammar.TermContext) *ExpressionResult {
	v.debugLog("VisitTerm: %s", ctx.GetText())

	// CORRECCIÓN: casting para factor
	var result *ExpressionResult
	if factorCtx, ok := ctx.Factor(0).(*grammar.FactorContext); ok && factorCtx != nil {
		result = v.VisitFactor(factorCtx)
	}

	factors := ctx.AllFactor()
	mulops := ctx.AllMulop()

	for i := 0; i < len(mulops); i++ {
		operator := mulops[i].GetText()

		var rightFactor *ExpressionResult
		if factorCtx, ok := factors[i+1].(*grammar.FactorContext); ok && factorCtx != nil {
			rightFactor = v.VisitFactor(factorCtx)
		}

		if result != nil && rightFactor != nil {
			v.debugLog("Processing mul/div: %v %s %v", result.Address, operator, rightFactor.Address)
			result = v.generateBinaryOperation(operator, result, rightFactor)
		}
	}

	return result
}

func (v *PureQuadrupleVisitor) VisitFactor(ctx *grammar.FactorContext) *ExpressionResult {
	v.debugLog("VisitFactor: %s", ctx.GetText())

	if ctx.Expression() != nil {
		if exprCtx, ok := ctx.Expression().(*grammar.ExpressionContext); ok && exprCtx != nil {
			return v.VisitExpression(exprCtx)
		}
	}

	if ctx.Addop() != nil {
		if valueCtx, ok := ctx.Value().(*grammar.ValueContext); ok && valueCtx != nil {
			valueResult := v.VisitValue(valueCtx)
			sign := ctx.Addop().GetText()

			if sign == "-" {
				tempAddress := v.getTemporaryAddress(valueResult.Type)
				constNegOne := v.ConstantTable.AddConstant("-1", valueResult.Type)
				v.quadQueue.Add("*", constNegOne, valueResult.Address, tempAddress)
				return &ExpressionResult{Address: tempAddress, Type: valueResult.Type}
			}
		}
	}

	if valueCtx, ok := ctx.Value().(*grammar.ValueContext); ok && valueCtx != nil {
		return v.VisitValue(valueCtx)
	}

	return &ExpressionResult{Address: 0, Type: "error"}
}

func (v *PureQuadrupleVisitor) VisitValue(ctx *grammar.ValueContext) *ExpressionResult {
	v.debugLog("VisitValue: %s", ctx.GetText())

	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		variable, found := v.Directory.FindVariableDeep(varName)
		if !found {
			v.addError(fmt.Sprintf("Variable '%s' no definida", varName))
			return &ExpressionResult{Address: 0, Type: "error"}
		}

		var address interface{}
		if variable.MemoryAddress == 0 {
			address = varName
		} else {
			address = variable.MemoryAddress
		}

		return &ExpressionResult{Address: address, Type: variable.Type}
	}

	if ctx.Constant() != nil {
		constValue := ctx.Constant().GetText()
		constType := semantic.InferTypeFromConstant(constValue)
		address := v.ConstantTable.AddConstant(constValue, constType)

		return &ExpressionResult{Address: address, Type: constType}
	}

	return &ExpressionResult{Address: 0, Type: "error"}
}

// ========== ASIGNACIONES ==========

func (v *PureQuadrupleVisitor) VisitAssignment(ctx *grammar.AssignContext) {
	variableName := ctx.ID().GetText()
	v.debugLog("Assignment: %s", variableName)

	// CORRECCIÓN: casting para expression
	if exprCtx, ok := ctx.Expression().(*grammar.ExpressionContext); ok && exprCtx != nil {
		exprResult := v.VisitExpression(exprCtx)
		if exprResult.Type == "error" {
			return
		}

		variable, exists := v.Directory.FindVariableDeep(variableName)
		if !exists {
			v.addError(fmt.Sprintf("Variable '%s' no encontrada", variableName))
			return
		}

		if exprResult.Type != variable.Type {
			v.addError(fmt.Sprintf("Tipo incompatible en asignación: %s = %s", variable.Type, exprResult.Type))
			return
		}

		address := variable.MemoryAddress
		if address == 0 {
			v.quadQueue.Add("=", exprResult.Address, "", variableName)
		} else {
			v.quadQueue.Add("=", exprResult.Address, "", address)
		}
	}
}

// ========== PRINT ==========

func (v *PureQuadrupleVisitor) VisitPrintStatement(ctx *grammar.Print_stmtContext) {
	v.debugLog("Print Statement")

	for _, pexpr := range ctx.AllPrintexpr() {
		switch {
		case pexpr.STRING() != nil:
			stringValue := pexpr.STRING().GetText()
			v.quadQueue.Add("print", stringValue, "", "")

		case pexpr.Expression() != nil:

			if exprCtx, ok := pexpr.Expression().(*grammar.ExpressionContext); ok && exprCtx != nil {
				exprResult := v.VisitExpression(exprCtx)
				if exprResult.Type != "error" {
					v.quadQueue.Add("print", exprResult.Address, "", "")
				}
			}
		}
	}
}

// ========== CONDICIONALES ==========

func (v *PureQuadrupleVisitor) VisitCondition(ctx *grammar.ConditionContext) {
	v.debugLog("=== Processing Condition ===")

	if exprCtx, ok := ctx.Expression().(*grammar.ExpressionContext); ok && exprCtx != nil {
		condResult := v.VisitExpression(exprCtx)
		if condResult.Type == "error" {
			return
		}

		if condResult.Type != "bool" {
			v.addError(fmt.Sprintf("Condición no booleana: %v", condResult.Type))
			return
		}

		gotoFIndex := v.quadQueue.Add("GOTOF", condResult.Address, "", nil)
		v.jumpStack = append(v.jumpStack, gotoFIndex)
	}

	if bodyCtx, ok := ctx.Body().(*grammar.BodyContext); ok && bodyCtx != nil {
		v.VisitBody(bodyCtx)
	}

	if ctx.Else_part() != nil && ctx.Else_part().Body() != nil {
		gotoEndIndex := v.quadQueue.Add("GOTO", "", "", nil)

		elseStart := v.quadQueue.Size()
		lastGotoF := v.jumpStack[len(v.jumpStack)-1]
		v.quadQueue.FillJump(lastGotoF, elseStart)
		v.jumpStack = v.jumpStack[:len(v.jumpStack)-1]

		if elseBodyCtx, ok := ctx.Else_part().Body().(*grammar.BodyContext); ok && elseBodyCtx != nil {
			v.VisitBody(elseBodyCtx)
		}

		endPos := v.quadQueue.Size()
		v.quadQueue.FillJump(gotoEndIndex, endPos)
	} else {
		endPos := v.quadQueue.Size()
		lastGotoF := v.jumpStack[len(v.jumpStack)-1]
		v.quadQueue.FillJump(lastGotoF, endPos)
		v.jumpStack = v.jumpStack[:len(v.jumpStack)-1]
	}
}

// ========== FUNCIONES ==========

func (v *PureQuadrupleVisitor) VisitFunction(ctx *grammar.FuncContext) {
	functionName := ctx.ID().GetText()
	v.debugLog("Processing Function: %s", functionName)

	err := v.Directory.EnterFunction(functionName)
	if err != nil {
		v.addError(fmt.Sprintf("Error entrando a función '%s': %v", functionName, err))
		return
	}

	startQuad := v.quadQueue.GenerateFUNC(functionName, "void")
	v.functionStarts[functionName] = startQuad

	// Procesar parámetros
	if ctx.Param_list() != nil {
		for i, paramCtx := range ctx.Param_list().AllParam() {
			paramType := paramCtx.Type_().GetText()
			paramName := paramCtx.ID().GetText()

			// Solo generar cuádruple, NO agregar a tabla (ya lo hizo DirectoryBuilder)
			address := 4000 + i
			v.quadQueue.GeneratePARAM(paramType, paramName, address)
			v.debugLog("Generated PARAM quadruple for '%s' (%s)", paramName, paramType)
		}
	}

	if ctx.Funcbody() != nil && ctx.Funcbody().Body() != nil {
		if funcBodyCtx, ok := ctx.Funcbody().Body().(*grammar.BodyContext); ok && funcBodyCtx != nil {
			v.VisitBody(funcBodyCtx)
		}
	}

	// Generar ENDFUNC
	v.quadQueue.GenerateENDFUNC()
	endQuad := v.quadQueue.Size() - 1
	v.Directory.SetFunctionQuadruples(functionName, startQuad, endQuad)

	err = v.Directory.ExitFunction()
	if err != nil {
		v.addError(fmt.Sprintf("Error saliendo de función '%s': %v", functionName, err))
	}

	v.debugLog("Function '%s' processed successfully", functionName)
}

func (v *PureQuadrupleVisitor) VisitFunctionCall(ctx *grammar.F_callContext) {
	funcName := ctx.ID().GetText()
	v.debugLog("Processing Function Call: %s", funcName)

	numArgs := 0
	if ctx.Arg_list() != nil {
		numArgs = len(ctx.Arg_list().AllExpression())
	}

	if err := v.Directory.ValidateFunctionCall(funcName, numArgs); err != nil {
		v.addError(err.Error())
		return
	}

	paramValues := make([]interface{}, numArgs)
	if ctx.Arg_list() != nil {
		for i, argExpr := range ctx.Arg_list().AllExpression() {
			// CORRECCIÓN: casting para expression en argumentos
			if exprCtx, ok := argExpr.(*grammar.ExpressionContext); ok && exprCtx != nil {
				argResult := v.VisitExpression(exprCtx)
				if argResult.Type == "error" {
					v.addError(fmt.Sprintf("Error en argumento %d de función %s", i+1, funcName))
					return
				}
				paramValues[i] = argResult.Address
			}
		}
	}

	funcInfo, err := v.Directory.GetFunctionInfo(funcName)
	if err != nil {
		v.addError(err.Error())
		return
	}

	localVars := v.Directory.CountLocalVariables(funcName)
	tempVars := v.Directory.CountTempVariables(funcName)
	totalSize := localVars + tempVars + numArgs
	v.quadQueue.GenerateERA(funcName, totalSize)

	for i, paramValue := range paramValues {
		v.quadQueue.GeneratePARAMETER(paramValue, i)
	}

	startAddress := funcInfo.StartQuadruple
	v.quadQueue.GenerateGOSUB(funcName, startAddress)
}

// ========== MÉTODOS AUXILIARES ==========

func (v *PureQuadrupleVisitor) generateBinaryOperation(operator string, left, right *ExpressionResult) *ExpressionResult {
	resultType, ok := semantic.Cube.GetResultType(left.Type, right.Type, operator)
	if !ok || resultType == "error" {
		v.addError(fmt.Sprintf("Tipo incompatible: %s %s %s", left.Type, operator, right.Type))
		return &ExpressionResult{Address: 0, Type: "error"}
	}

	tempAddress := v.getTemporaryAddress(resultType)
	v.quadQueue.Add(operator, left.Address, right.Address, tempAddress)

	return &ExpressionResult{Address: tempAddress, Type: resultType}
}

func (v *PureQuadrupleVisitor) getTemporaryAddress(resultType string) interface{} {
	tempVar := v.Directory.NewTempVar(resultType)
	return tempVar.MemoryAddress
}

func (v *PureQuadrupleVisitor) addError(msg string) {
	for _, e := range v.Errors {
		if e == msg {
			return
		}
	}
	v.Errors = append(v.Errors, msg)
}

func (v *PureQuadrupleVisitor) debugLog(msg string, args ...interface{}) {
	if v.debugExpressions {
		fmt.Printf("[PURE VISITOR] "+msg+"\n", args...)
	}
}

// ========== INTERFAZ PÚBLICA ==========

func (v *PureQuadrupleVisitor) PrintQuadruples() {
	fmt.Println("=== Cuádruplos generados ===")
	for i, quad := range v.quadQueue.GetAll() {
		fmt.Printf("%d: (%s, %v, %v, %v)\n", i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result)
	}
}

func (v *PureQuadrupleVisitor) GetQuadruples() []quads.Quadruple {
	if v.quadQueue == nil {
		return []quads.Quadruple{}
	}
	return v.quadQueue.GetAll()
}

func (v *PureQuadrupleVisitor) HasErrors() bool {
	return len(v.Errors) > 0
}

func (v *PureQuadrupleVisitor) GetErrors() []string {
	return v.Errors
}
