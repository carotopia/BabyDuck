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
// ACTUALIZAR EL MÉTODO AddVariable EN TU variables.go:

func (fd *FunctionDirectory) AddVariable(name string, varType string) error {
	if len(fd.CurrentScope) == 0 {
		return fmt.Errorf("error: no active scope defined")
	}

	context := fd.CurrentScope[len(fd.CurrentScope)-1]
	funcInfo, exists := fd.Directory[context]
	if !exists {
		return fmt.Errorf("error: function '%s' is not declared", context)
	}

	if _, exists := funcInfo.Variables[name]; exists {
		return fmt.Errorf("error: variable '%s' already declared in scope '%s'", name, context)
	}

	var address int
	if context == "program" {
		// Variables globales - usar contadores globales
		switch varType {
		case "int":
			address = globalIntCounter
			globalIntCounter++
		case "float":
			address = globalFloatCounter
			globalFloatCounter++
		case "bool":
			address = globalBoolCounter
			globalBoolCounter++
		default:
			return fmt.Errorf("error: tipo de variable '%s' no soportado", varType)
		}
	} else {
		// Variables locales - usar contadores locales
		switch varType {
		case "int":
			address = localIntCounter
			localIntCounter++
		case "float":
			address = localFloatCounter
			localFloatCounter++
		case "bool":
			address = localBoolCounter
			localBoolCounter++
		default:
			return fmt.Errorf("error: tipo de variable '%s' no soportado", varType)
		}
	}

	funcInfo.Variables[name] = Variable{
		Type:          varType,
		Value:         nil,
		MemoryAddress: address,
	}

	return nil
}
func (fd *FunctionDirectory) ValidateVariable(name string) error {
	_, exists := fd.FindVariableDeep(name)
	if exists {
		return nil
	}
	return fmt.Errorf("error: undefined variable '%s'", name)
}

// Looks for variable in a specific scope
// Returns the type and wether it was found or not
func (fd *FunctionDirectory) FindVariable(scope string, name string) (Variable, bool) {
	funcInfo, ok := fd.Directory[scope]
	if !ok {
		return Variable{}, false
	}
	v, exists := funcInfo.Variables[name]
	return v, exists
}

// Busca una variable desde el scope actual hacia los scopes padres y luego en el global
func (fd *FunctionDirectory) FindVariableDeep(name string) (Variable, bool) {
	// Busca en los scopes actuales (desde el más interno hacia el global)
	for i := len(fd.CurrentScope) - 1; i >= 0; i-- {
		scope := fd.CurrentScope[i]
		v, exists := fd.FindVariable(scope, name)
		if exists {
			return v, true
		}
	}

	// Busca al final en el scope global (por si no está en ninguno anterior)
	v, exists := fd.FindVariable("program", name)
	return v, exists
}
