package main

import (
	"BabyDuckCompiler/builder"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/symbols"
	"BabyDuckCompiler/vm" // <- NUEVO IMPORT
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
		if leftOp == "" || leftOp == "<nil>" {
			leftOp = "_"
		}
		if rightOp == "" || rightOp == "<nil>" {
			rightOp = "_"
		}
		if result == "" || result == "<nil>" {
			result = "_"
		}

		fmt.Printf("%4d: %-12s %-15s %-15s %-15s\n",
			i, quad.Operator, leftOp, rightOp, result)
	}

	fmt.Println(strings.Repeat("=", 60) + "\n")
}

// printErrors prints compilation errors in a formatted way
func printErrors(errors []string) {
	if len(errors) == 0 {
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("COMPILATION ERRORS")
	fmt.Println(strings.Repeat("=", 60))

	for i, err := range errors {
		fmt.Printf("%d. %s\n", i+1, err)
	}

	fmt.Println(strings.Repeat("=", 60))
}

func main() {
	// Check if a source file was provided
	if len(os.Args) < 2 {
		fmt.Println("Error: Source file path required")
		fmt.Println("Usage: go run main.go <filename.bdck> [debug] [--execute]")
		fmt.Println("  debug: optional flag to enable debug output")
		fmt.Println("  --execute: automatically execute after compilation")
		os.Exit(1)
	}

	sourceFile := os.Args[1]

	// Check flags
	debug := false
	autoExecute := false
	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "debug", "-d", "--debug":
			debug = true
		case "--execute", "-e", "--run":
			autoExecute = true
		}
	}

	// Read the source file
	sourceCode, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading source file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Compiling %s...\n", sourceFile)
	if debug {
		fmt.Println("Debug mode enabled")
	}
	fmt.Println(strings.Repeat("=", 60))

	// Create parser with optional debug
	newParser := builder.NewParser(string(sourceCode), debug)

	// Parse the source code
	symbolTable, errors := newParser.Parse()

	// Print symbol table
	printSymbolTable(symbolTable)

	// Get the directory builder to access quadruples
	dirBuilder := newParser.GetDirectoryBuilder()

	// Print quadruples if available
	if dirBuilder != nil && dirBuilder.QuadVisitor != nil {
		// Access quadruples from the visitor
		quadruples := dirBuilder.QuadVisitor.GetQuadruples()
		printQuadruples(quadruples)
	} else {
		fmt.Println("\n⚠️  No quadruples generated or visitor not initialized.")
	}

	// Print constant table if available
	if dirBuilder != nil && dirBuilder.ConstantTable != nil {
		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("CONSTANT TABLE")
		fmt.Println(strings.Repeat("=", 60))
		dirBuilder.ConstantTable.Print()
	}

	// Print compilation errors
	printErrors(errors)

	// Print compilation summary
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("COMPILATION SUMMARY")
	fmt.Println(strings.Repeat("=", 60))

	if len(errors) > 0 {
		fmt.Printf("❌ Compilation failed with %d error(s).\n", len(errors))
		os.Exit(1)
	} else {
		fmt.Println("✅ Compilation successful! No errors detected.")

		// NUEVA SECCIÓN: Ejecutar con VM
		if dirBuilder != nil && dirBuilder.QuadVisitor != nil {
			quadruples := dirBuilder.QuadVisitor.GetQuadruples()

			if autoExecute {
				fmt.Println("\n🚀 Ejecutando automáticamente...")
				err := vm.ExecuteProgram(quadruples, dirBuilder, debug)
				if err != nil {
					fmt.Printf("❌ Error durante la ejecución: %v\n", err)
				} else {
					fmt.Println("✅ Ejecución completada exitosamente.")
				}
			} else {
				fmt.Print("\n🚀 ¿Ejecutar el programa? (y/n): ")
				var response string
				fmt.Scanln(&response)

				if response == "y" || response == "Y" || response == "yes" || response == "" {
					err := vm.ExecuteProgram(quadruples, dirBuilder, debug)
					if err != nil {
						fmt.Printf("❌ Error durante la ejecución: %v\n", err)
					} else {
						fmt.Println("✅ Ejecución completada exitosamente.")
					}
				}
			}
		}
	}

	fmt.Println("\nFin de ejecución. Revisa tabla de símbolos y cuádruplos arriba.")
}
