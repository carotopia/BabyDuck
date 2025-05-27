package symbols

import (
	"fmt"
)

type ConstantTable struct {
	Constants map[string]int // valor como string -> dirección virtual
	// Memory    *memory.MemoryManager // ← COMENTADO temporalmente
}

// Contadores para constantes (similar a los de function.go)
var constIntCounter = 10000
var constFloatCounter = 11000
var constBoolCounter = 12000
var constStringCounter = 13000

// NewConstantTable crea una nueva tabla de constantes.
func NewConstantTable() *ConstantTable { // ← SIN parámetros
	return &ConstantTable{
		Constants: make(map[string]int),
		// Memory:    mem,  // ← COMENTADO
	}
}

// AddConstant agrega la constante si no existe y regresa la dirección virtual
// typ debe ser "int", "float", "bool" o "string"
func (ct *ConstantTable) AddConstant(value string, typ string) int {
	key := fmt.Sprintf("%s|%s", typ, value) // Asegura unicidad por tipo+valor
	if addr, ok := ct.Constants[key]; ok {
		return addr // Ya existe
	}

	var addr int
	switch typ {
	case "int":
		addr = constIntCounter
		constIntCounter++
	case "float":
		addr = constFloatCounter
		constFloatCounter++
	case "bool":
		addr = constBoolCounter
		constBoolCounter++
	case "string":
		addr = constStringCounter
		constStringCounter++
	default:
		panic("Tipo de constante no soportado: " + typ)
	}

	ct.Constants[key] = addr
	return addr
}

// Para debug: imprime la tabla de constantes
func (ct *ConstantTable) Print() {
	fmt.Println("=== Tabla de Constantes ===")
	for k, addr := range ct.Constants {
		fmt.Printf("Constante %s -> %d\n", k, addr)
	}
}

func (ct *ConstantTable) GetConstAddress(value string, typ string) (int, error) {
	key := fmt.Sprintf("%s|%s", typ, value)
	if addr, ok := ct.Constants[key]; ok {
		return addr, nil
	}

	var addr int
	switch typ {
	case "int":
		addr = constIntCounter
		constIntCounter++
	case "float":
		addr = constFloatCounter
		constFloatCounter++
	case "bool":
		addr = constBoolCounter
		constBoolCounter++
	case "string":
		addr = constStringCounter
		constStringCounter++
	default:
		return -1, fmt.Errorf("tipo de constante no soportado: %s", typ)
	}

	ct.Constants[key] = addr
	return addr, nil
}
