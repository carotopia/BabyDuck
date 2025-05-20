package main

import (
	"BabyDuckCompiler/builder"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/symbols"
	"fmt"
	"os"
	"strings"
)

// printSymbolTable displays the contents of the function directory,
// showing all variables and their details for each scope.
func printSymbolTable(directory *symbols.FunctionDirectory) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("SYMBOL TABLE")
	fmt.Println(strings.Repeat("=", 60))

	if len(directory.Directory) == 0 {
		fmt.Println("No functions or variables defined.")
		return
	}

	for scope, funcInfo := range directory.Directory {
		fmt.Printf("\n%s SCOPE: '%s' %s\n",
			strings.Repeat("-", 10),
			scope,
			strings.Repeat("-", 10))

		if len(funcInfo.Params) > 0 {
			fmt.Println("\n  Parameters:")
			fmt.Println("  " + strings.Repeat("-", 30))
			for i, param := range funcInfo.Params {
				fmt.Printf("    %d. Type: %s\n", i+1, param.Type)
			}
		}

		if len(funcInfo.Variables) > 0 {
			fmt.Println("\n  Variables:")
			fmt.Println("  " + strings.Repeat("-", 30))
			fmt.Printf("  %-20s %-10s %-10s\n", "Name", "Type", "Address")
			fmt.Println("  " + strings.Repeat("-", 40))

			for varName, varDetails := range funcInfo.Variables {
				fmt.Printf("  %-20s %-10s %-10d\n",
					varName, varDetails.Type, varDetails.MemoryAddress)
			}
		} else {
			fmt.Println("\n  No local variables defined.")
		}

	}

	fmt.Println("\n" + strings.Repeat("=", 60))
}

// printQuadruples prints the generated quadruples in a formatted table
func printQuadruples(quads []quads.Quadruple) {
	if len(quads) == 0 {
		fmt.Println("No quadruples generated.")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("GENERATED QUADRUPLES")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Printf("%-5s %-12s %-15s %-15s %-15s\n",
		"Idx", "Operator", "Left Operand", "Right Operand", "Result")
	fmt.Println(strings.Repeat("-", 60))

	for i, quad := range quads {
		leftOp := fmt.Sprintf("%v", quad.LeftOperand)
		rightOp := fmt.Sprintf("%v", quad.RightOperand)
		result := fmt.Sprintf("%v", quad.Result)

		// Handle empty values
		if leftOp == "" {
			leftOp = "_"
		}
		if rightOp == "" {
			rightOp = "_"
		}
		if result == "" {
			result = "_"
		}

		fmt.Printf("%4d: %-12s %-15s %-15s %-15s\n",
			i, quad.Operator, leftOp, rightOp, result)
	}

	fmt.Println(strings.Repeat("=", 60) + "\n")
}
func main() {
	// Check if a source file was provided
	if len(os.Args) < 2 {
		fmt.Println("Error: Source file path required")
		fmt.Println("Usage: go run main.go <filename.bdck>")
		os.Exit(1)
	}

	sourceFile := os.Args[1]

	// Read the source file
	sourceCode, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading source file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Compiling %s...\n", sourceFile)

	// Parse the source code
	newParser := builder.NewParser(string(sourceCode))
	symbolTable, errors := newParser.Parse() // Assume Parse now returns quadruples too

	// Print symbol table
	printSymbolTable(symbolTable)

	// Print quadruples if available
	if dirBuilder, ok := newParser.GetDirectoryBuilder(); ok && dirBuilder.QuadQueue != nil {
		printQuadruples(dirBuilder.QuadQueue.GetAll())

	}

	// Print compilation status
	if len(errors) > 0 {
		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("COMPILATION ERRORS")
		fmt.Println(strings.Repeat("=", 60))
		for i, err := range errors {
			fmt.Printf("%d. %s\n", i+1, err)
		}
		fmt.Println("\nCompilation failed with errors.")
		os.Exit(1)
	} else {
		fmt.Println("\nCompilation successful! No errors detected.")
	}
	fmt.Println("\nFin de ejecución. Revisa tabla de símbolos y cuádruplos arriba.")

}
