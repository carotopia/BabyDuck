package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/symbols"
	"fmt"
)

// DirectoryBuilder tabla de simbolos
type DirectoryBuilder struct {
	*grammar.BaseBabyDuckListener

	Directory     *symbols.FunctionDirectory
	ConstantTable *symbols.ConstantTable
	Errors        []string
	Debug         bool
}

// NewDirectoryBuilder
func NewDirectoryBuilder(debug bool, funcDir *symbols.FunctionDirectory, constTable *symbols.ConstantTable) *DirectoryBuilder {
	return &DirectoryBuilder{
		BaseBabyDuckListener: &grammar.BaseBabyDuckListener{},
		Directory:            funcDir,
		ConstantTable:        constTable,
		Errors:               []string{},
		Debug:                debug,
	}
}

// ========== SOLO MANEJO DE PROGRAMA ==========

func (d *DirectoryBuilder) EnterProgram(ctx *grammar.ProgramContext) {
	programName := ctx.ID().GetText()
	d.debugLog("Entrando al programa: %s", programName)
	d.Directory.CurrentScope = append(d.Directory.CurrentScope, "program")
}

// ========== SOLO MANEJO DE VARIABLES ==========

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

// ========== SOLO MANEJO DE FUNCIONES ==========

func (d *DirectoryBuilder) EnterFunc(ctx *grammar.FuncContext) {
	functionName := ctx.ID().GetText()
	params := d.extractParameters(ctx)

	if err := d.Directory.AddFunction(functionName, params); err != nil {
		d.addError(err.Error())
	}

	d.Directory.CurrentScope = append(d.Directory.CurrentScope, functionName)
	d.debugLog("Función declarada: %s", functionName)
}

func (d *DirectoryBuilder) ExitFunc(ctx *grammar.FuncContext) {
	functionName := ctx.ID().GetText()
	d.debugLog("Saliendo de función: %s", functionName)
	d.Directory.CurrentScope = d.Directory.CurrentScope[:len(d.Directory.CurrentScope)-1]
}

func (d *DirectoryBuilder) extractParameters(ctx *grammar.FuncContext) []symbols.Variable {
	var params []symbols.Variable

	if ctx.Param_list() == nil {
		return params
	}

	paramListCtx := ctx.Param_list()
	for _, paramCtx := range paramListCtx.AllParam() {
		param := paramCtx
		varType := param.Type_().GetText()
		params = append(params, symbols.Variable{Type: varType})
	}

	return params
}

func (d *DirectoryBuilder) EnterParam(ctx *grammar.ParamContext) {
	paramName := ctx.ID().GetText()
	paramType := ctx.Type_().GetText()

	if err := d.Directory.AddVariable(paramName, paramType); err != nil {
		d.addError(err.Error())
	}

	d.debugLog("Parámetro declarado: %s (%s)", paramName, paramType)
}

func (d *DirectoryBuilder) addError(msg string) {
	// Evitar duplicados
	for _, e := range d.Errors {
		if e == msg {
			return
		}
	}
	d.Errors = append(d.Errors, msg)
}

func (d *DirectoryBuilder) debugLog(format string, args ...interface{}) {
	if d.Debug {
		fmt.Printf("[Builder] "+format+"\n", args...)
	}
}
