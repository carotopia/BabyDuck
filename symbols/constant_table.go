package symbols

import (
	"BabyDuckCompiler/memory"
	"fmt"
)

type ConstantTable struct {
	Constants map[string]int        // valor como string -> dirección virtual
	Memory    *memory.MemoryManager // referencia al MemoryManager global
}

// NewConstantTable crea una nueva tabla de constantes.
func NewConstantTable(mem *memory.MemoryManager) *ConstantTable {
	return &ConstantTable{
		Constants: make(map[string]int),
		Memory:    mem,
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
		addr = ct.Memory.NextConstInt()
	case "float":
		addr = ct.Memory.NextConstFloat()
	case "bool":
		addr = ct.Memory.NextConstBool()

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
