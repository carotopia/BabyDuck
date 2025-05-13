package symbols

import "fmt"

// Variable is a structure to create a new variable, including type,value and memory address
type Variable struct {
	Type          string
	Value         interface{}
	MemoryAddress int
}

// VariableTable a map of variables names to variable structures
type VariableTable map[string]Variable

// AddVariable adds a new variable to the current scope, and checks if there is an current scope defined
// If the variable is already in the current scope returns error
// If the variable is missing an attribute returns error
func (fd *FunctionDirectory) AddVariable(name string, varType string) error {
	if len(fd.CurrentScope) == 0 {
		return fmt.Errorf("error: no active scope defined")
	}

	context := fd.CurrentScope[len(fd.CurrentScope)-1]
	varTable := fd.Directory[context]

	if _, exists := varTable[name]; exists {
		return fmt.Errorf("error: variable '%s' already declared in scope '%s'", name, context)
	}
	mockAddress := 1000 + len(varType)

	varTable[name] = Variable{
		Type:          varType,
		Value:         nil,
		MemoryAddress: mockAddress,
	}

	return nil
}

// Validates if the variable exists within the current scope
// LookUpVariable is called in every scope from inside to outside
// If FindVariable founds the variable returns nil, return error
func (fd *FunctionDirectory) ValidateVariable(scopes []string, name string) error {
	for i := len(scopes) - 1; i >= 0; i-- {
		scope := scopes[i]
		_, exists := fd.FindVariable(scope, name)
		if exists {
			return nil
		}
	}

	return fmt.Errorf("error: undefined variable '%s'", name)
}

// Looks for variable in a specific scope
// Returns the type and wether it was found or not
func (fd *FunctionDirectory) FindVariable(scope string, name string) (string, bool) {
	varTable, scopeExists := fd.Directory[scope]
	if !scopeExists {
		return "", false
	}

	variable, varExists := varTable[name]
	return variable.Type, varExists
}
