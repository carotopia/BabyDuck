package quads

// OperatorStack implementa una pila para operadores
type OperatorStack struct {
	items []string
}

// NewOperatorStack crea una nueva pila de operadores
func NewOperatorStack() *OperatorStack {
	return &OperatorStack{
		items: make([]string, 0),
	}
}

// Push agrega un operador a la pila
func (s *OperatorStack) Push(op string) {
	s.items = append(s.items, op)
}

// Pop remueve y retorna el operador en el tope de la pila
func (s *OperatorStack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item, true
}

// Peek retorna el operador en el tope sin removerlo
func (s *OperatorStack) Peek() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty verifica si la pila está vacía
func (s *OperatorStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size retorna el tamaño de la pila
func (s *OperatorStack) Size() int {
	return len(s.items)
}

// Stack implementa una pila genérica para operandos
type Stack struct {
	items []interface{}
}

// NewOperandStack crea una nueva pila de operandos
func NewOperandStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push agrega un operando a la pila
func (s *Stack) Push(operand interface{}) {
	s.items = append(s.items, operand)
}

// Pop remueve y retorna el operando en el tope de la pila
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item, true
}

// Peek retorna el operando en el tope sin removerlo
func (s *Stack) Peek() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty verifica si la pila está vacía
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size retorna el tamaño de la pila
func (s *Stack) Size() int {
	return len(s.items)
}

// TypeStack implementa una pila específica para tipos
type TypeStack struct {
	items []string
}

// NewTypeStack crea una nueva pila de tipos
func NewTypeStack() *TypeStack {
	return &TypeStack{
		items: make([]string, 0),
	}
}

// Push agrega un tipo a la pila
func (s *TypeStack) Push(typ string) {
	s.items = append(s.items, typ)
}

// Pop remueve y retorna el tipo en el tope de la pila
func (s *TypeStack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item, true
}

// Peek retorna el tipo en el tope sin removerlo
func (s *TypeStack) Peek() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty verifica si la pila está vacía
func (s *TypeStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size retorna el tamaño de la pila
func (s *TypeStack) Size() int {
	return len(s.items)
}
