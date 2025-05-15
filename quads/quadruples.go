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
	Items []Quadruple
}

// NewQuadrupleQueue crea una nueva fila de cuádruplos vacía
func NewQuadrupleQueue() *QuadrupleQueue {
	return &QuadrupleQueue{
		Items: make([]Quadruple, 0),
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
func (q *QuadrupleQueue) UpdateResult(index int, result interface{}) bool {
	if index < 0 || index >= len(q.Items) {
		return false
	}
	q.Items[index].Result = result
	return true
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
