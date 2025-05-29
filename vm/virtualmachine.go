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

	// Stack para manejo de funciones (simplificado)
	callStack  []CallFrame   // Stack de contextos de llamada
	paramStack []interface{} // Stack de parámetros temporales

	// Información de funciones
	functionTable map[string]int // nombre -> dirección de inicio
}

// CallFrame representa un contexto de llamada a función
type CallFrame struct {
	FunctionName  string              // Nombre de la función
	ReturnAddress int                 // Dirección de retorno
	LocalMemory   map[int]interface{} // Memoria local de la función
	Parameters    []interface{}       // Parámetros de la función
}

// NewVirtualMachine crea una nueva instancia de la máquina virtual
func NewVirtualMachine(debug bool) *VirtualMachine {
	return &VirtualMachine{
		quadruples:    make([]Quadruple, 0),
		memory:        make(map[int]interface{}),
		constants:     make(map[int]interface{}),
		pc:            0,
		debug:         debug,
		callStack:     make([]CallFrame, 0),
		paramStack:    make([]interface{}, 0),
		functionTable: make(map[string]int),
	}
}

// LoadQuadruples carga los cuádruplos en la máquina virtual
func (vm *VirtualMachine) LoadQuadruples(quads []Quadruple) {
	vm.quadruples = quads

	// Preprocesar para encontrar funciones
	vm.preprocessFunctions()

	if vm.debug {
		fmt.Printf("VM: Cargados %d cuádruplos\n", len(quads))
		if len(vm.functionTable) > 0 {
			fmt.Printf("VM: Encontradas %d funciones\n", len(vm.functionTable))
		}
	}
}

// preprocessFunctions busca y registra todas las funciones antes de la ejecución
func (vm *VirtualMachine) preprocessFunctions() {
	for i, quad := range vm.quadruples {
		if quad.Operator == "FUNC" {
			if funcName, ok := quad.LeftOperand.(string); ok {
				vm.functionTable[funcName] = i
				if vm.debug {
					fmt.Printf("VM: Registrando función '%s' en PC: %d\n", funcName, i)
				}
			}
		}
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

		// Verificar memoria local si hay contexto de función activo
		if len(vm.callStack) > 0 {
			currentFrame := &vm.callStack[len(vm.callStack)-1]
			if val, exists := currentFrame.LocalMemory[v]; exists {
				if vm.debug {
					fmt.Printf("    Obteniendo memoria local[%d] = %v (función: %s)\n", v, val, currentFrame.FunctionName)
				}
				return val
			}
		}

		// Verificar memoria global
		if val, exists := vm.memory[v]; exists {
			if vm.debug {
				fmt.Printf("    Obteniendo memoria global[%d] = %v\n", v, val)
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

// setValue asigna un valor a una dirección de memoria (global o local)
func (vm *VirtualMachine) setValue(address interface{}, value interface{}) {
	if addr, ok := address.(int); ok {
		// Si hay contexto de función activo, verificar si es memoria local
		if len(vm.callStack) > 0 {
			currentFrame := &vm.callStack[len(vm.callStack)-1]
			// Si es una dirección local (rango 1000-4999 para parámetros y locales)
			if addr >= 1000 && addr < 5000 {
				currentFrame.LocalMemory[addr] = value
				if vm.debug {
					fmt.Printf("    Asignando memoria local[%d] = %v (función: %s)\n", addr, value, currentFrame.FunctionName)
				}
				return
			}
		}

		// Asignar a memoria global
		vm.memory[addr] = value
		if vm.debug {
			fmt.Printf("    Asignando memoria global[%d] = %v\n", addr, value)
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

	// Inicializar PC en 0
	vm.pc = 0

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
// executeQuadruple ejecuta un cuádruple individual
func (vm *VirtualMachine) executeQuadruple(quad Quadruple) error {
	op := quad.Operator

	// ==================== OPERACIONES DE FUNCIÓN (CON IF) ====================
	if op == "ERA" {
		if vm.debug {
			funcName := "desconocida"
			if name, ok := quad.LeftOperand.(string); ok {
				funcName = name
			}
			fmt.Printf("  ERA: Preparando espacio para función '%s'\n", funcName)
		}
		return nil
	}

	if op == "PARAMETER" {
		value := vm.getValue(quad.LeftOperand)
		vm.paramStack = append(vm.paramStack, value)
		if vm.debug {
			fmt.Printf("  PARAMETER: Guardando parámetro %v (total: %d)\n", value, len(vm.paramStack))
		}
		return nil
	}

	if op == "GOSUB" {
		var funcAddr int
		var funcName string

		// Determinar función a llamar - ADAPTADO para tu formato específico
		if name, ok := quad.LeftOperand.(string); ok {
			funcName = name
			// Buscar la función en la tabla
			if addr, exists := vm.functionTable[name]; exists {
				funcAddr = addr
			} else {
				return fmt.Errorf("función '%s' no encontrada", name)
			}
		} else if addr, ok := quad.Result.(int); ok {
			// El Result contiene la dirección de la función
			funcAddr = addr
			// Buscar nombre por dirección
			for name, address := range vm.functionTable {
				if address == addr {
					funcName = name
					break
				}
			}
			if funcName == "" {
				funcName = fmt.Sprintf("func_at_%d", addr)
			}
		} else {
			// Intentar como formato específico de tu compilador
			// Si LeftOperand es string y Result es int
			if leftStr, ok := quad.LeftOperand.(string); ok {
				funcName = leftStr
				if resultInt, ok := quad.Result.(int); ok {
					funcAddr = resultInt
				} else {
					return fmt.Errorf("formato de GOSUB no válido: %v", quad)
				}
			} else {
				return fmt.Errorf("dirección de función inválida en GOSUB: %v", quad)
			}
		}

		if vm.debug {
			fmt.Printf("  GOSUB: Intentando llamar función '%s' en PC: %d\n", funcName, funcAddr)
		}

		// Crear frame de ejecución
		frame := CallFrame{
			FunctionName:  funcName,
			ReturnAddress: vm.pc + 1,
			LocalMemory:   make(map[int]interface{}),
			Parameters:    make([]interface{}, len(vm.paramStack)),
		}

		// Copiar parámetros
		copy(frame.Parameters, vm.paramStack)

		// Asignar parámetros a memoria local
		for i, param := range vm.paramStack {
			paramAddr := 1000 + i
			frame.LocalMemory[paramAddr] = param
			if vm.debug {
				fmt.Printf("    Parámetro[%d] = %v -> memoria local[%d]\n", i, param, paramAddr)
			}
		}

		// Limpiar stack de parámetros
		vm.paramStack = vm.paramStack[:0]

		// Agregar frame al call stack
		vm.callStack = append(vm.callStack, frame)

		// Saltar a la función
		vm.pc = funcAddr - 1
		if vm.debug {
			fmt.Printf("  GOSUB: Llamando función '%s' en PC: %d\n", funcName, funcAddr)
		}
		return nil
	}

	if op == "FUNC" {
		if vm.debug {
			if funcName, ok := quad.LeftOperand.(string); ok {
				fmt.Printf("  FUNC: Definición de función '%s'\n", funcName)
			}
		}
		return nil
	}

	if op == "PARAM" {
		if vm.debug {
			fmt.Printf("  PARAM: Declaración de parámetro (saltando)\n")
		}
		return nil
	}

	if op == "ENDFUNC" {
		if len(vm.callStack) > 0 {
			frame := vm.callStack[len(vm.callStack)-1]
			vm.callStack = vm.callStack[:len(vm.callStack)-1]

			vm.pc = frame.ReturnAddress - 1
			if vm.debug {
				fmt.Printf("  ENDFUNC: Retornando de función '%s' a PC: %d\n", frame.FunctionName, frame.ReturnAddress)
			}
		} else {
			if vm.debug {
				fmt.Printf("  ENDFUNC: Fin de función principal\n")
			}
		}
		return nil
	}

	if op == "RET" {
		if len(vm.callStack) > 0 {
			frame := vm.callStack[len(vm.callStack)-1]
			vm.callStack = vm.callStack[:len(vm.callStack)-1]

			if quad.LeftOperand != nil {
				returnValue := vm.getValue(quad.LeftOperand)
				if quad.Result != nil {
					vm.setValue(quad.Result, returnValue)
				}
				if vm.debug {
					fmt.Printf("  RET: Retornando valor %v de función '%s'\n", returnValue, frame.FunctionName)
				}
			}

			vm.pc = frame.ReturnAddress - 1
			if vm.debug {
				fmt.Printf("  RET: Retornando de función '%s' a PC: %d\n", frame.FunctionName, frame.ReturnAddress)
			}
		}
		return nil
	}

	// ==================== OPERACIONES NORMALES (CON SWITCH) ====================
	switch op {
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
		var values []interface{}

		// Recopilar todos los valores no nulos
		if quad.LeftOperand != nil {
			val := vm.getValue(quad.LeftOperand)
			if val != nil {
				values = append(values, val)
			}
		}
		if quad.RightOperand != nil {
			val := vm.getValue(quad.RightOperand)
			if val != nil {
				values = append(values, val)
			}
		}
		if quad.Result != nil {
			val := vm.getValue(quad.Result)
			if val != nil {
				values = append(values, val)
			}
		}

		// Imprimir todos los valores
		if len(values) > 0 {
			fmt.Print(">>> ")
			for i, value := range values {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(value)
			}
			fmt.Println()
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
		// Mayor o equal que
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

	default:
		if vm.debug {
			fmt.Printf("  ⚠️  Operación no reconocida: %s\n", op)
		}
		return fmt.Errorf("operación no implementada: %s", op)
	}

	return nil
}

// PrintMemoryState imprime el estado actual de la memoria
func (vm *VirtualMachine) PrintMemoryState() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("ESTADO FINAL DE MEMORIA")
	fmt.Println(strings.Repeat("=", 60))

	// Memoria global
	fmt.Println("Memoria Global:")
	if len(vm.memory) == 0 {
		fmt.Println("  Memoria global vacía")
	} else {
		for addr, value := range vm.memory {
			fmt.Printf("  Memoria[%d] = %v (%T)\n", addr, value, value)
		}
	}

	// Call stack activo
	if len(vm.callStack) > 0 {
		fmt.Printf("\nCall Stack activo (%d frames):\n", len(vm.callStack))
		for i, frame := range vm.callStack {
			fmt.Printf("  [%d] Función '%s' -> Retorno: PC %d\n", i, frame.FunctionName, frame.ReturnAddress)
			if len(frame.LocalMemory) > 0 {
				fmt.Printf("      Memoria local:\n")
				for addr, value := range frame.LocalMemory {
					fmt.Printf("        Local[%d] = %v (%T)\n", addr, value, value)
				}
			}
			if len(frame.Parameters) > 0 {
				fmt.Printf("      Parámetros: %v\n", frame.Parameters)
			}
		}
	}

	// Constantes
	if len(vm.constants) > 0 {
		fmt.Println("\nConstantes:")
		for addr, value := range vm.constants {
			fmt.Printf("  Constante[%d] = %v (%T)\n", addr, value, value)
		}
	}

	// Información de funciones
	if len(vm.functionTable) > 0 {
		fmt.Println("\nTabla de Funciones:")
		for name, addr := range vm.functionTable {
			fmt.Printf("  %s -> PC: %d\n", name, addr)
		}
	}

	// Parámetros pendientes
	if len(vm.paramStack) > 0 {
		fmt.Println("\nParámetros pendientes:")
		for i, param := range vm.paramStack {
			fmt.Printf("  [%d] %v\n", i, param)
		}
	}

	fmt.Println(strings.Repeat("=", 60))
}
