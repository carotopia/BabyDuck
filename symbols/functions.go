package symbols

import "fmt"

// FunctionDirectory structure with two key components
// Directory is a dictionary where every entrey represents a function or a scope
// and every value is a new table of variables
type FunctionDirectory struct {
	Directory    map[string]VariableTable
	CurrentScope []string
}

// NewFunctionDirectory creates a new instance of function directory
// initializes Directory as an empty map of functions
// sets current scope to program the root of the global

func NewFunctionDirectory() *FunctionDirectory {
	directory := &FunctionDirectory{
		Directory:    make(map[string]VariableTable),
		CurrentScope: []string{"program"},
	}

	directory.Directory["program"] = make(VariableTable)

	return directory
}

// Adds a new function to the directory
// First checks if it already exists and returns an error if it already exists
// Creates a new table of variables for this function

func (fd *FunctionDirectory) AddFunction(functionName string) error {

	if _, exists := fd.Directory[functionName]; exists {
		return fmt.Errorf("error: function '%s' already declared in current scope", functionName)
	}

	fd.Directory[functionName] = make(VariableTable)

	return nil
}
