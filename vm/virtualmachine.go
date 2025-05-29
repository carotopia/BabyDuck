package vm

import (
	"fmt"
	"io"
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

// CallFrame representa un contexto de llamada a función
type CallFrame struct {
	FunctionName  string              // Nombre de la función
	ReturnAddress int                 // Dirección de retorno
	LocalMemory   map[int]interface{} // Memoria local de la función
	Parameters    []interface{}       // Parámetros de la función
}

// VirtualMachine representa la máquina virtual
type VirtualMachine struct {
	quadruples []Quadruple
	memory     map[int]interface{} // Memoria global por direcciones
	constants  map[int]interface{} // Tabla de constantes
	pc         int                 // Program Counter
	debug      bool                // Modo debug

	// Stack para manejo de funciones
	callStack  []CallFrame   // Stack de contextos de llamada
	paramStack []interface{} // Stack de parámetros temporales

	// Información de funciones
	functionTable map[string]int // nombre -> dirección de inicio

	// Para capturar output
	outputWriter io.Writer // Writer para capturar output
	debugWriter  io.Writer // Writer para debug output
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
		outputWriter:  nil,
		debugWriter:   nil,
	}
}

// SetOutputWriter configura el writer para capturar el output del programa
func (vm *VirtualMachine) SetOutputWriter(writer io.Writer) {
	vm.outputWriter = writer
}

// SetDebugWriter configura el writer para el debug
func (vm *VirtualMachine) SetDebugWriter(writer io.Writer) {
	vm.debugWriter = writer
}

// print wraps the printing logic to use the configured writer
func (vm *VirtualMachine) print(args ...interface{}) {
	if vm.outputWriter != nil {
		fmt.Fprint(vm.outputWriter, args...)
	} else {
		fmt.Print(args...)
	}
}

// println wraps the printing logic to use the configured writer
func (vm *VirtualMachine) println(args ...interface{}) {
	if vm.outputWriter != nil {
		fmt.Fprintln(vm.outputWriter, args...)
	} else {
		fmt.Println(args...)
	}
}

// debugPrint prints debug information
func (vm *VirtualMachine) debugPrint(args ...interface{}) {
	if !vm.debug {
		return
	}

	if vm.debugWriter != nil {
		fmt.Fprint(vm.debugWriter, args...)
	} else if vm.outputWriter != nil {
		fmt.Fprint(vm.outputWriter, args...)
	} else {
		fmt.Print(args...)
	}
}

// debugPrintln prints debug information with newline
func (vm *VirtualMachine) debugPrintln(args ...interface{}) {
	if !vm.debug {
		return
	}

	if vm.debugWriter != nil {
		fmt.Fprintln(vm.debugWriter, args...)
	} else if vm.outputWriter != nil {
		fmt.Fprintln(vm.outputWriter, args...)
	} else {
		fmt.Println(args...)
	}
}

// LoadQuadruples carga los cuádruplos en la máquina virtual
func (vm *VirtualMachine) LoadQuadruples(quads []Quadruple) {
	vm.quadruples = quads

	// Preprocesar para encontrar funciones
	vm.preprocessFunctions()

	if vm.debug {
		vm.debugPrintln("VM: Cargados", len(quads), "cuádruplos")
		if len(vm.functionTable) > 0 {
			vm.debugPrintln("VM: Encontradas", len(vm.functionTable), "funciones")
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
					vm.debugPrintln("VM: Registrando función", funcName, "en PC:", i)
				}
			}
		}
	}
}

// LoadConstants carga la tabla de constantes
func (vm *VirtualMachine) LoadConstants(constants map[int]interface{}) {
	vm.constants = constants
	if vm.debug {
		vm.debugPrintln("VM: Cargadas", len(constants), "constantes")
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
				vm.debugPrint("    Obteniendo constante[", v, "] = ", val, "\n")
			}
			return val
		}

		// Verificar memoria local si hay contexto de función activo
		if len(vm.callStack) > 0 {
			currentFrame := &vm.callStack[len(vm.callStack)-1]
			if val, exists := currentFrame.LocalMemory[v]; exists {
				if vm.debug {
					vm.debugPrint("    Obteniendo memoria local[", v, "] = ", val, " (función: ", currentFrame.FunctionName, ")\n")
				}
				return val
			}
		}

		// Verificar memoria global
		if val, exists := vm.memory[v]; exists {
			if vm.debug {
				vm.debugPrint("    Obteniendo memoria global[", v, "] = ", val, "\n")
			}
			return val
		}

		// Si no existe en ningún lado, devolver el valor directo
		if vm.debug {
			vm.debugPrint("    Usando valor directo: ", v, "\n")
		}
		return v
	case string:
		// Si es una cadena literal, remover comillas si las tiene
		if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
			cleaned := v[1 : len(v)-1]
			if vm.debug {
				vm.debugPrint("    String literal: ", v, " -> ", cleaned, "\n")
			}
			return cleaned
		}
		if vm.debug {
			vm.debugPrint("    String value: ", v, "\n")
		}
		return v
	case float64:
		if vm.debug {
			vm.debugPrint("    Float value: ", v, "\n")
		}
		return v
	default:
		if vm.debug {
			vm.debugPrint("    Other value (", fmt.Sprintf("%T", v), "): ", v, "\n")
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
					vm.debugPrint("    Asignando memoria local[", addr, "] = ", value, " (función: ", currentFrame.FunctionName, ")\n")
				}
				return
			}
		}

		// Asignar a memoria global
		vm.memory[addr] = value
		if vm.debug {
			vm.debugPrint("    Asignando memoria global[", addr, "] = ", value, "\n")
		}
	} else if vm.debug {
		vm.debugPrint("    ⚠️  Dirección inválida para asignación: ", address, " (", fmt.Sprintf("%T", address), ")\n")
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
		vm.println("⚠️  No hay cuádruplos para ejecutar")
		return nil
	}

	vm.debugPrintln("\n", strings.Repeat("=", 60))
	vm.debugPrintln("🚀 EJECUTANDO CON MÁQUINA VIRTUAL")
	vm.debugPrintln(strings.Repeat("=", 60))

	// Inicializar PC en 0
	vm.pc = 0

	for vm.pc < len(vm.quadruples) {
		if vm.pc < 0 || vm.pc >= len(vm.quadruples) {
			return fmt.Errorf("PC fuera de rango: %d (máximo: %d)", vm.pc, len(vm.quadruples)-1)
		}

		quad := vm.quadruples[vm.pc]

		if vm.debug {
			vm.debugPrintln("\nPC:", vm.pc, "- Ejecutando:", quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result)
		}

		err := vm.executeQuadruple(quad)
		if err != nil {
			return fmt.Errorf("error en PC %d: %v", vm.pc, err)
		}

		vm.pc++
	}

	vm.debugPrintln("\n✅ Ejecución completada exitosamente.")
	return nil
}

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
			vm.debugPrintln("  ERA: Preparando espacio para función", funcName)
		}
		return nil
	}

	if op == "PARAMETER" {
		value := vm.getValue(quad.LeftOperand)
		vm.paramStack = append(vm.paramStack, value)
		if vm.debug {
			vm.debugPrintln("  PARAMETER: Guardando parámetro", value, "(total:", len(vm.paramStack), ")")
		}
		return nil
	}

	if op == "GOSUB" {
		var funcAddr int
		var funcName string

		// Determinar función a llamar
		if name, ok := quad.LeftOperand.(string); ok {
			funcName = name
			if addr, exists := vm.functionTable[name]; exists {
				funcAddr = addr
			} else {
				return fmt.Errorf("función '%s' no encontrada", name)
			}
		} else if addr, ok := quad.Result.(int); ok {
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
			vm.debugPrintln("  GOSUB: Intentando llamar función", funcName, "en PC:", funcAddr)
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
				vm.debugPrint("    Parámetro[", i, "] = ", param, " -> memoria local[", paramAddr, "]\n")
			}
		}

		// Limpiar stack de parámetros
		vm.paramStack = vm.paramStack[:0]

		// Agregar frame al call stack
		vm.callStack = append(vm.callStack, frame)

		// Saltar a la función
		vm.pc = funcAddr - 1
		if vm.debug {
			vm.debugPrintln("  GOSUB: Llamando función", funcName, "en PC:", funcAddr)
		}
		return nil
	}

	if op == "FUNC" {
		if vm.debug {
			if funcName, ok := quad.LeftOperand.(string); ok {
				vm.debugPrintln("  FUNC: Definición de función", funcName)
			}
		}
		return nil
	}

	if op == "PARAM" {
		if vm.debug {
			vm.debugPrintln("  PARAM: Declaración de parámetro (saltando)")
		}
		return nil
	}

	if op == "ENDFUNC" {
		if len(vm.callStack) > 0 {
			frame := vm.callStack[len(vm.callStack)-1]
			vm.callStack = vm.callStack[:len(vm.callStack)-1]

			vm.pc = frame.ReturnAddress - 1
			if vm.debug {
				vm.debugPrintln("  ENDFUNC: Retornando de función", frame.FunctionName, "a PC:", frame.ReturnAddress)
			}
		} else {
			if vm.debug {
				vm.debugPrintln("  ENDFUNC: Fin de función principal")
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
					vm.debugPrintln("  RET: Retornando valor", returnValue, "de función", frame.FunctionName)
				}
			}

			vm.pc = frame.ReturnAddress - 1
			if vm.debug {
				vm.debugPrintln("  RET: Retornando de función", frame.FunctionName, "a PC:", frame.ReturnAddress)
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
			vm.print(">>> ")
			for i, value := range values {
				if i > 0 {
					vm.print(" ")
				}
				vm.print(value)
			}
			vm.println()
		} else {
			vm.println(">>> <valor nulo>")
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
			vm.debugPrint("  GOTOF - Condición: ", condition, " (", conditionBool, ")")
		}

		if !conditionBool { // Si la condición es FALSA
			if jumpAddr, ok := quad.Result.(int); ok {
				vm.pc = jumpAddr - 1 // -1 porque se incrementará al final del loop
				if vm.debug {
					vm.debugPrintln(" -> Saltando a PC:", jumpAddr)
				}
			}
		} else {
			if vm.debug {
				vm.debugPrintln(" -> Continúa secuencial")
			}
		}

	case "GOTO":
		// Salto incondicional
		if jumpAddr, ok := quad.Result.(int); ok {
			vm.pc = jumpAddr - 1 // -1 porque se incrementará al final del loop
			if vm.debug {
				vm.debugPrintln("  Salto incondicional a PC:", jumpAddr)
			}
		}

	default:
		if vm.debug {
			vm.debugPrintln("  ⚠️  Operación no reconocida:", op)
		}
		return fmt.Errorf("operación no implementada: %s", op)
	}

	return nil
}

// PrintMemoryState imprime el estado actual de la memoria
func (vm *VirtualMachine) PrintMemoryState() {
	vm.println("\n", strings.Repeat("=", 60))
	vm.println("ESTADO FINAL DE MEMORIA")
	vm.println(strings.Repeat("=", 60))

	// Memoria global
	vm.println("Memoria Global:")
	if len(vm.memory) == 0 {
		vm.println("  Memoria global vacía")
	} else {
		for addr, value := range vm.memory {
			vm.println("  Memoria[", addr, "] = ", value, " (", fmt.Sprintf("%T", value), ")")
		}
	}

	// Call stack activo
	if len(vm.callStack) > 0 {
		vm.println("\nCall Stack activo (", len(vm.callStack), " frames):")
		for i, frame := range vm.callStack {
			vm.println("  [", i, "] Función '", frame.FunctionName, "' -> Retorno: PC ", frame.ReturnAddress)
			if len(frame.LocalMemory) > 0 {
				vm.println("      Memoria local:")
				for addr, value := range frame.LocalMemory {
					vm.println("        Local[", addr, "] = ", value, " (", fmt.Sprintf("%T", value), ")")
				}
			}
			if len(frame.Parameters) > 0 {
				vm.println("      Parámetros: ", frame.Parameters)
			}
		}
	}

	// Constantes
	if len(vm.constants) > 0 {
		vm.println("\nConstantes:")
		for addr, value := range vm.constants {
			vm.println("  Constante[", addr, "] = ", value, " (", fmt.Sprintf("%T", value), ")")
		}
	}

	// Información de funciones
	if len(vm.functionTable) > 0 {
		vm.println("\nTabla de Funciones:")
		for name, addr := range vm.functionTable {
			vm.println("  ", name, " -> PC: ", addr)
		}
	}

	// Parámetros pendientes
	if len(vm.paramStack) > 0 {
		vm.println("\nParámetros pendientes:")
		for i, param := range vm.paramStack {
			vm.println("  [", i, "] ", param)
		}
	}

	vm.println(strings.Repeat("=", 60))
}
