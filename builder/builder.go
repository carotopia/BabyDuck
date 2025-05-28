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

	// Generar cuádruplos para parámetros
	if ctx.Param_list() != nil {
		paramListCtx := ctx.Param_list().(*grammar.Param_listContext)
		for i, paramCtx := range paramListCtx.AllParam() {
			param := paramCtx.(*grammar.ParamContext)
			paramType := param.Type_().GetText()
			paramName := param.ID().GetText()

			// Asignar direcciones en el rango local (4000-4999)
			address := 4000 + i
			d.QuadVisitor.HandleParameterDeclaration(paramType, paramName, address)
		}
	}
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

func (d *DirectoryBuilder) EnterParam(ctx *grammar.ParamContext) {
	paramName := ctx.ID().GetText()
	paramType := ctx.Type_().GetText()

	if err := d.Directory.AddVariable(paramName, paramType); err != nil {
		d.addError(err.Error())
	}
}

// ========== CALLBACKS PARA GENERACIÓN DE CUÁDRUPLOS ==========
// ENFOQUE LIMPIO: Solo callbacks esenciales, sin manejo de pilas

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

func (d *DirectoryBuilder) ExitCycle(ctx *grammar.CycleContext) {
	// Los cycles también pueden tener expresiones que necesitan evaluación
	// pero el visitor las manejará automáticamente cuando se encuentren
}

func (d *DirectoryBuilder) ExitProgram(ctx *grammar.ProgramContext) {
	if d.QuadVisitor != nil {
		d.QuadVisitor.PrintQuadruples()
	}
	if d.ConstantTable != nil {
		d.ConstantTable.Print()
	}
}

// ========== MÉTODOS DE VALIDACIÓN SIMPLIFICADOS ==========
// Ya no necesitamos validar expresiones manualmente porque el visitor las maneja

func (d *DirectoryBuilder) EnterExpression(ctx *grammar.ExpressionContext) {
	d.debugLog("Entrando a expresión: %s", ctx.GetText())
}

func (d *DirectoryBuilder) EnterMain(ctx interface{}) {
	d.debugLog("Entrando al bloque main")
	if len(d.Directory.CurrentScope) == 0 {
		d.Directory.CurrentScope = append(d.Directory.CurrentScope, "main")
	}
}

func (d *DirectoryBuilder) ExitMain(ctx interface{}) {
	d.debugLog("Saliendo del bloque main")
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