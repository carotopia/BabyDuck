package quads

type Stack struct {
	items []interface{}
}

// TypeStack implements a stack for types
type TypeStack struct {
	items []string
}

// OperatorStack implements a stack for operators
type OperatorStack struct {
	items []string
}

// OperandStack implements a stack for operands
type OperandStack struct {
	items []interface{}
}
func NewOperatorStack() *OperatorStack {
	return &OperatorStack{
		items: make([]string, 0),
	}
}

// Constructor para OperandStack (puedes usar Stack genérico directamente)
func NewOperandStack() *Stack {
	return NewStack()
}


// Stack structure
func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}
func (s *Stack) Size() int {
	return len(s.items)
}
func (s *Stack) Clear() {
	s.items = make([]interface{}, 0)
}

func NewTypeStack() *TypeStack {
	return &TypeStack{
		items: make([]string, 0),
	}
}

func (s *TypeStack) Push(dataType string) {
	s.items = append(s.items, dataType)
}

func (s *TypeStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}
func (s *TypeStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *OperatorStack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	return s.items[len(s.items)-1], true
}

func (s *OperatorStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *OperatorStack) Size() int {
	return len(s.items)
}

// Clear elimina todos los elementos de la pila
func (s *OperatorStack) Clear() {
	s.items = make([]string, 0)
}

