package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/symbols"
	"fmt"
)

// DirectoryBuilder maneja la construcción del directorio de símbolos y validación semántica
type DirectoryBuilder struct {
	*grammar.BaseBabyDuckListener

	Directory     *symbols.FunctionDirectory
	ConstantTable *symbols.ConstantTable
	Errors        []string
	Debug         bool

	// Visitor para generación de cuádruplos
	QuadVisitor *QuadrupleVisitor
}

// NewDirectoryBuilder crea e inicializa un nuevo DirectoryBuilder
func NewDirectoryBuilder(debug bool, funcDir *symbols.FunctionDirectory, constTable *symbols.ConstantTable) *DirectoryBuilder {
	builder := &DirectoryBuilder{
		BaseBabyDuckListener: &grammar.BaseBabyDuckListener{},
		Directory:            funcDir,
		ConstantTable:        constTable,
		Errors:               []string{},
		Debug:                debug,
	}

	// Inicializar el visitor de cuádruplos
	builder.QuadVisitor = NewQuadrupleVisitor(builder)

	return builder
}

// ========== MANEJO DE FUNCIONES ==========

func (d *DirectoryBuilder) EnterProgram(ctx *grammar.ProgramContext) {
	programName := ctx.ID().GetText()
	d.debugLog("Entrando al programa: %s", programName)

	// Inicializar scope del programa
	d.Directory.CurrentScope = append(d.Directory.CurrentScope, "program")
}

func (d *DirectoryBuilder) EnterFunc(ctx *grammar.FuncContext) {
	functionName := ctx.ID().GetText()
	params := d.extractParameters(ctx)

	if err := d.Directory.AddFunction(functionName, params); err != nil {
		d.addError(err.Error())
	}

	d.Directory.CurrentScope = append(d.Directory.CurrentScope, functionName)

	// Notificar al visitor que estamos entrando a una función
	d.QuadVisitor.EnterFunction(functionName)
}

func (d *DirectoryBuilder) ExitFunc(ctx *grammar.FuncContext) {
	functionName := ctx.ID().GetText()

	// Notificar al visitor que estamos saliendo de la función
	d.QuadVisitor.ExitFunction(functionName)

	// Salir del scope de la función
	d.Directory.CurrentScope = d.Directory.CurrentScope[:len(d.Directory.CurrentScope)-1]
}

func (d *DirectoryBuilder) extractParameters(ctx *grammar.FuncContext) []symbols.Variable {
	var params []symbols.Variable

	if ctx.Param_list() == nil {
		return params
	}

	paramListCtx := ctx.Param_list().(*grammar.Param_listContext)
	for _, paramCtx := range paramListCtx.AllParam() {
		param := paramCtx.(*grammar.ParamContext)
		varType := param.Type_().GetText()
		params = append(params, symbols.Variable{Type: varType})
	}

	return params
}

// ========== MANEJO DE VARIABLES ==========

func (d *DirectoryBuilder) EnterVars(ctx *grammar.VarsContext) {
	d.debugLog("Entrando a bloque de declaraciones de variables")
}

func (d *DirectoryBuilder) ExitVar_decl(ctx *grammar.Var_declContext) {
	variableType := ctx.Type_().GetText()

	for _, idToken := range ctx.Id_list().AllID() {
		varName := idToken.GetText()

		if err := d.Directory.AddVariable(varName, variableType); err != nil {
			d.addError(err.Error())
			continue
		}

		d.debugLog("Variable declarada: %s (%s)", varName, variableType)
	}
}

// Alternativo: si tu gramática usa diferentes nombres de contexto
func (d *DirectoryBuilder) ExitVariable_declaration(ctx interface{}) {
	// Handle different context types for variable declarations
	d.debugLog("Manejando declaración de variable alternativa")
}

func (d *DirectoryBuilder) EnterParam(ctx *grammar.ParamContext) {
	paramName := ctx.ID().GetText()
	paramType := ctx.Type_().GetText()

	if err := d.Directory.AddVariable(paramName, paramType); err != nil {
		d.addError(err.Error())
	}
}

// ========== VALIDACIÓN DE EXPRESIONES ==========

func (d *DirectoryBuilder) validateExpression(exp grammar.IExpressionContext) {
	if exp != nil && exp.Rel_expr() != nil {
		d.validateRelExpr(exp.Rel_expr())
	}
}

func (d *DirectoryBuilder) validateRelExpr(ctx grammar.IRel_exprContext) {
	if ctx == nil {
		return
	}

	if left := ctx.Add_expr(0); left != nil {
		d.validateAddExpr(left)
	}

	if ctx.Relop() != nil && len(ctx.AllAdd_expr()) > 1 {
		if right := ctx.Add_expr(1); right != nil {
			d.validateAddExpr(right)
		}
	}
}

func (d *DirectoryBuilder) validateAddExpr(add grammar.IAdd_exprContext) {
	if add == nil {
		return
	}

	for _, term := range add.AllTerm() {
		d.validateTerm(term)
	}
}

func (d *DirectoryBuilder) validateTerm(ctx grammar.ITermContext) {
	for _, factor := range ctx.AllFactor() {
		d.validateFactor(factor)
	}
}

func (d *DirectoryBuilder) validateFactor(ctx grammar.IFactorContext) {
	switch {
	case ctx.Expression() != nil:
		d.validateExpression(ctx.Expression())
	case ctx.Value() != nil:
		d.validateValue(ctx.Value())
	}
}

func (d *DirectoryBuilder) validateValue(ctx grammar.IValueContext) {
	if id := ctx.ID(); id != nil {
		name := id.GetText()
		if err := d.Directory.ValidateVariable(name); err != nil {
			d.addError(err.Error())
		}
	}
}

// ========== CALLBACKS PARA GENERACIÓN DE CUÁDRUPLOS ==========

func (d *DirectoryBuilder) ExitAssign(ctx *grammar.AssignContext) {
	d.QuadVisitor.VisitAssignment(ctx)
}

func (d *DirectoryBuilder) ExitPrint_stmt(ctx *grammar.Print_stmtContext) {
	d.QuadVisitor.VisitPrintStatement(ctx)
}

func (d *DirectoryBuilder) ExitCondition(ctx *grammar.ConditionContext) {
	d.QuadVisitor.VisitCondition(ctx)
}

func (d *DirectoryBuilder) ExitBody(ctx *grammar.BodyContext) {
	d.QuadVisitor.VisitBody(ctx)
}

func (d *DirectoryBuilder) ExitElse_part(ctx *grammar.Else_partContext) {
	d.QuadVisitor.VisitElsePart(ctx)
}

func (d *DirectoryBuilder) ExitF_call(ctx *grammar.F_callContext) {
	// Validar que la función existe
	funcName := ctx.ID().GetText()
	numArgs := 0

	if ctx.Arg_list() != nil {
		argListCtx := ctx.Arg_list().(*grammar.Arg_listContext)
		numArgs = len(argListCtx.AllExpression())
	}

	if err := d.Directory.ValidateFunctionCall(funcName, numArgs); err != nil {
		d.addError(err.Error())
	}

	// Generar cuádruplos de llamada a función
	d.QuadVisitor.VisitFunctionCall(ctx)
}

func (d *DirectoryBuilder) ExitFactor(ctx *grammar.FactorContext) {
	d.QuadVisitor.VisitFactor(ctx)
}

func (d *DirectoryBuilder) ExitTerm(ctx *grammar.TermContext) {
	d.QuadVisitor.VisitTerm(ctx)
}

func (d *DirectoryBuilder) EnterMulop(ctx *grammar.MulopContext) {
	d.QuadVisitor.EnterMulop(ctx)
}

func (d *DirectoryBuilder) ExitAdd_expr(ctx *grammar.Add_exprContext) {
	d.QuadVisitor.VisitAddExpression(ctx)
}

func (d *DirectoryBuilder) EnterAddop(ctx *grammar.AddopContext) {
	d.QuadVisitor.EnterAddop(ctx)
}

func (d *DirectoryBuilder) ExitValue(ctx *grammar.ValueContext) {
	d.QuadVisitor.VisitValue(ctx)
}

func (d *DirectoryBuilder) ExitRel_expr(ctx *grammar.Rel_exprContext) {
	d.QuadVisitor.VisitRelationalExpression(ctx)
}

func (d *DirectoryBuilder) ExitCycle(ctx *grammar.CycleContext) {
	if ctx.Expression() != nil {
		d.validateExpression(ctx.Expression())
	}
}

func (d *DirectoryBuilder) ExitProgram(ctx *grammar.ProgramContext) {
	if d.QuadVisitor != nil {
		d.QuadVisitor.PrintQuadruples()
	}
	if d.ConstantTable != nil { // ← Agregar verificación nil
		d.ConstantTable.Print()
	}
}

func (d *DirectoryBuilder) EnterExpression(ctx *grammar.ExpressionContext) {
	d.debugLog("Entrando a expresión")
}

func (d *DirectoryBuilder) EnterTerm(ctx *grammar.TermContext) {
	// Este método puede estar vacío o manejar lógica específica si es necesario
}

// Métodos adicionales para manejar diferentes contextos de declaración
func (d *DirectoryBuilder) EnterMain(ctx interface{}) {
	d.debugLog("Entrando al bloque main")
	// Asegurar que estamos en el scope correcto
	if len(d.Directory.CurrentScope) == 0 {
		d.Directory.CurrentScope = append(d.Directory.CurrentScope, "main")
	}
}

func (d *DirectoryBuilder) ExitMain(ctx interface{}) {
	d.debugLog("Saliendo del bloque main")
}

// Método genérico para declaraciones que podrían tener nombres diferentes
func (d *DirectoryBuilder) handleVariableDeclaration(varType string, varNames []string) {
	for _, varName := range varNames {
		if err := d.Directory.AddVariable(varName, varType); err != nil {
			d.addError(err.Error())
			continue
		}
		d.debugLog("Variable declarada: %s (%s)", varName, varType)
	}
}

// ========== UTILIDADES ==========

func (d *DirectoryBuilder) addError(msg string) {
	for _, e := range d.Errors {
		if e == msg {
			return // Evitar duplicados
		}
	}
	d.Errors = append(d.Errors, msg)
}

func (d *DirectoryBuilder) debugLog(format string, args ...interface{}) {
	if d.Debug {
		fmt.Printf("[Debug] "+format+"\n", args...)
	}
}

// GetQuadVisitor returns the quadruple visitor for external access
func (d *DirectoryBuilder) GetQuadVisitor() *QuadrupleVisitor {
	return d.QuadVisitor
}
