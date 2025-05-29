package symbols

import (
	"fmt"
	"strconv"
	"strings"
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

// ==================== NUEVOS MÉTODOS PARA LA VM ====================

// GetConstants devuelve un mapa con dirección -> valor convertido para la VM
func (ct *ConstantTable) GetConstants() map[int]interface{} {
	constants := make(map[int]interface{})

	for key, addr := range ct.Constants {
		// El key tiene formato "tipo|valor"
		parts := strings.SplitN(key, "|", 2)
		if len(parts) != 2 {
			continue // Skip malformed keys
		}

		typ := parts[0]
		value := parts[1]

		// Convertir el valor según su tipo
		switch typ {
		case "int":
			if intVal, err := strconv.Atoi(value); err == nil {
				constants[addr] = intVal
			}
		case "float":
			if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
				constants[addr] = floatVal
			}
		case "bool":
			if boolVal, err := strconv.ParseBool(value); err == nil {
				constants[addr] = boolVal
			}
		case "string":
			constants[addr] = value // Las strings se mantienen como están
		}
	}

	return constants
}

// GetConstantValue obtiene el valor convertido de una constante por su dirección
func (ct *ConstantTable) GetConstantValue(addr int) (interface{}, bool) {
	// Buscar la constante por dirección
	for key, constAddr := range ct.Constants {
		if constAddr == addr {
			parts := strings.SplitN(key, "|", 2)
			if len(parts) != 2 {
				continue
			}

			typ := parts[0]
			value := parts[1]

			switch typ {
			case "int":
				if intVal, err := strconv.Atoi(value); err == nil {
					return intVal, true
				}
			case "float":
				if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
					return floatVal, true
				}
			case "bool":
				if boolVal, err := strconv.ParseBool(value); err == nil {
					return boolVal, true
				}
			case "string":
				return value, true
			}
		}
	}

	return nil, false
}

// GetConstantByValue busca una constante por valor y tipo, devuelve su dirección
func (ct *ConstantTable) GetConstantByValue(value string, typ string) (int, bool) {
	key := fmt.Sprintf("%s|%s", typ, value)
	addr, exists := ct.Constants[key]
	return addr, exists
}

// HasConstant verifica si existe una constante
func (ct *ConstantTable) HasConstant(value string, typ string) bool {
	key := fmt.Sprintf("%s|%s", typ, value)
	_, exists := ct.Constants[key]
	return exists
}

// Size devuelve el número de constantes almacenadas
func (ct *ConstantTable) Size() int {
	return len(ct.Constants)
}

// Clear limpia todas las constantes (útil para testing)
func (ct *ConstantTable) Clear() {
	ct.Constants = make(map[string]int)
}

// PrintDetailed imprime la tabla con más detalles para debug
func (ct *ConstantTable) PrintDetailed() {
	fmt.Println("=== Tabla de Constantes (Detallada) ===")
	fmt.Printf("Total de constantes: %d\n", len(ct.Constants))

	if len(ct.Constants) == 0 {
		fmt.Println("No hay constantes definidas.")
		return
	}

	fmt.Printf("%-15s %-10s %-15s %s\n", "Tipo", "Valor", "Dirección", "Valor Convertido")
	fmt.Println(strings.Repeat("-", 60))

	for key, addr := range ct.Constants {
		parts := strings.SplitN(key, "|", 2)
		if len(parts) != 2 {
			continue
		}

		typ := parts[0]
		value := parts[1]
		convertedValue, _ := ct.GetConstantValue(addr)

		fmt.Printf("%-15s %-10s %-15d %v\n", typ, value, addr, convertedValue)
	}

	fmt.Println(strings.Repeat("=", 60))
}