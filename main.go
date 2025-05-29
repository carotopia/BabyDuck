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

	// ========== CORRECCIÓN: Usar el parser final ==========
	parser := builder.NewPureVisitorParser(string(sourceCode), debug)

	// ========== CORRECCIÓN 2: Nombre correcto de variable ==========
	symbolTable, errors := parser.Parse()

	// Print symbol table
	printSymbolTable(symbolTable)

	// ========== CORRECCIÓN 3: Obtener cuádruplos del parser puro ==========
	// El nuevo parser ya no tiene DirectoryBuilder, sino que accede directo
	quadruples := parser.GetQuadruples()

	// Convertir []interface{} a []quads.Quadruple
	var typedQuadruples []quads.Quadruple
	for _, q := range quadruples {
		if quad, ok := q.(quads.Quadruple); ok {
			typedQuadruples = append(typedQuadruples, quad)
		}
	}

	// Print quadruples
	printQuadruples(typedQuadruples)

	// ========== CORRECCIÓN 4: Obtener tabla de constantes del parser puro ==========
	constantTable := parser.GetConstantTable()
	if constantTable != nil {
		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("CONSTANT TABLE")
		fmt.Println(strings.Repeat("=", 60))
		constantTable.Print()
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

		// ========== CORRECCIÓN 5: Ejecutar VM con los nuevos datos ==========
		if len(typedQuadruples) > 0 {
			if autoExecute {
				fmt.Println("\n🚀 Ejecutando automáticamente...")
				fmt.Println(strings.Repeat("=", 60))
				fmt.Println("🚀 EJECUTANDO CON MÁQUINA VIRTUAL")
				fmt.Println(strings.Repeat("=", 60))

				// ========== NOTA: Aquí necesitas adaptar tu VM ==========
				// Tu VM probablemente espere un DirectoryBuilder, pero ahora tenemos parser puro
				// Opción 1: Modificar VM para aceptar parser puro
				// Opción 2: Crear adaptador

				// Por ahora, simulamos la ejecución:
				fmt.Println(">>> Simulando ejecución del programa...")
				fmt.Println(">>> (VM necesita ser adaptada para el nuevo parser)")

				// Descomenta y adapta cuando tengas la VM lista:
				// err := vm.ExecuteProgram(typedQuadruples, parser, debug)
				// if err != nil {
				//     fmt.Printf("❌ Error durante la ejecución: %v\n", err)
				// } else {
				//     fmt.Println("✅ Ejecución completada exitosamente.")
				// }

			} else {
				fmt.Print("\n🚀 ¿Ejecutar el programa? (y/n): ")
				var response string
				fmt.Scanln(&response)

				if response == "y" || response == "Y" || response == "yes" || response == "" {
					fmt.Println(strings.Repeat("=", 60))
					fmt.Println("🚀 EJECUTANDO CON MÁQUINA VIRTUAL")
					fmt.Println(strings.Repeat("=", 60))

					fmt.Println(">>> Simulando ejecución del programa...")
					fmt.Println(">>> (VM necesita ser adaptada para el nuevo parser)")

					// Descomenta y adapta cuando tengas la VM lista:
					// err := vm.ExecuteProgram(typedQuadruples, parser, debug)
					// if err != nil {
					//     fmt.Printf("❌ Error durante la ejecución: %v\n", err)
					// } else {
					//     fmt.Println("✅ Ejecución completada exitosamente.")
					// }
				}
			}
		}
	}

	fmt.Println("\nFin de ejecución. Revisa tabla de símbolos y cuádruplos arriba.")
}
