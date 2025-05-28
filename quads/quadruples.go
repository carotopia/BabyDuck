package quads

import (
	"fmt"
	"strings"
)

// Quadruple structure
// Represents a quadruple in the form of (operator, left operand, right operand, result)
type Quadruple struct {
	Operator     string
	LeftOperand  interface{}
	RightOperand interface{}
	Result       interface{}
}

// QuadrupleQueue implementa una fila para almacenar cuádruplos
type QuadrupleQueue struct {
	Items      []Quadruple
	TemCounter int
}

// NewQuadrupleQueue crea una nueva fila de cuádruplos vacía
func NewQuadrupleQueue() *QuadrupleQueue {
	return &QuadrupleQueue{
		Items:      make([]Quadruple, 0),
		TemCounter: 0,
	}
}

// Add agrega un nuevo cuádruplo a la fila
func (q *QuadrupleQueue) Add(operator string, leftOperand, rightOperand, result interface{}) int {
	quad := Quadruple{
		Operator:     operator,
		LeftOperand:  leftOperand,
		RightOperand: rightOperand,
		Result:       result,
	}
	q.Items = append(q.Items, quad)
	return len(q.Items) - 1 // Retorna el índice del cuádruplo agregado
}

// AddQuadruple agrega un cuádruplo ya construido a la fila
func (q *QuadrupleQueue) AddQuadruple(quad Quadruple) int {
	q.Items = append(q.Items, quad)
	return len(q.Items) - 1 // Retorna el índice del cuádruplo agregado
}

// Get obtiene un cuádruplo en un índice específico
func (q *QuadrupleQueue) Get(index int) (Quadruple, bool) {
	if index < 0 || index >= len(q.Items) {
		return Quadruple{}, false
	}
	return q.Items[index], true
}

// GetAll retorna todos los cuádruplos
func (q *QuadrupleQueue) GetAll() []Quadruple {
	return q.Items
}

// Size retorna el número de cuádruplos en la fila
func (q *QuadrupleQueue) Size() int {
	return len(q.Items)
}

// Update actualiza un cuádruplo existente en un índice específico
// Útil para el backpatching
func (q *QuadrupleQueue) Update(index int, quad Quadruple) bool {
	if index < 0 || index >= len(q.Items) {
		return false
	}
	q.Items[index] = quad
	return true
}

// UpdateResult actualiza solo el resultado de un cuádruplo
// Comúnmente usado para llenar los saltos pendientes
func (q *QuadrupleQueue) UpdateResult(index int, value interface{}) {
	if index >= 0 && index < len(q.Items) {
		q.Items[index].Result = value
	}
}

// String convierte la fila de cuádruplos a una representación legible
func (q *QuadrupleQueue) String() string {
	var output strings.Builder
	for i, quad := range q.Items {
		output.WriteString(fmt.Sprintf("%d: (%v, %v, %v, %v)\n",
			i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result))
	}
	return output.String()
}
func (q *QuadrupleQueue) NewTemp(resultType string) string {
	q.TemCounter++
	return fmt.Sprintf("t%d", q.TemCounter)
}

func (q *QuadrupleQueue) FillJump(index int, value interface{}) {
	if index >= 0 && index < len(q.Items) {
		q.Items[index].Result = value
	}
}

// FunctionOperations - Extensión para tu QuadrupleQueue existente
type FunctionOperations struct {
	*QuadrupleQueue
	JumpStack     []int                    // Para manejar saltos pendientes
	FunctionTable map[string]*FunctionInfo // Tabla de funciones
}

// FunctionInfo almacena información de funciones
type FunctionInfo struct {
	Name         string
	ReturnType   string
	StartAddress int
	Parameters   []Parameter
	LocalVars    map[string]int // nombre -> dirección
	Size         int
}

type Parameter struct {
	Name    string
	Type    string
	Address int
}

// NewFunctionOperations crea un nuevo manejador de operaciones de función
func NewFunctionOperations() *FunctionOperations {
	return &FunctionOperations{
		QuadrupleQueue: NewQuadrupleQueue(),
		JumpStack:      make([]int, 0),
		FunctionTable:  make(map[string]*FunctionInfo),
	}
}

// === CUÁDRUPLOS ESPECÍFICOS DE FUNCIONES ===

// GenerateFUNC genera cuádruple para declaración de función
func (fo *FunctionOperations) GenerateFUNC(name, returnType string) int {
	startAddr := fo.Size()
	fo.Add("FUNC", name, returnType, startAddr)

	// Crear entrada en tabla de funciones
	fo.FunctionTable[name] = &FunctionInfo{
		Name:         name,
		ReturnType:   returnType,
		StartAddress: startAddr,
		Parameters:   make([]Parameter, 0),
		LocalVars:    make(map[string]int),
		Size:         0,
	}

	return startAddr
}

// GeneratePARAM genera cuádruple para parámetro de función
func (fo *FunctionOperations) GeneratePARAM(funcName, paramType, paramName string, address int) {
	fo.Add("PARAM", paramType, paramName, address)

	// Agregar a la función
	if funcInfo, exists := fo.FunctionTable[funcName]; exists {
		param := Parameter{
			Name:    paramName,
			Type:    paramType,
			Address: address,
		}
		funcInfo.Parameters = append(funcInfo.Parameters, param)
		funcInfo.Size++
	}
}

// GenerateENDFUNC genera cuádruple para fin de función
func (fo *FunctionOperations) GenerateENDFUNC() {
	fo.Add("ENDFUNC", nil, nil, nil)
}

// GenerateERA genera cuádruple ERA (Espacio Registro Activación)
func (fo *FunctionOperations) GenerateERA(functionName string) {
	if funcInfo, exists := fo.FunctionTable[functionName]; exists {
		fo.Add("ERA", functionName, nil, funcInfo.Size)
	}
}

// GeneratePARAMETER genera cuádruple para pasar parámetro
func (fo *FunctionOperations) GeneratePARAMETER(argument interface{}, position int) {
	fo.Add("PARAMETER", argument, nil, position)
}

// GenerateGOSUB genera cuádruple GOSUB (llamada a función)
func (fo *FunctionOperations) GenerateGOSUB(functionName string) {
	if funcInfo, exists := fo.FunctionTable[functionName]; exists {
		fo.Add("GOSUB", functionName, nil, funcInfo.StartAddress)
	}
}

// GenerateRETURN genera cuádruple RETURN
func (fo *FunctionOperations) GenerateRETURN(returnValue interface{}) {
	fo.Add("RETURN", returnValue, nil, nil)
}
// === NUEVOS MÉTODOS PARA CUÁDRUPLOS DE FUNCIONES ===

// GenerateFUNC genera cuádruple para declaración de función
func (q *QuadrupleQueue) GenerateFUNC(name, returnType string) int {
	return q.Add("FUNC", name, returnType, q.Size())
}

// GeneratePARAM genera cuádruple para parámetro de función
func (q *QuadrupleQueue) GeneratePARAM(paramType, paramName string, address interface{}) {
	q.Add("PARAM", paramType, paramName, address)
}

// GenerateENDFUNC genera cuádruple para fin de función
func (q *QuadrupleQueue) GenerateENDFUNC() {
	q.Add("ENDFUNC", nil, nil, nil)
}

// GenerateERA genera cuádruple ERA (Espacio Registro Activación)
func (q *QuadrupleQueue) GenerateERA(functionName string, size int) {
	q.Add("ERA", functionName, nil, size)
}

// GeneratePARAMETER genera cuádruple para pasar parámetro
func (q *QuadrupleQueue) GeneratePARAMETER(argument interface{}, position int) {
	q.Add("PARAMETER", argument, nil, position)
}

// GenerateGOSUB genera cuádruple GOSUB (llamada a función)
func (q *QuadrupleQueue) GenerateGOSUB(functionName string, startAddr int) {
	q.Add("GOSUB", functionName, nil, startAddr)
}

// GenerateRETURN genera cuádruple RETURN
func (q *QuadrupleQueue) GenerateRETURN(returnValue interface{}) {
	q.Add("RETURN", returnValue, nil, nil)
}