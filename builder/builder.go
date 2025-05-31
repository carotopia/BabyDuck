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

// ========== MANEJO DE PROGRAMA ==========

func (d *DirectoryBuilder) EnterProgram(ctx *grammar.ProgramContext) {
	programName := ctx.ID().GetText()
	d.debugLog("Entrando al programa: %s", programName)
	if len(d.Directory.CurrentScope) == 0 || d.Directory.CurrentScope[0] != "program" {
		d.Directory.CurrentScope = []string{"program"} // ← Inicializar limpio
	}
}

// ========== MANEJO DE VARIABLES GLOBALES ==========

func (d *DirectoryBuilder) EnterVars(ctx *grammar.VarsContext) {
	d.debugLog("Entrando a bloque de declaraciones de variables en scope: %s", d.Directory.GetCurrentScope())
}

func (d *DirectoryBuilder) ExitVar_decl(ctx *grammar.Var_declContext) {
	variableType := ctx.Type_().GetText()
	currentScope := d.Directory.GetCurrentScope()

	for _, idToken := range ctx.Id_list().AllID() {
		varName := idToken.GetText()

		if err := d.Directory.AddVariable(varName, variableType); err != nil {
			d.addError(err.Error())
			continue
		}

		d.debugLog("Variable declarada: %s (%s) en scope: %s", varName, variableType, currentScope)
	}
}

func (d *DirectoryBuilder) EnterFunc(ctx *grammar.FuncContext) {
	functionName := ctx.ID().GetText()
	d.debugLog("Entrando a función: %s", functionName)

	// 🔧 CREAR FUNCIÓN SIN PARÁMETROS INICIALMENTE
	params := []symbols.Variable{}
	if err := d.Directory.AddFunction(functionName, params); err != nil {
		d.addError(err.Error())
		return
	}
	d.Directory.CurrentScope = []string{"program", functionName}
	d.debugLog("Función declarada: %s, scope actual: %v", functionName, d.Directory.CurrentScope)
}

func (d *DirectoryBuilder) ExitFunc(ctx *grammar.FuncContext) {
	functionName := ctx.ID().GetText()
	d.debugLog("Saliendo de función: %s", functionName)

	d.Directory.CurrentScope = []string{"program"}
	d.debugLog("Scope después de salir: %v", d.Directory.CurrentScope)
}

func (d *DirectoryBuilder) EnterParam(ctx *grammar.ParamContext) {
	paramName := ctx.ID().GetText()
	paramType := ctx.Type_().GetText()

	currentScope := d.Directory.GetCurrentScope()
	if currentScope == "program" {
		d.addError("Error interno: parámetro fuera de función")
		return
	}

	if err := d.Directory.AddFunctionParameter(currentScope, paramName, paramType); err != nil {
		d.addError(err.Error())
		return
	}

	d.debugLog("Parámetro declarado: %s (%s) en función: %s", paramName, paramType, currentScope)
}

func (d *DirectoryBuilder) EnterFuncbody(ctx *grammar.FuncbodyContext) {
	d.debugLog("Entrando al cuerpo de función en scope: %s", d.Directory.GetCurrentScope())

	if ctx.Vars() != nil {
		d.debugLog("Función tiene variables locales que serán procesadas")
	}
}

func (d *DirectoryBuilder) ExitFuncbody(ctx *grammar.FuncbodyContext) {
	d.debugLog("Saliendo del cuerpo de función")
}

func (d *DirectoryBuilder) extractParameters(ctx *grammar.FuncContext) []symbols.Variable {
	var params []symbols.Variable

	if ctx.Param_list() == nil {
		return params
	}

	paramListCtx := ctx.Param_list()
	for i, paramCtx := range paramListCtx.AllParam() {
		param := paramCtx
		paramName := param.ID().GetText()
		varType := param.Type_().GetText()

		// Crear variable con dirección local
		address := 4000 + i
		params = append(params, symbols.Variable{
			Type:          varType,
			Value:         nil,
			MemoryAddress: address,
		})

		d.debugLog("Parámetro extraído: %s (%s) -> %d", paramName, varType, address)
	}

	return params
}

func (d *DirectoryBuilder) addError(msg string) {

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
