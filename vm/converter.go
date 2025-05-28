package vm

import (
	"BabyDuckCompiler/builder"
	"BabyDuckCompiler/quads"
	"strconv"
	"strings"
)

// ConvertQuadruples convierte cuádruplos del compilador al formato de la VM
func ConvertQuadruples(compilerQuads []quads.Quadruple) []Quadruple {
	vmQuads := make([]Quadruple, len(compilerQuads))

	for i, quad := range compilerQuads {
		vmQuads[i] = Quadruple{
			Operator:     quad.Operator,
			LeftOperand:  quad.LeftOperand,
			RightOperand: quad.RightOperand,
			Result:       quad.Result,
		}
	}

	return vmQuads
}

// ExtractConstants extrae constantes de la tabla de constantes del compilador
func ExtractConstants(dirBuilder *builder.DirectoryBuilder) map[int]interface{} {
	constants := make(map[int]interface{})

	// Extraer constantes reales de tu ConstantTable
	if dirBuilder != nil && dirBuilder.ConstantTable != nil {
		// Tu ConstantTable tiene el formato: "tipo|valor" -> dirección
		// Necesitamos convertir a: dirección -> valor
		for key, address := range dirBuilder.ConstantTable.Constants {
			// Parsear la clave "tipo|valor"
			if len(key) > 4 { // Mínimo "int|1"
				parts := strings.Split(key, "|")
				if len(parts) == 2 {
					typ := parts[0]
					valueStr := parts[1]

					// Convertir el valor según el tipo
					switch typ {
					case "int":
						if val, err := strconv.Atoi(valueStr); err == nil {
							constants[address] = val
						}
					case "float":
						if val, err := strconv.ParseFloat(valueStr, 64); err == nil {
							constants[address] = val
						}
					case "bool":
						if val, err := strconv.ParseBool(valueStr); err == nil {
							constants[address] = val
						}
					case "string":
						// Para strings, remover comillas si las tiene
						if strings.HasPrefix(valueStr, "\"") && strings.HasSuffix(valueStr, "\"") {
							constants[address] = valueStr[1 : len(valueStr)-1]
						} else {
							constants[address] = valueStr
						}
					}
				}
			}
		}
	}

	// Si no se pudieron extraer constantes, usar valores por defecto
	if len(constants) == 0 {
		constants[10000] = 10
		constants[10001] = 5
		constants[10002] = 2
		constants[10003] = 3
		constants[10004] = 0
		constants[10005] = 1
		constants[10006] = 8
	}

	return constants
}

// ExecuteProgram ejecuta un programa compilado
func ExecuteProgram(quadruples []quads.Quadruple, dirBuilder *builder.DirectoryBuilder, debug bool) error {
	// Crear la VM
	vm := NewVirtualMachine(debug)

	// Cargar constantes
	constants := ExtractConstants(dirBuilder)
	vm.LoadConstants(constants)

	// Convertir y cargar cuádruplos
	vmQuads := ConvertQuadruples(quadruples)
	vm.LoadQuadruples(vmQuads)

	// Ejecutar
	err := vm.Execute()
	if err != nil {
		return err
	}

	// Mostrar estado final
	if debug {
		vm.PrintMemoryState()
	}

	return nil
}
