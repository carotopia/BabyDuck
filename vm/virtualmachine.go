package vm

import (
	"fmt"
	"strconv"
	"strings"
)

// Quadruple representa un cuádruple para la VM
type Quadruple struct {
	Operator     string
	LeftOperand  interface{}
	RightOperand interface{}
	Result       interface{}
}

// VirtualMachine representa la máquina virtual
type VirtualMachine struct {
	quadruples []Quadruple
	memory     map[int]interface{} // Memoria global por direcciones
	constants  map[int]interface{} // Tabla de constantes
	pc         int                 // Program Counter
	debug      bool                // Modo debug
}

// NewVirtualMachine crea una nueva instancia de la máquina virtual
func NewVirtualMachine(debug bool) *VirtualMachine {
	return &VirtualMachine{
		quadruples: make([]Quadruple, 0),
		memory:     make(map[int]interface{}),
		constants:  make(map[int]interface{}),
		pc:         0,
		debug:      debug,
	}
}

// LoadQuadruples carga los cuádruplos en la máquina virtual
func (vm *VirtualMachine) LoadQuadruples(quads []Quadruple) {
	vm.quadruples = quads
}

// LoadConstants carga la tabla de constantes
func (vm *VirtualMachine) LoadConstants(constants map[int]interface{}) {
	vm.constants = constants
}

// getValue obtiene el valor de un operando (puede ser dirección de memoria o constante)
func (vm *VirtualMachine) getValue(operand interface{}) interface{} {
	if operand == nil {
		return nil
	}

	switch v := operand.(type) {
	case int:
		// Si es una dirección de memoria
		if val, exists := vm.memory[v]; exists {
			return val
		}
		// Si es una constante
		if val, exists := vm.constants[v]; exists {
			return val
		}
		// Si no existe, devolver el valor directo
		return v
	case string:
		// Si es una cadena literal, remover comillas si las tiene
		if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
			return v[1 : len(v)-1]
		}
		return v
	default:
		return v
	}
}

// setValue asigna un valor a una dirección de memoria
func (vm *VirtualMachine) setValue(address interface{}, value interface{}) {
	if addr, ok := address.(int); ok {
		vm.memory[addr] = value
	}
}

// toInt convierte un valor a entero si es posible
func (vm *VirtualMachine) toInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		if val, err := strconv.Atoi(v); err == nil {
			return val, nil
		}
		return 0, fmt.Errorf("cannot convert string '%s' to int", v)
	default:
		return 0, fmt.Errorf("cannot convert %T to int", value)
	}
}

// Execute ejecuta los cuádruplos
func (vm *VirtualMachine) Execute() error {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🚀 EJECUTANDO CON MÁQUINA VIRTUAL")
	fmt.Println(strings.Repeat("=", 60))

	// Preprocesar cuádruplos para arreglar orden incorrecto del compilador
	vm.preprocessQuadruples()

	for vm.pc < len(vm.quadruples) {
		quad := vm.quadruples[vm.pc]

		if vm.debug {
			fmt.Printf("PC: %d - Ejecutando: %s %v %v %v\n",
				vm.pc, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result)
		}

		err := vm.executeQuadruple(quad)
		if err != nil {
			return fmt.Errorf("error en PC %d: %v", vm.pc, err)
		}

		vm.pc++
	}

	return nil
}

// preprocessQuadruples arregla el orden incorrecto de los cuádruplos del compilador
func (vm *VirtualMachine) preprocessQuadruples() {
	newQuads := make([]Quadruple, 0, len(vm.quadruples))

	for i := 0; i < len(vm.quadruples); i++ {
		quad := vm.quadruples[i]

		// Detectar patrón incorrecto de if/else
		if quad.Operator == ">" || quad.Operator == "<" || quad.Operator == "==" || quad.Operator == "!=" {
			// Buscar si hay un GOTOF mal ubicado después
			gotofIndex := -1
			for j := i + 1; j < len(vm.quadruples) && j < i+10; j++ {
				if vm.quadruples[j].Operator == "GOTOF" {
					gotofIndex = j
					break
				}
			}

			if gotofIndex != -1 {
				// Patrón detectado: comparación seguida de prints y GOTOF al final
				// Reordenar para que funcione correctamente

				// 1. Añadir la comparación
				newQuads = append(newQuads, quad)

				// 2. Añadir GOTOF inmediatamente después de la comparación
				gotofQuad := vm.quadruples[gotofIndex]
				// Calcular la dirección correcta para saltar al else
				elseIndex := -1
				eeeIndex := -1

				// Buscar EEE y la instrucción después
				for j := i + 1; j < gotofIndex; j++ {
					if vm.quadruples[j].Operator == "EEE" {
						eeeIndex = j
						elseIndex = j + 1
						break
					}
				}

				if elseIndex != -1 {
					// Crear GOTOF que salte al else
					fixedGotof := Quadruple{
						Operator:     "GOTOF",
						LeftOperand:  gotofQuad.LeftOperand,
						RightOperand: gotofQuad.RightOperand,
						Result:       len(newQuads) + 2, // Saltar después del if
					}
					newQuads = append(newQuads, fixedGotof)

					// 3. Añadir solo el print del if (sin el else)
					for j := i + 1; j < eeeIndex; j++ {
						if vm.quadruples[j].Operator != "EEE" {
							newQuads = append(newQuads, vm.quadruples[j])
						}
					}

					// 4. Añadir el else
					if elseIndex < len(vm.quadruples) {
						newQuads = append(newQuads, vm.quadruples[elseIndex])
					}

					// Saltar todos los cuádruplos que ya procesamos
					i = gotofIndex
					continue
				}
			}
		}

		// Si no es un patrón de if/else, añadir normalmente
		if quad.Operator != "GOTOF" || (quad.Result != nil && quad.Result != "<nil>") {
			newQuads = append(newQuads, quad)
		}
	}

	// Actualizar los cuádruplos
	vm.quadruples = newQuads

	if vm.debug {
		fmt.Println("=== CUÁDRUPLOS DESPUÉS DEL PREPROCESAMIENTO ===")
		for i, quad := range vm.quadruples {
			fmt.Printf("%d: %s %v %v %v\n", i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result)
		}
		fmt.Println("===========================================")
	}
}

// executeQuadruple ejecuta un cuádruple individual
func (vm *VirtualMachine) executeQuadruple(quad Quadruple) error {
	switch quad.Operator {
	case "=":
		// Asignación
		value := vm.getValue(quad.LeftOperand)
		vm.setValue(quad.Result, value)
		if vm.debug {
			fmt.Printf("  Asignación: %v -> memoria[%v] = %v\n", quad.LeftOperand, quad.Result, value)
		}

	case "+":
		// Suma
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftInt, err1 := vm.toInt(left)
		rightInt, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en suma: %v + %v", left, right)
		}

		result := leftInt + rightInt
		vm.setValue(quad.Result, result)
		if vm.debug {
			fmt.Printf("  Suma: %d + %d = %d -> memoria[%v]\n", leftInt, rightInt, result, quad.Result)
		}

	case "-":
		// Resta
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftInt, err1 := vm.toInt(left)
		rightInt, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en resta: %v - %v", left, right)
		}

		result := leftInt - rightInt
		vm.setValue(quad.Result, result)
		if vm.debug {
			fmt.Printf("  Resta: %d - %d = %d -> memoria[%v]\n", leftInt, rightInt, result, quad.Result)
		}

	case "*":
		// Multiplicación
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftInt, err1 := vm.toInt(left)
		rightInt, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en multiplicación: %v * %v", left, right)
		}

		result := leftInt * rightInt
		vm.setValue(quad.Result, result)
		if vm.debug {
			fmt.Printf("  Multiplicación: %d * %d = %d -> memoria[%v]\n", leftInt, rightInt, result, quad.Result)
		}

	case "/":
		// División
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftInt, err1 := vm.toInt(left)
		rightInt, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en división: %v / %v", left, right)
		}

		if rightInt == 0 {
			return fmt.Errorf("error: división por cero")
		}

		result := leftInt / rightInt
		vm.setValue(quad.Result, result)
		if vm.debug {
			fmt.Printf("  División: %d / %d = %d -> memoria[%v]\n", leftInt, rightInt, result, quad.Result)
		}

	case "print":
		// Imprimir
		value := vm.getValue(quad.LeftOperand)
		fmt.Printf(">>> %v\n", value)

	case ">":
		// Mayor que
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftInt, err1 := vm.toInt(left)
		rightInt, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en comparación: %v > %v", left, right)
		}

		result := leftInt > rightInt
		vm.setValue(quad.Result, result)
		if vm.debug {
			fmt.Printf("  Comparación: %d > %d = %t -> memoria[%v]\n", leftInt, rightInt, result, quad.Result)
		}

	case "<":
		// Menor que
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftInt, err1 := vm.toInt(left)
		rightInt, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en comparación: %v < %v", left, right)
		}

		result := leftInt < rightInt
		vm.setValue(quad.Result, result)
		if vm.debug {
			fmt.Printf("  Comparación: %d < %d = %t -> memoria[%v]\n", leftInt, rightInt, result, quad.Result)
		}

	case "GOTOF":
		// Salto condicional (Go To False)
		condition := vm.getValue(quad.LeftOperand)
		if vm.debug {
			fmt.Printf("  GOTOF - Condición: %v", condition)
		}

		// Verificar si tenemos una dirección de salto válida
		if quad.Result == nil || quad.Result == "<nil>" {
			if vm.debug {
				fmt.Printf(" (sin dirección de salto válida - ignorando)\n")
			}
			return nil
		}

		if condition == false {
			if jumpAddr, ok := quad.Result.(int); ok {
				vm.pc = jumpAddr - 1 // -1 porque se incrementará al final del loop
				if vm.debug {
					fmt.Printf(" -> Saltando a PC: %d\n", jumpAddr)
				}
			}
		} else {
			if vm.debug {
				fmt.Printf(" -> Continúa secuencial\n")
			}
		}

	case "GOTO":
		// Salto incondicional
		if jumpAddr, ok := quad.Result.(int); ok {
			vm.pc = jumpAddr - 1 // -1 porque se incrementará al final del loop
			if vm.debug {
				fmt.Printf("  Salto incondicional a PC: %d\n", jumpAddr)
			}
		}

	case "FUNC", "ENDFUNC", "PARAM", "ERA", "PARAMETER", "GOSUB":
		// Operaciones de función - implementar más tarde
		if vm.debug {
			fmt.Printf("  Operación de función (no implementada): %s\n", quad.Operator)
		}

	case "EEE":
		// Operación especial - por ahora ignorar
		if vm.debug {
			fmt.Printf("  Marcador EEE (fin de bloque)\n")
		}

	default:
		if vm.debug {
			fmt.Printf("  Operación no implementada: %s\n", quad.Operator)
		}
	}

	return nil
}

// PrintMemoryState imprime el estado actual de la memoria
func (vm *VirtualMachine) PrintMemoryState() {
	fmt.Println("\n=== ESTADO FINAL DE MEMORIA ===")
	if len(vm.memory) == 0 {
		fmt.Println("Memoria vacía")
		return
	}

	for addr, value := range vm.memory {
		fmt.Printf("Memoria[%d] = %v\n", addr, value)
	}
	fmt.Println("==============================")
}
