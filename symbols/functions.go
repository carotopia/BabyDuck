package symbols

import (
	"BabyDuckCompiler/memory"
	"fmt"
)

// FunctionDirectory structure with two key components
// Directory is a dictionary where every entrey represents a function or a scope
// and every value is a new table of variables
type FunctionDirectory struct {
	Directory    map[string]*FunctionInfo
	CurrentScope []string
	TempCounter  int
	Memory       *memory.MemoryManager
}

// NewFunctionDirectory creates a new instance of function directory
// initializes Directory as an empty map of functions
// sets current scope to program the root of the global

func NewFunctionDirectory(mem *memory.MemoryManager) *FunctionDirectory {
	directory := &FunctionDirectory{
		Directory:    make(map[string]*FunctionInfo),
		CurrentScope: []string{"program"},
		Memory:       mem,
	}

	directory.Directory["program"] = &FunctionInfo{
		Params:    []Variable{},
		Variables: make(VariableTable),
	}

	return directory
}

type FunctionInfo struct {
	Params    []Variable
	Variables VariableTable
}

// Adds a new function to the directory
// First checks if it already exists and returns an error if it already exists
// Creates a new table of variables for this function

func (fd *FunctionDirectory) AddFunction(functionName string, params []Variable) error {

	if _, exists := fd.Directory[functionName]; exists {
		return fmt.Errorf("error: function '%s' already declared in current scope", functionName)
	}

	fd.Directory[functionName] = &FunctionInfo{
		Params:    params,
		Variables: make(VariableTable),
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

func (fd *FunctionDirectory) NewTempVar(resultType string) Variable {
	fd.TempCounter++

	var address int
	switch resultType {
	case "int":
		address = fd.Memory.NextTempInt()
	case "float":
		address = fd.Memory.NextTempFloat()
	case "bool":
		address = fd.Memory.NextTempBool()
	default:
		address = fd.Memory.NextTempInt()
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
