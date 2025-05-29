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
	if vm.debug {
		fmt.Printf("VM: Cargados %d cuádruplos\n", len(quads))
	}
}

// LoadConstants carga la tabla de constantes
func (vm *VirtualMachine) LoadConstants(constants map[int]interface{}) {
	vm.constants = constants
	if vm.debug {
		fmt.Printf("VM: Cargadas %d constantes\n", len(constants))
	}
}

// getValue obtiene el valor de un operando (puede ser dirección de memoria o constante)
func (vm *VirtualMachine) getValue(operand interface{}) interface{} {
	if operand == nil {
		return nil
	}

	switch v := operand.(type) {
	case int:
		// Primero verificar si es una constante
		if val, exists := vm.constants[v]; exists {
			if vm.debug {
				fmt.Printf("    Obteniendo constante[%d] = %v\n", v, val)
			}
			return val
		}
		// Luego verificar si es una dirección de memoria
		if val, exists := vm.memory[v]; exists {
			if vm.debug {
				fmt.Printf("    Obteniendo memoria[%d] = %v\n", v, val)
			}
			return val
		}
		// Si no existe en ningún lado, devolver el valor directo
		if vm.debug {
			fmt.Printf("    Usando valor directo: %d\n", v)
		}
		return v
	case string:
		// Si es una cadena literal, remover comillas si las tiene
		if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
			cleaned := v[1 : len(v)-1]
			if vm.debug {
				fmt.Printf("    String literal: %s -> %s\n", v, cleaned)
			}
			return cleaned
		}
		if vm.debug {
			fmt.Printf("    String value: %s\n", v)
		}
		return v
	case float64:
		if vm.debug {
			fmt.Printf("    Float value: %f\n", v)
		}
		return v
	default:
		if vm.debug {
			fmt.Printf("    Other value (%T): %v\n", v, v)
		}
		return v
	}
}

// setValue asigna un valor a una dirección de memoria
func (vm *VirtualMachine) setValue(address interface{}, value interface{}) {
	if addr, ok := address.(int); ok {
		vm.memory[addr] = value
		if vm.debug {
			fmt.Printf("    Asignando memoria[%d] = %v\n", addr, value)
		}
	} else if vm.debug {
		fmt.Printf("    ⚠️  Dirección inválida para asignación: %v (%T)\n", address, address)
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
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int", value)
	}
}

// toFloat convierte un valor a float64 si es posible
func (vm *VirtualMachine) toFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case string:
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			return val, nil
		}
		return 0, fmt.Errorf("cannot convert string '%s' to float", v)
	default:
		return 0, fmt.Errorf("cannot convert %T to float", value)
	}
}

// toBool convierte un valor a booleano
func (vm *VirtualMachine) toBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0.0
	case string:
		return v != "" && v != "false" && v != "0"
	default:
		return false
	}
}

// Execute ejecuta los cuádruplos
func (vm *VirtualMachine) Execute() error {
	if len(vm.quadruples) == 0 {
		fmt.Println("⚠️  No hay cuádruplos para ejecutar")
		return nil
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🚀 EJECUTANDO CON MÁQUINA VIRTUAL")
	fmt.Println(strings.Repeat("=", 60))

	for vm.pc < len(vm.quadruples) {
		if vm.pc < 0 || vm.pc >= len(vm.quadruples) {
			return fmt.Errorf("PC fuera de rango: %d (máximo: %d)", vm.pc, len(vm.quadruples)-1)
		}

		quad := vm.quadruples[vm.pc]

		if vm.debug {
			fmt.Printf("\nPC: %d - Ejecutando: %s %v %v %v\n",
				vm.pc, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result)
		}

		err := vm.executeQuadruple(quad)
		if err != nil {
			return fmt.Errorf("error en PC %d: %v", vm.pc, err)
		}

		vm.pc++
	}

	fmt.Println("\n✅ Ejecución completada exitosamente.")
	return nil
}

// executeQuadruple ejecuta un cuádruple individual
func (vm *VirtualMachine) executeQuadruple(quad Quadruple) error {
	switch quad.Operator {
	case "=":
		// Asignación
		value := vm.getValue(quad.LeftOperand)
		vm.setValue(quad.Result, value)

	case "+":
		// Suma
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toInt(left)
		rightNum, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			// Intentar como float
			leftFloat, ferr1 := vm.toFloat(left)
			rightFloat, ferr2 := vm.toFloat(right)
			if ferr1 != nil || ferr2 != nil {
				return fmt.Errorf("error en suma: %v + %v", left, right)
			}
			result := leftFloat + rightFloat
			vm.setValue(quad.Result, result)
		} else {
			result := leftNum + rightNum
			vm.setValue(quad.Result, result)
		}

	case "-":
		// Resta
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toInt(left)
		rightNum, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			leftFloat, ferr1 := vm.toFloat(left)
			rightFloat, ferr2 := vm.toFloat(right)
			if ferr1 != nil || ferr2 != nil {
				return fmt.Errorf("error en resta: %v - %v", left, right)
			}
			result := leftFloat - rightFloat
			vm.setValue(quad.Result, result)
		} else {
			result := leftNum - rightNum
			vm.setValue(quad.Result, result)
		}

	case "*":
		// Multiplicación
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toInt(left)
		rightNum, err2 := vm.toInt(right)

		if err1 != nil || err2 != nil {
			leftFloat, ferr1 := vm.toFloat(left)
			rightFloat, ferr2 := vm.toFloat(right)
			if ferr1 != nil || ferr2 != nil {
				return fmt.Errorf("error en multiplicación: %v * %v", left, right)
			}
			result := leftFloat * rightFloat
			vm.setValue(quad.Result, result)
		} else {
			result := leftNum * rightNum
			vm.setValue(quad.Result, result)
		}

	case "/":
		// División
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		// Siempre usar float para división
		leftFloat, err1 := vm.toFloat(left)
		rightFloat, err2 := vm.toFloat(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en división: %v / %v", left, right)
		}

		if rightFloat == 0 {
			return fmt.Errorf("error: división por cero")
		}

		result := leftFloat / rightFloat
		vm.setValue(quad.Result, result)

	case "print", "PRINT":
		// Imprimir - buscar valor en todos los operandos posibles
		var value interface{}

		if quad.LeftOperand != nil {
			value = vm.getValue(quad.LeftOperand)
		} else if quad.RightOperand != nil {
			value = vm.getValue(quad.RightOperand)
		} else if quad.Result != nil {
			value = vm.getValue(quad.Result)
		}

		if value != nil {
			fmt.Printf(">>> %v\n", value)
		} else {
			fmt.Println(">>> <valor nulo>")
		}

	case ">":
		// Mayor que
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toFloat(left)
		rightNum, err2 := vm.toFloat(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en comparación: %v > %v", left, right)
		}

		result := leftNum > rightNum
		vm.setValue(quad.Result, result)

	case "<":
		// Menor que
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toFloat(left)
		rightNum, err2 := vm.toFloat(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en comparación: %v < %v", left, right)
		}

		result := leftNum < rightNum
		vm.setValue(quad.Result, result)

	case ">=":
		// Mayor o igual que
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toFloat(left)
		rightNum, err2 := vm.toFloat(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en comparación: %v >= %v", left, right)
		}

		result := leftNum >= rightNum
		vm.setValue(quad.Result, result)

	case "<=":
		// Menor o igual que
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toFloat(left)
		rightNum, err2 := vm.toFloat(right)

		if err1 != nil || err2 != nil {
			return fmt.Errorf("error en comparación: %v <= %v", left, right)
		}

		result := leftNum <= rightNum
		vm.setValue(quad.Result, result)

	case "==":
		// Igual que
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		// Intentar comparación numérica primero
		leftNum, err1 := vm.toFloat(left)
		rightNum, err2 := vm.toFloat(right)

		if err1 == nil && err2 == nil {
			result := leftNum == rightNum
			vm.setValue(quad.Result, result)
		} else {
			// Comparación directa
			result := left == right
			vm.setValue(quad.Result, result)
		}

	case "!=":
		// Diferente
		left := vm.getValue(quad.LeftOperand)
		right := vm.getValue(quad.RightOperand)

		leftNum, err1 := vm.toFloat(left)
		rightNum, err2 := vm.toFloat(right)

		if err1 == nil && err2 == nil {
			result := leftNum != rightNum
			vm.setValue(quad.Result, result)
		} else {
			result := left != right
			vm.setValue(quad.Result, result)
		}

	case "GOTOF":
		// Salto condicional (Go To False)
		condition := vm.getValue(quad.LeftOperand)
		conditionBool := vm.toBool(condition)

		if vm.debug {
			fmt.Printf("  GOTOF - Condición: %v (%t)", condition, conditionBool)
		}

		if !conditionBool { // Si la condición es FALSA
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

	case "FUNC", "ENDFUNC", "PARAM", "ERA", "PARAMETER", "GOSUB", "RET":
		// Operaciones de función - para implementar más tarde
		if vm.debug {
			fmt.Printf("  Operación de función (no implementada): %s\n", quad.Operator)
		}

	default:
		if vm.debug {
			fmt.Printf("  ⚠️  Operación no reconocida: %s\n", quad.Operator)
		}
	}

	return nil
}

// PrintMemoryState imprime el estado actual de la memoria
func (vm *VirtualMachine) PrintMemoryState() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("ESTADO FINAL DE MEMORIA")
	fmt.Println(strings.Repeat("=", 60))

	if len(vm.memory) == 0 {
		fmt.Println("Memoria vacía")
	} else {
		for addr, value := range vm.memory {
			fmt.Printf("Memoria[%d] = %v (%T)\n", addr, value, value)
		}
	}

	if len(vm.constants) > 0 {
		fmt.Println("\nConstantes:")
		for addr, value := range vm.constants {
			fmt.Printf("Constante[%d] = %v (%T)\n", addr, value, value)
		}
	}

	fmt.Println(strings.Repeat("=", 60))
}
