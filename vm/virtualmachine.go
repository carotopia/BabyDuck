package vm

import (
	"BabyDuckCompiler/memory" // Ajusta esta ruta
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

// ExecutionContext representa el contexto de ejecución de una función
type ExecutionContext struct {
	FunctionName   string
	ReturnAddress  int
	PreviousMemory *memory.ExecutionMemoryMap
	ParameterCount int
}

// VirtualMachine representa la máquina virtual mejorada
type VirtualMachine struct {
	globalMemory   *memory.ExecutionMemoryMap
	currentMemory  *memory.ExecutionMemoryMap
	quadruples     []Quadruple
	executionStack []ExecutionContext // Stack para contextos de función
	parameterStack []interface{}      // Stack para parámetros
	pc             int                // Program Counter
	debug          bool
	functionTable  map[string]int

	// Para capturar output
	outputWriter io.Writer
	debugWriter  io.Writer
}

func NewVirtualMachine(debug bool) *VirtualMachine {

	globalMemory := memory.NewExecutionMemoryMap()

	return &VirtualMachine{
		globalMemory:   globalMemory,
		currentMemory:  globalMemory, // 🔧 LA MISMA INSTANCIA
		quadruples:     make([]Quadruple, 0),
		executionStack: make([]ExecutionContext, 0),
		parameterStack: make([]interface{}, 0),
		pc:             0,
		debug:          debug,
		functionTable:  make(map[string]int),
		outputWriter:   nil,
		debugWriter:    nil,
	}
}

// SetOutputWriter configura el writer para capturar el output
func (vm *VirtualMachine) SetOutputWriter(writer io.Writer) {
	vm.outputWriter = writer
}

// SetDebugWriter configura el writer para el debug
func (vm *VirtualMachine) SetDebugWriter(writer io.Writer) {
	vm.debugWriter = writer
}

// print wraps the printing logic
func (vm *VirtualMachine) print(args ...interface{}) {
	if vm.outputWriter != nil {
		fmt.Fprint(vm.outputWriter, args...)
	} else {
		fmt.Print(args...)
	}
}

// println juntas el print con un salto de línea
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
	vm.preprocessFunctions()
	if vm.debug {
		vm.debugPrintln("VM: Cargados", len(quads), "cuádruplos")
		if len(vm.functionTable) > 0 {
			vm.debugPrintln("VM: Encontradas", len(vm.functionTable), "funciones")
		}
	}
}

// preprocessFunctions busca y registra todas las funciones
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
	for addr, value := range constants {
		switch {
		case addr >= 10000 && addr <= 10999:
			if intVal, err := vm.toInt(value); err == nil {
				vm.globalMemory.ConstantInts[addr] = intVal
			}
		case addr >= 11000 && addr <= 11999:
			if floatVal, err := vm.toFloat(value); err == nil {
				vm.globalMemory.ConstantFloats[addr] = floatVal
			}
		case addr >= 12000 && addr <= 12999:
			if boolVal, ok := value.(bool); ok {
				vm.globalMemory.ConstantBools[addr] = boolVal
			}
		case addr >= 13000 && addr <= 13999:
			if strVal, ok := value.(string); ok {
				vm.globalMemory.ConstantStrings[addr] = strVal
			}
		}
	}
	if vm.debug {
		vm.debugPrintln("VM: Cargadas", len(constants), "constantes")
	}
}

// getCurrentMemoryValue obtiene un valor de la memoria apropiada (similar al ejemplo)
func (vm *VirtualMachine) getCurrentMemoryValue(address int) (interface{}, error) {
	// Verificar si es una dirección global
	if vm.isGlobalAddress(address) {
		return vm.globalMemory.GetValue(address)
	}

	// Intentar con memoria actual primero
	val, err := vm.currentMemory.GetValue(address)
	if err != nil && vm.isTempAddress(address) {
		if vm.debug {
			vm.debugPrintln("    ⚠️ Temporal no encontrado en local, buscando en global:", address)
		}
		return vm.globalMemory.GetValue(address)
	}
	return val, err
}

func (vm *VirtualMachine) isTempAddress(address int) bool {
	return address >= 7000 && address <= 7999
}

// setCurrentMemoryValue establece un valor en la memoria apropiada
func (vm *VirtualMachine) setCurrentMemoryValue(address int, value interface{}) error {
	if vm.isGlobalAddress(address) {
		return vm.globalMemory.SetValue(address, value)
	}

	err := vm.currentMemory.SetValue(address, value)
	if err != nil && vm.isTempAddress(address) {
		// Si falla con memoria local y es temporal, intentar con global
		if vm.debug {
			vm.debugPrintln("    ⚠️ Temporal falló en local, usando global:", address)
		}
		return vm.globalMemory.SetValue(address, value)
	}
	return err
}

// isGlobalAddress determina si una dirección pertenece a memoria global
func (vm *VirtualMachine) isGlobalAddress(address int) bool {
	return (address >= 1000 && address <= 3999) || // Variables globales
		(address >= 10000 && address <= 13999) || // Constantes
		(address >= 7000 && address <= 7999) //
}

// getValue obtiene el valor usando el sistema de memoria mejorado
func (vm *VirtualMachine) getValue(operand interface{}) interface{} {
	if operand == nil {
		return nil
	}

	switch v := operand.(type) {
	case int:
		if val, err := vm.getCurrentMemoryValue(v); err == nil {
			if vm.debug {
				vm.debugPrint("    Obteniendo memoria[", v, "] = ", val, "\n")
			}
			return val
		}

		// 🔧 MEJOR MANEJO DE TEMPORALES NO ENCONTRADOS
		if v >= 7000 && v <= 9999 {
			if vm.debug {
				vm.debugPrint("  ⚠️ TEMPORAL[", v, "] no encontrado - esto puede ser un error\n")
			}
			// No inicializar automáticamente, devolver error
			return nil
		}

		// Si no se encuentra y es una dirección de memoria, inicializar con 0
		if v >= 1000 {
			if vm.debug {
				vm.debugPrint("  Memoria[", v, "] no encontrada, inicializando con 0\n")
			}
			vm.setValue(v, 0)
			return 0
		}

		// Para valores literales pequeños
		if vm.debug {
			vm.debugPrint("    Usando valor literal: ", v, "\n")
		}
		return v

	case string:
		if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
			cleaned := v[1 : len(v)-1]
			if vm.debug {
				vm.debugPrint("    String literal: ", v, " -> ", cleaned, "\n")
			}
			return cleaned
		}
		return v
	case float64:
		return v
	default:
		return v
	}
}

// setValue asigna un valor usando el sistema de memoria tipado
func (vm *VirtualMachine) setValue(address interface{}, value interface{}) {
	if addr, ok := address.(int); ok {
		// 🔧 CONVERSIÓN AUTOMÁTICA SEGÚN RANGO DE MEMORIA
		convertedValue := vm.convertValueForMemoryType(addr, value)

		err := vm.setCurrentMemoryValue(addr, convertedValue)
		if err != nil && vm.debug {
			vm.debugPrintln("    ⚠️ Error setValue:", err)
		} else if vm.debug {
			vm.debugPrintln("    ✅ Guardado memoria[", addr, "] =", convertedValue)
		}
	}
}
func (vm *VirtualMachine) Execute() error {
	if len(vm.quadruples) == 0 {
		vm.println("⚠️  No hay cuádruplos para ejecutar")
		return nil
	}

	vm.debugPrintln("\n", strings.Repeat("=", 60))
	vm.debugPrintln(" EJECUTANDO CON MÁQUINA VIRTUAL")
	vm.debugPrintln(strings.Repeat("=", 60))

	vm.pc = 0
	maxInstructions := len(vm.quadruples) * 1000 // Límite para evitar ciclos infinitos
	instructionCount := 0

	for vm.pc < len(vm.quadruples) {
		// Protección contra ciclos infinitos
		instructionCount++
		if instructionCount > maxInstructions {
			return fmt.Errorf("posible ciclo infinito detectado después de %d instrucciones", instructionCount)
		}

		if vm.pc < 0 || vm.pc >= len(vm.quadruples) {
			return fmt.Errorf("PC fuera de rango: %d (máximo: %d)", vm.pc, len(vm.quadruples)-1)
		}

		quad := vm.quadruples[vm.pc]
		shouldIncrementPC := true

		if vm.debug {
			vm.debugPrintln(fmt.Sprintf("\n[%d] PC:%d - %s %v %v %v",
				instructionCount, vm.pc, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result))
		}

		// Dispatcher similar al ejemplo
		switch quad.Operator {
		case "ERA":
			vm.executeERA(quad)
		case "PARAMETER":
			vm.executeParam(quad)
		case "GOSUB":
			shouldIncrementPC = vm.executeGosub(quad)
		case "FUNC":
			vm.executeFunc(quad)
		case "PARAM":
			vm.executeParamDeclaration(quad)
		case "ENDFUNC":
			shouldIncrementPC = vm.executeEndFunc()
		case "RET":
			shouldIncrementPC = vm.executeReturn(quad)
		case "=":
			vm.executeAssignment(quad)
		case "+", "-", "*", "/":
			vm.executeArithmetic(quad)
		case "print", "PRINT":
			vm.executePrint(quad)
		case ">", "<", ">=", "<=", "==", "!=":
			vm.executeComparison(quad)
		case "GOTOF":
			shouldIncrementPC = vm.executeConditionalJump(quad)
		case "GOTO":
			shouldIncrementPC = vm.executeUnconditionalJump(quad)
		default:
			if vm.debug {
				vm.debugPrintln("  ⚠️  Operación no reconocida:", quad.Operator)
			}
			// No hacer panic, solo continuar
		}

		if shouldIncrementPC {
			vm.pc++
			if vm.debug && vm.pc >= len(vm.quadruples) {
				vm.debugPrintln("  → PC incrementado a", vm.pc, "(fin de programa)")
			}
		} else {
			if vm.debug {
				vm.debugPrintln("  → PC cambiado a", vm.pc, "por salto/llamada")
			}
		}
	}

	vm.debugPrintln("\n✅ Ejecución completada exitosamente después de", instructionCount, "instrucciones.")
	return nil
}

// executeERA maneja la instrucción ERA (similar al ejemplo)
func (vm *VirtualMachine) executeERA(quad Quadruple) {
	var funcName string

	// Determinar el nombre de la función
	if name, ok := quad.LeftOperand.(string); ok {
		funcName = name
	} else if name, ok := quad.Result.(string); ok {
		funcName = name
	} else {
		funcName = "unknown" // Fallback
	}

	// Crear contexto de ejecución
	context := ExecutionContext{
		FunctionName:   funcName,
		ReturnAddress:  0,   // Se establecerá en GOSUB
		PreviousMemory: nil, // Se establecerá en GOSUB
		ParameterCount: 0,   // Se calculará según los parámetros
	}

	vm.executionStack = append(vm.executionStack, context)

	if vm.debug {
		vm.debugPrintln("  ERA: Preparando espacio para función", funcName, "(contextos activos:", len(vm.executionStack), ")")
	}
}

// executeParam maneja parámetros (similar al ejemplo)
func (vm *VirtualMachine) executeParam(quad Quadruple) {
	paramAddr := quad.LeftOperand.(int)
	paramValue := vm.getValue(paramAddr)
	vm.parameterStack = append(vm.parameterStack, paramValue)

	if vm.debug {
		vm.debugPrintln("  PARAMETER: Guardando parámetro", paramValue, "(total:", len(vm.parameterStack), ")")
	}
}

// executeGosub maneja llamadas a función (basado en el ejemplo mejorado)
func (vm *VirtualMachine) executeGosub(quad Quadruple) bool {
	var funcName string

	if name, ok := quad.LeftOperand.(string); ok {
		funcName = name
	} else if name, ok := quad.Result.(string); ok {
		funcName = name
	} else {
		panic(fmt.Sprintf("No se puede determinar el nombre de función en GOSUB: %v", quad))
	}

	// Si no hay contexto de ERA, crear uno automáticamente
	if len(vm.executionStack) == 0 {
		if vm.debug {
			vm.debugPrintln("  GOSUB: No hay contexto ERA, creando automáticamente para", funcName)
		}
		context := ExecutionContext{
			FunctionName:   funcName,
			ReturnAddress:  0,
			PreviousMemory: nil,
			ParameterCount: len(vm.parameterStack),
		}
		vm.executionStack = append(vm.executionStack, context)
	}

	contextIndex := len(vm.executionStack) - 1
	context := &vm.executionStack[contextIndex]
	context.FunctionName = funcName
	context.ReturnAddress = vm.pc + 1

	vm.currentMemory.PushActivationRecord(funcName, context.ReturnAddress)

	vm.copyParametersToLocalMemory(funcName)

	// Saltar a la función
	if funcAddr, exists := vm.functionTable[funcName]; exists {
		if vm.debug {
			vm.debugPrintln("  GOSUB: Llamando función", funcName, "PC:", vm.pc, "→", funcAddr, "(retorno a", context.ReturnAddress, ")")
		}
		vm.pc = funcAddr
		return false
	} else {
		panic(fmt.Sprintf("Función '%s' no encontrada en tabla de funciones", funcName))
	}
}

// copyParametersToLocalMemory copia parámetros a memoria local (SOLUCIÓN MEJORADA)
func (vm *VirtualMachine) copyParametersToLocalMemory(funcName string) {
	// Buscar la función para obtener información de parámetros
	funcAddr := vm.functionTable[funcName]

	// Buscar declaraciones PARAM después de FUNC
	paramAddresses := make([]int, 0)
	for i := funcAddr + 1; i < len(vm.quadruples) && vm.quadruples[i].Operator == "PARAM"; i++ {
		if paramAddr, ok := vm.quadruples[i].Result.(int); ok {
			paramAddresses = append(paramAddresses, paramAddr)
		}
	}

	for i, param := range vm.parameterStack {
		if i < len(paramAddresses) {
			paramAddr := paramAddresses[i]

			err := vm.currentMemory.SetValue(paramAddr, param)
			if err != nil {
				if vm.debug {
					vm.debugPrintln("    ⚠️ Error asignando parámetro:", err)
				}
			} else if vm.debug {
				vm.debugPrint("    Parámetro[", i, "] = ", param, " -> memoria[", paramAddr, "]\n")
			}
		}
	}

	// Limpiar stack de parámetros
	vm.parameterStack = vm.parameterStack[:0]
}

// executeEndFunc maneja retorno de función (similar al ejemplo)
func (vm *VirtualMachine) executeEndFunc() bool {
	if len(vm.executionStack) == 0 {
		if vm.debug {
			vm.debugPrintln("  ENDFUNC: Fin de función principal - terminando programa")
		}
		vm.pc = len(vm.quadruples)
		return false
	}

	// Obtener y remover el último contexto
	contextIndex := len(vm.executionStack) - 1
	context := vm.executionStack[contextIndex]
	vm.executionStack = vm.executionStack[:contextIndex]

	// 🔧 USAR PopActivationRecord EN LUGAR DE CAMBIAR MEMORIA
	vm.currentMemory.PopActivationRecord()

	if vm.debug {
		vm.debugPrintln("  ENDFUNC: Retornando de función", context.FunctionName, "PC:", vm.pc, "→", context.ReturnAddress)
	}

	vm.pc = context.ReturnAddress
	return false
}

// executeFunc maneja declaración de función
func (vm *VirtualMachine) executeFunc(quad Quadruple) {
	if vm.debug {
		if funcName, ok := quad.LeftOperand.(string); ok {
			vm.debugPrintln("  FUNC: Definición de función", funcName)
		}
	}
}

// executeParamDeclaration maneja declaración de parámetros
func (vm *VirtualMachine) executeParamDeclaration(quad Quadruple) {
	if vm.debug {
		vm.debugPrintln("  PARAM: Declaración de parámetro (saltando)")
	}
}

// executeReturn maneja retorno con valor
func (vm *VirtualMachine) executeReturn(quad Quadruple) bool {
	if len(vm.executionStack) > 0 {
		// Obtener y remover el último contexto
		contextIndex := len(vm.executionStack) - 1
		context := vm.executionStack[contextIndex]
		vm.executionStack = vm.executionStack[:contextIndex]

		if quad.LeftOperand != nil {
			returnValue := vm.getValue(quad.LeftOperand)
			if quad.Result != nil {
				vm.setValue(quad.Result, returnValue)
			}
			if vm.debug {
				vm.debugPrintln("  RET: Retornando valor", returnValue, "de función", context.FunctionName)
			}
		}

		// Restaurar memoria anterior
		if context.PreviousMemory != nil {
			vm.currentMemory = context.PreviousMemory
		} else {
			vm.currentMemory = vm.globalMemory
		}

		vm.pc = context.ReturnAddress - 1
	}
	return false
}

// executeAssignment maneja asignaciones
func (vm *VirtualMachine) executeAssignment(quad Quadruple) {
	value := vm.getValue(quad.LeftOperand)
	vm.setValue(quad.Result, value)
	if vm.debug {
		vm.debugPrintln("  =: Asignando", value, "a dirección", quad.Result)
	}
}

// executeArithmetic maneja operaciones aritméticas
func (vm *VirtualMachine) executeArithmetic(quad Quadruple) {
	left := vm.getValue(quad.LeftOperand)
	right := vm.getValue(quad.RightOperand)
	operator := quad.Operator

	result, err := vm.performArithmetic(left, right, operator)
	if err != nil {
		panic(fmt.Sprintf("Error en operación aritmética: %v", err))
	}

	vm.setValue(quad.Result, result)
	if vm.debug {
		vm.debugPrintln("  ", operator, ":", left, operator, right, "=", result)
	}
}

// executeComparison maneja operaciones de comparación
func (vm *VirtualMachine) executeComparison(quad Quadruple) {
	left := vm.getValue(quad.LeftOperand)
	right := vm.getValue(quad.RightOperand)
	operator := quad.Operator

	result, err := vm.performComparison(left, right, operator)
	if err != nil {
		panic(fmt.Sprintf("Error en comparación: %v", err))
	}

	vm.setValue(quad.Result, result)
	if vm.debug {
		vm.debugPrintln("  ", operator, ":", left, operator, right, "=", result)
	}
}

// executePrint maneja operaciones de impresión
func (vm *VirtualMachine) executePrint(quad Quadruple) {
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

	isFirstPrint := vm.pc == 0 || vm.quadruples[vm.pc-1].Operator != "print"

	isLastPrint := vm.pc+1 >= len(vm.quadruples) || vm.quadruples[vm.pc+1].Operator != "print"

	// Imprimir valores
	if len(values) > 0 {
		if isFirstPrint {
			vm.print(">>> ") // Solo mostrar >>> al comienzo
		}

		for i, value := range values {
			if i > 0 {
				vm.print(" ")
			}
			vm.print(value)
		}

		if isLastPrint {
			vm.println() // Solo nueva línea al final de la secuencia
		} else {
			vm.print(" ") // Espacio entre prints consecutivos
		}
	}
}

// executeConditionalJump maneja saltos condicionales
func (vm *VirtualMachine) executeConditionalJump(quad Quadruple) bool {
	condition := vm.getValue(quad.LeftOperand)
	conditionBool := vm.toBool(condition)

	if vm.debug {
		vm.debugPrint("  GOTOF - Condición: ", condition, " (", conditionBool, ")")
	}

	if !conditionBool { // Si la condición es FALSA
		if jumpAddr, ok := quad.Result.(int); ok {
			if vm.debug {
				vm.debugPrintln(" → Saltando a PC:", jumpAddr)
			}
			vm.pc = jumpAddr // Ir directamente, NO restar 1
			return false     // No incrementar PC
		}
	} else {
		if vm.debug {
			vm.debugPrintln(" → Continúa secuencial")
		}
	}
	return true // Incrementar PC normalmente
}

// executeUnconditionalJump maneja saltos incondicionales
func (vm *VirtualMachine) executeUnconditionalJump(quad Quadruple) bool {
	if jumpAddr, ok := quad.Result.(int); ok {
		if vm.debug {
			vm.debugPrintln("  GOTO: Salto incondicional PC:", vm.pc, "→", jumpAddr)
		}
		vm.pc = jumpAddr // Ir directamente a la dirección, NO restar 1
		return false     // No incrementar PC
	}
	if vm.debug {
		vm.debugPrintln("  GOTO: Error - no se pudo obtener dirección de salto")
	}
	return true
}

// performArithmetic realiza operaciones aritméticas
func (vm *VirtualMachine) performArithmetic(left, right interface{}, operator string) (interface{}, error) {
	// Intentar como enteros primero
	leftNum, err1 := vm.toInt(left)
	rightNum, err2 := vm.toInt(right)

	if err1 == nil && err2 == nil {
		switch operator {
		case "+":
			return leftNum + rightNum, nil
		case "-":
			return leftNum - rightNum, nil
		case "*":
			return leftNum * rightNum, nil
		case "/":
			if rightNum == 0 {
				return nil, fmt.Errorf("división por cero")
			}
			return leftNum / rightNum, nil
		}
	}

	// Intentar como floats
	leftFloat, ferr1 := vm.toFloat(left)
	rightFloat, ferr2 := vm.toFloat(right)

	if ferr1 == nil && ferr2 == nil {
		switch operator {
		case "+":
			return leftFloat + rightFloat, nil
		case "-":
			return leftFloat - rightFloat, nil
		case "*":
			return leftFloat * rightFloat, nil
		case "/":
			if rightFloat == 0 {
				return nil, fmt.Errorf("división por cero")
			}
			return leftFloat / rightFloat, nil
		}
	}

	return nil, fmt.Errorf("tipos incompatibles para operación %s: %T y %T", operator, left, right)
}

// performComparison realiza operaciones de comparación
func (vm *VirtualMachine) performComparison(left, right interface{}, operator string) (bool, error) {
	// Intentar como números
	leftNum, err1 := vm.toFloat(left)
	rightNum, err2 := vm.toFloat(right)

	if err1 == nil && err2 == nil {
		switch operator {
		case ">":
			return leftNum > rightNum, nil
		case "<":
			return leftNum < rightNum, nil
		case ">=":
			return leftNum >= rightNum, nil
		case "<=":
			return leftNum <= rightNum, nil
		case "==":
			return leftNum == rightNum, nil
		case "!=":
			return leftNum != rightNum, nil
		}
	}

	// Comparación directa para otros tipos
	switch operator {
	case "==":
		return left == right, nil
	case "!=":
		return left != right, nil
	default:
		return false, fmt.Errorf("tipos incompatibles para comparación %s: %T y %T", operator, left, right)
	}
}

// Funciones de utilidad
func (vm *VirtualMachine) toInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int", value)
	}
}

func (vm *VirtualMachine) toFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to float", value)
	}
}

// PrintQuadruples imprime todos los cuádruplos para debug
func (vm *VirtualMachine) PrintQuadruples() {
	vm.println("\n", strings.Repeat("=", 80))
	vm.println("CUÁDRUPLOS CARGADOS:")
	vm.println(strings.Repeat("=", 80))

	for i, quad := range vm.quadruples {
		vm.println(fmt.Sprintf("PC[%3d]: %-10s | %-15v | %-15v | %-15v",
			i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result))
	}

	vm.println(strings.Repeat("=", 80))
}

// PrintMemoryState imprime el estado actual de la memoria
func (vm *VirtualMachine) PrintMemoryState() {
	vm.println("\n", strings.Repeat("=", 60))
	vm.println("ESTADO FINAL DE MEMORIA")
	vm.println(strings.Repeat("=", 60))

	// Memoria global
	vm.println("Memoria Global:")

	if len(vm.globalMemory.GlobalInts) > 0 {
		vm.println("  Enteros globales:")
		for addr, value := range vm.globalMemory.GlobalInts {
			vm.println("    Global[", addr, "] = ", value)
		}
	}

	if len(vm.globalMemory.GlobalFloats) > 0 {
		vm.println("  Floats globales:")
		for addr, value := range vm.globalMemory.GlobalFloats {
			vm.println("    Global[", addr, "] = ", value)
		}
	}

	if len(vm.globalMemory.GlobalBools) > 0 {
		vm.println("  Bools globales:")
		for addr, value := range vm.globalMemory.GlobalBools {
			vm.println("    Global[", addr, "] = ", value)
		}
	}

	// Constantes
	vm.println("\nConstantes:")

	if len(vm.globalMemory.ConstantInts) > 0 {
		vm.println("  Constantes enteras:")
		for addr, value := range vm.globalMemory.ConstantInts {
			vm.println("    Const[", addr, "] = ", value)
		}
	}

	if len(vm.globalMemory.ConstantFloats) > 0 {
		vm.println("  Constantes float:")
		for addr, value := range vm.globalMemory.ConstantFloats {
			vm.println("    Const[", addr, "] = ", value)
		}
	}

	if len(vm.globalMemory.ConstantBools) > 0 {
		vm.println("  Constantes bool:")
		for addr, value := range vm.globalMemory.ConstantBools {
			vm.println("    Const[", addr, "] = ", value)
		}
	}

	if len(vm.globalMemory.ConstantStrings) > 0 {
		vm.println("  Constantes string:")
		for addr, value := range vm.globalMemory.ConstantStrings {
			vm.println("    Const[", addr, "] = ", value)
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
	if len(vm.parameterStack) > 0 {
		vm.println("\nParámetros pendientes:")
		for i, param := range vm.parameterStack {
			vm.println("  [", i, "] ", param)
		}
	}

	vm.println(strings.Repeat("=", 60))
}

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

// Reset reinicia el estado de la VM para una nueva ejecución
func (vm *VirtualMachine) Reset() {
	// Preservar constantes
	constantInts := make(map[int]int)
	constantFloats := make(map[int]float64)
	constantBools := make(map[int]bool)
	constantStrings := make(map[int]string)

	// Copiar constantes existentes
	for addr, val := range vm.globalMemory.ConstantInts {
		constantInts[addr] = val
	}
	for addr, val := range vm.globalMemory.ConstantFloats {
		constantFloats[addr] = val
	}
	for addr, val := range vm.globalMemory.ConstantBools {
		constantBools[addr] = val
	}
	for addr, val := range vm.globalMemory.ConstantStrings {
		constantStrings[addr] = val
	}

	// Reiniciar memoria
	vm.globalMemory = memory.NewExecutionMemoryMap()
	vm.currentMemory = vm.globalMemory

	// Restaurar constantes
	vm.globalMemory.ConstantInts = constantInts
	vm.globalMemory.ConstantFloats = constantFloats
	vm.globalMemory.ConstantBools = constantBools
	vm.globalMemory.ConstantStrings = constantStrings

	// Reiniciar program counter
	vm.pc = 0

	// Limpiar stacks
	vm.executionStack = make([]ExecutionContext, 0)
	vm.parameterStack = make([]interface{}, 0)

	if vm.debug {
		vm.debugPrintln("🔄 VM: Estado reiniciado para nueva ejecución")
	}
}
func (vm *VirtualMachine) convertValueForMemoryType(address int, value interface{}) interface{} {
	switch {
	case address >= 1000 && address <= 1999: // Int globales
		if intVal, err := vm.toInt(value); err == nil {
			return intVal
		}
	case address >= 2000 && address <= 2999: // Float globales
		if floatVal, err := vm.toFloat(value); err == nil {
			return floatVal
		}
	case address >= 4000 && address <= 4999: // Int locales
		if intVal, err := vm.toInt(value); err == nil {
			return intVal
		}
	case address >= 5000 && address <= 5999: // Float locales
		if floatVal, err := vm.toFloat(value); err == nil {
			return floatVal
		}
	case address >= 7000 && address <= 7999: // Temporales int
		if intVal, err := vm.toInt(value); err == nil {
			return intVal
		}
	case address >= 8000 && address <= 8999: // Temporales float
		if floatVal, err := vm.toFloat(value); err == nil {
			return floatVal
		}
	case address >= 9000 && address <= 9999: // Temporales bool
		return vm.toBool(value)
	}

	// Si no se puede convertir, devolver el valor original
	return value
}
