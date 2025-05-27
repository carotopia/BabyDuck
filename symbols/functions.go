package symbols

import (
	"fmt"
)

// FunctionDirectory structure with two key components
// Directory is a dictionary where every entry represents a function or a scope
// and every value is a new table of variables
type FunctionDirectory struct {
	Directory    map[string]*FunctionInfo
	CurrentScope []string
	TempCounter  int
	// No usar MemoryManager por ahora - asignar direcciones manualmente
}

// NewFunctionDirectory creates a new instance of function directory
// initializes Directory as an empty map of functions
// sets current scope to program the root of the global
func NewFunctionDirectory() *FunctionDirectory { // ← Sin parámetros
	directory := &FunctionDirectory{
		Directory:    make(map[string]*FunctionInfo),
		CurrentScope: []string{"program"},
		TempCounter:  0,
	}

	directory.Directory["program"] = &FunctionInfo{
		Params:         []Variable{},
		StartQuadruple: 0,
		EndQuadruple:   -1,
		LocalVarCount:  0,
		TempVarCount:   0,
		Variables:      make(VariableTable),
	}

	return directory
}

type FunctionInfo struct {
	Params         []Variable
	StartQuadruple int
	EndQuadruple   int
	LocalVarCount  int
	TempVarCount   int
	Variables      VariableTable
}

// Adds a new function to the directory
func (fd *FunctionDirectory) AddFunction(functionName string, params []Variable) error {
	if _, exists := fd.Directory[functionName]; exists {
		return fmt.Errorf("error: function '%s' already declared in current scope", functionName)
	}

	fd.Directory[functionName] = &FunctionInfo{
		Params:         params,
		StartQuadruple: -1,
		EndQuadruple:   -1,
		LocalVarCount:  0,
		TempVarCount:   0,
		Variables:      make(VariableTable),
	}

	return nil
}

func (fd *FunctionDirectory) ValidateFunctionCall(name string, numArgs int) error {
	funcInfo, exists := fd.Directory[name]
	if !exists {
		return fmt.Errorf("error: function '%s' is not declared", name)
	}

	expected := len(funcInfo.Params)
	if expected != numArgs {
		return fmt.Errorf("error: function '%s' expects %d arguments, got %d", name, expected, numArgs)
	}

	return nil
}

func (fd *FunctionDirectory) Error() {

}

// Contadores para asignar direcciones manualmente
var globalIntCounter = 1000
var globalFloatCounter = 2000
var globalBoolCounter = 3000
var localIntCounter = 4000
var localFloatCounter = 5000
var localBoolCounter = 6000
var tempIntCounter = 7000
var tempFloatCounter = 8000
var tempBoolCounter = 9000

func (fd *FunctionDirectory) NewTempVar(resultType string) Variable {
	fd.TempCounter++

	var address int
	switch resultType {
	case "int":
		address = tempIntCounter
		tempIntCounter++
	case "float":
		address = tempFloatCounter
		tempFloatCounter++
	case "bool":
		address = tempBoolCounter
		tempBoolCounter++
	default:
		address = tempIntCounter
		tempIntCounter++
	}

	tempName := fmt.Sprintf("temp%d", fd.TempCounter)
	newVar := Variable{
		Type:          resultType,
		Value:         nil,
		MemoryAddress: address,
	}
	currentScope := fd.CurrentScope[len(fd.CurrentScope)-1]
	fd.Directory[currentScope].Variables[tempName] = newVar

	return newVar
}

func (fd *FunctionDirectory) GetCurrentScope() string {
	if len(fd.CurrentScope) == 0 {
		return "global"
	}
	return fd.CurrentScope[len(fd.CurrentScope)-1]
}

// ========== MÉTODOS NUEVOS PARA FUNCIONES ==========

func (fd *FunctionDirectory) SetFunctionQuadruples(functionName string, start, end int) error {
	funcInfo, exists := fd.Directory[functionName]
	if !exists {
		return fmt.Errorf("función '%s' no encontrada", functionName)
	}

	funcInfo.StartQuadruple = start
	if end != -1 {
		funcInfo.EndQuadruple = end
	}
	return nil
}

func (fd *FunctionDirectory) GetFunctionQuadruples(functionName string) (int, int, error) {
	funcInfo, exists := fd.Directory[functionName]
	if !exists {
		return -1, -1, fmt.Errorf("función '%s' no encontrada", functionName)
	}

	return funcInfo.StartQuadruple, funcInfo.EndQuadruple, nil
}

func (fd *FunctionDirectory) GetFunctionInfo(functionName string) (*FunctionInfo, error) {
	funcInfo, exists := fd.Directory[functionName]
	if !exists {
		return nil, fmt.Errorf("función '%s' no encontrada", functionName)
	}

	return funcInfo, nil
}

func (fd *FunctionDirectory) CountLocalVariables(functionName string) int {
	funcInfo, exists := fd.Directory[functionName]
	if !exists {
		return 0
	}

	count := 0
	for _, variable := range funcInfo.Variables {
		// Contar solo variables locales (rango 4000-6999)
		if variable.MemoryAddress >= 4000 && variable.MemoryAddress <= 6999 {
			count++
		}
	}

	funcInfo.LocalVarCount = count
	return count
}

func (fd *FunctionDirectory) CountTempVariables(functionName string) int {
	funcInfo, exists := fd.Directory[functionName]
	if !exists {
		return 0
	}

	count := 0
	for _, variable := range funcInfo.Variables {
		// Contar temporales (rango 7000-9999)
		if variable.MemoryAddress >= 7000 && variable.MemoryAddress <= 9999 {
			count++
		}
	}

	funcInfo.TempVarCount = count
	return count
}

func (fd *FunctionDirectory) UpdateFunctionStats(functionName string) {
	fd.CountLocalVariables(functionName)
	fd.CountTempVariables(functionName)
}

func (fd *FunctionDirectory) PrintFunctionInfo() {
	fmt.Println("\n=== FUNCTION INFORMATION ===")
	for name, info := range fd.Directory {
		fmt.Printf("\nFunction: %s\n", name)
		fmt.Printf("  Start Quadruple: %d\n", info.StartQuadruple)
		fmt.Printf("  End Quadruple: %d\n", info.EndQuadruple)
		fmt.Printf("  Parameters: %d\n", len(info.Params))
		fmt.Printf("  Local Variables: %d\n", info.LocalVarCount)
		fmt.Printf("  Temp Variables: %d\n", info.TempVarCount)
		fmt.Printf("  Total Variables: %d\n", len(info.Variables))
	}
	fmt.Println("=============================")
}
