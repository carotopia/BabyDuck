package main

import (
	"BabyDuckCompiler/builder"
	"BabyDuckCompiler/symbols"
	"fmt"
	"os"
)

// printSymbolTable displays the contents of the function directory,
// showing all variables and their details for each scope.
func printSymbolTable(directory *symbols.FunctionDirectory) {
	for scope, variables := range directory.Directory {
		fmt.Printf("Scope (Function Name): %s\n", scope)
		for varName, varDetails := range variables {
			fmt.Printf("Variable: %s, Details: %v\n", varName, varDetails)
		}
		fmt.Println()
	}
}

func main() {
	// Check if a source file was provided
	if len(os.Args) < 2 {
		fmt.Println("Error: source file path required")
		os.Exit(1)
	}

	sourceFile := os.Args[1]

	// Read the source file
	sourceCode, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading source file: %v\n", err)
		os.Exit(1)
	}

	// Parse the source code
	newParser := builder.NewParser(string(sourceCode))
	symbolTable, errors := newParser.Parse()

	fmt.Println("Function Directory built successfully")
	printSymbolTable(symbolTable)

	if len(errors) > 0 {
		fmt.Println("Compilation errors:")
		for _, err := range errors {
			fmt.Printf("- %s\n", err)
		}
	}
}
