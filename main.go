package main

import (
	"BabyDuckCompiler/builder"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/symbols"
	"BabyDuckCompiler/vm" // ← AGREGAR ESTA IMPORTACIÓN
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

// convertToVMQuadruples convierte cuádruplos del compilador a formato de VM
func convertToVMQuadruples(compilerQuads []quads.Quadruple) []vm.Quadruple {
	vmQuads := make([]vm.Quadruple, len(compilerQuads))

	for i, quad := range compilerQuads {
		vmQuads[i] = vm.Quadruple{
			Operator:     quad.Operator,
			LeftOperand:  quad.LeftOperand,
			RightOperand: quad.RightOperand,
			Result:       quad.Result,
		}
	}

	return vmQuads
}

// executeProgram ejecuta el programa usando la máquina virtual
func executeProgram(typedQuadruples []quads.Quadruple, parser *builder.PureVisitorParser, debug bool) error {
	// Crear la máquina virtual
	virtualMachine := vm.NewVirtualMachine(debug)

	// Convertir cuádruplos al formato de la VM
	vmQuadruples := convertToVMQuadruples(typedQuadruples)

	// Cargar cuádruplos en la VM
	virtualMachine.LoadQuadruples(vmQuadruples)

	// Obtener y cargar tabla de constantes
	constantTable := parser.GetConstantTable()
	if constantTable != nil {
		// Usar el nuevo método GetConstants() de la tabla de constantes
		constants := constantTable.GetConstants()
		virtualMachine.LoadConstants(constants)

		if debug {
			fmt.Printf("Cargadas %d constantes en la VM\n", len(constants))
		}
	} else {
		// Cargar mapa vacío si no hay tabla de constantes
		virtualMachine.LoadConstants(make(map[int]interface{}))
	}

	// Ejecutar el programa
	err := virtualMachine.Execute()
	if err != nil {
		return err
	}

	// Mostrar estado final de memoria si está en modo debug
	if debug {
		virtualMachine.PrintMemoryState()
	}

	return nil
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

	// Usar el parser final
	parser := builder.NewPureVisitorParser(string(sourceCode), debug)

	// Parsear el código
	symbolTable, errors := parser.Parse()

	// Print symbol table
	printSymbolTable(symbolTable)

	// Obtener cuádruplos del parser
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

	// Obtener tabla de constantes del parser
	constantTable := parser.GetConstantTable()
	if constantTable != nil && constantTable.Size() > 0 {
		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("CONSTANT TABLE")
		fmt.Println(strings.Repeat("=", 60))
		if debug {
			constantTable.PrintDetailed()
		} else {
			constantTable.Print()
		}
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

		// Ejecutar con máquina virtual si hay cuádruplos
		if len(typedQuadruples) > 0 {
			if autoExecute {
				fmt.Println("\n🚀 Ejecutando automáticamente...")

				err := executeProgram(typedQuadruples, parser, debug)
				if err != nil {
					fmt.Printf("❌ Error durante la ejecución: %v\n", err)
					os.Exit(1)
				} else {
					fmt.Println("✅ Ejecución completada exitosamente.")
				}

			} else {
				fmt.Print("\n🚀 ¿Ejecutar el programa? (y/n): ")
				var response string
				fmt.Scanln(&response)

				if response == "y" || response == "Y" || response == "yes" || response == "" {
					err := executeProgram(typedQuadruples, parser, debug)
					if err != nil {
						fmt.Printf("❌ Error durante la ejecución: %v\n", err)
						os.Exit(1)
					} else {
						fmt.Println("✅ Ejecución completada exitosamente.")
					}
				}
			}
		} else {
			fmt.Println("⚠️  No hay cuádruplos para ejecutar.")
		}
	}

	fmt.Println("\nFin de ejecución. Revisa tabla de símbolos y cuádruplos arriba.")
}
