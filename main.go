package main

import (
	"BabyDuckCompiler/builder"
	"BabyDuckCompiler/quads"
	"BabyDuckCompiler/vm"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("🐣 BabyDuck Compiler - Terminal  🐣")
	fmt.Println(strings.Repeat("=", 50))

	// Verificar argumentos de línea de comandos
	if len(os.Args) < 2 {
		fmt.Println("Error: Debes proporcionar un archivo .bd para compilar")
		fmt.Println("Uso: go run main.go archivo.bd [debug]")
		fmt.Println("Ejemplo: go run main.go programa.bd")
		fmt.Println("Ejemplo con debug: go run main.go programa.bd debug")
		os.Exit(1)
	}

	filename := os.Args[1]
	debug := false

	// Verificar si se proporcionó el flag de debug
	if len(os.Args) > 2 && strings.ToLower(os.Args[2]) == "debug" {
		debug = true
		fmt.Println("🐣 Modo debug activado")
	}

	// Leer el archivo
	sourceCode, err := readFile(filename)
	if err != nil {
		fmt.Printf("🐣 Error al leer el archivo '%s': %v\n", filename, err)
		os.Exit(1)
	}

	// Compilar y ejecutar
	compileAndRun(sourceCode, debug)
}

// readFile lee el contenido de un archivo y devuelve su contenido como string
func readFile(filename string) (string, error) {
	// Verificar si el archivo existe
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("el archivo '%s' no existe", filename)
	}

	// Leer el contenido del archivo
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("no se pudo leer el archivo: %v", err)
	}

	return string(content), nil
}

// compileAndRun compila y ejecuta el código fuente proporcionado
func compileAndRun(sourceCode string, debug bool) {
	fmt.Println("🌟 COMPILANDO...")

	parser := builder.NewPureVisitorParser(sourceCode, debug)

	symbolTable, errors := parser.Parse()

	if len(errors) > 0 {
		fmt.Println("🍓 ERRORES DE COMPILACIÓN:")
		for i, err := range errors {
			fmt.Printf("%d. %s\n", i+1, err)
		}
		fmt.Println("🍓 Compilación fallida. No se puede ejecutar el programa.")
		return
	}

	fmt.Println("✅ Compilación exitosa!")

	// Mostrar tabla de símbolos si está en debug
	if debug && symbolTable != nil {
		fmt.Println("\n🍰 TABLA DE SÍMBOLOS:")
		for scope, funcInfo := range symbolTable.Directory {
			fmt.Printf("  Scope: %s\n", scope)
			if len(funcInfo.Variables) > 0 {
				for varName, varDetails := range funcInfo.Variables {
					fmt.Printf("    %s (%s) -> %d\n",
						varName, varDetails.Type, varDetails.MemoryAddress)
				}
			}
		}
	}

	// Obtener cuádruplos
	quadruples := parser.GetQuadruples()
	var builderQuadruples []quads.Quadruple
	for _, q := range quadruples {
		if quad, ok := q.(quads.Quadruple); ok {
			builderQuadruples = append(builderQuadruples, quad)
		}
	}

	if len(builderQuadruples) == 0 {
		fmt.Println("🍓 No hay código válido para ejecutar")
		return
	}

	// Mostrar cuádruplos si está en debug
	if debug {
		fmt.Println("\n💫 CUÁDRUPLOS GENERADOS:")
		for i, quad := range builderQuadruples {
			fmt.Printf("%3d: %-10s %-12v %-12v %-12v\n",
				i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result)
		}
	}

	fmt.Println("\n⚡ EJECUTANDO PROGRAMA:")
	fmt.Println(strings.Repeat("-", 50))

	// Crear VM y configurar para terminal
	virtualMachine := vm.NewVirtualMachine(debug)
	virtualMachine.Reset()

	// Convertir cuádruplos al formato VM
	vmQuadruples := make([]vm.Quadruple, len(builderQuadruples))
	for i, quad := range builderQuadruples {
		vmQuadruples[i] = vm.Quadruple{
			Operator:     quad.Operator,
			LeftOperand:  quad.LeftOperand,
			RightOperand: quad.RightOperand,
			Result:       quad.Result,
		}
	}

	// Cargar programa
	virtualMachine.LoadQuadruples(vmQuadruples)

	// Cargar constantes
	constantTable := parser.GetConstantTable()
	if constantTable != nil {
		constants := constantTable.GetConstants()
		virtualMachine.LoadConstants(constants)
	} else {
		virtualMachine.LoadConstants(make(map[int]interface{}))
	}

	// Ejecutar
	err := virtualMachine.Execute()

	fmt.Println(strings.Repeat("-", 50))

	if err != nil {
		fmt.Printf("🍓 Error de ejecución: %v\n", err)
	} else {
		fmt.Println("✅ Programa ejecutado correctamente!")
	}

	// Mostrar estado de memoria si está en debug
	if debug {
		virtualMachine.PrintMemoryState()
	}
}
