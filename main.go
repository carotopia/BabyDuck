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
	for scope, funcInfo := range directory.Directory {
		fmt.Printf("Scope (Function Name): %s\n", scope)

		if len(funcInfo.Params) > 0 {
			fmt.Println("  Parameters:")
			for _, param := range funcInfo.Params {
				fmt.Printf("    - Name: (unknown), Type: %s\n", param.Type) // Si guardas nombre, cámbialo aquí
			}
		}

		if len(funcInfo.Variables) > 0 {
			fmt.Println("  Variables:")
			for varName, varDetails := range funcInfo.Variables {
				fmt.Printf("    - %s: Type=%s, Address=%d\n",
					varName, varDetails.Type, varDetails.MemoryAddress)
			}
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
