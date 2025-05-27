package memory

import "fmt"

// ExecutionMemoryMap - Mapa de memoria para la máquina virtual
type ExecutionMemoryMap struct {
	// Memoria Global
	GlobalInts   map[int]int     // 1000-1999
	GlobalFloats map[int]float64 // 2000-2999
	GlobalBools  map[int]bool    // 3000-3999

	// Pila de Activación (para funciones)
	ActivationStack []*ActivationRecord

	// Memoria de Constantes
	ConstantInts    map[int]int     // 10000-10999
	ConstantFloats  map[int]float64 // 11000-11999
	ConstantBools   map[int]bool    // 12000-12999
	ConstantStrings map[int]string  // 13000-13999
}

// ActivationRecord - Registro de activación para cada función
type ActivationRecord struct {
	FunctionName string

	// Memoria Local de la función
	LocalInts   map[int]int     // 4000-4999
	LocalFloats map[int]float64 // 5000-5999
	LocalBools  map[int]bool    // 6000-6999

	// Memoria Temporal de la función
	TempInts   map[int]int     // 7000-7999
	TempFloats map[int]float64 // 8000-8999
	TempBools  map[int]bool    // 9000-9999

	// Dirección de retorno
	ReturnAddress int

	// Parámetros pasados a la función
	Parameters map[int]interface{}
}

// NewExecutionMemoryMap crea un nuevo mapa de memoria
func NewExecutionMemoryMap() *ExecutionMemoryMap {
	return &ExecutionMemoryMap{
		GlobalInts:      make(map[int]int),
		GlobalFloats:    make(map[int]float64),
		GlobalBools:     make(map[int]bool),
		ActivationStack: make([]*ActivationRecord, 0),
		ConstantInts:    make(map[int]int),
		ConstantFloats:  make(map[int]float64),
		ConstantBools:   make(map[int]bool),
		ConstantStrings: make(map[int]string),
	}
}

// GetValue obtiene un valor de cualquier segmento de memoria
func (emm *ExecutionMemoryMap) GetValue(address int) (interface{}, error) {
	switch {
	// Variables Globales
	case address >= 1000 && address <= 1999:
		if val, exists := emm.GlobalInts[address]; exists {
			return val, nil
		}
	case address >= 2000 && address <= 2999:
		if val, exists := emm.GlobalFloats[address]; exists {
			return val, nil
		}
	case address >= 3000 && address <= 3999:
		if val, exists := emm.GlobalBools[address]; exists {
			return val, nil
		}

	// Variables Locales (registro de activación actual)
	case address >= 4000 && address <= 4999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			if val, exists := currentRecord.LocalInts[address]; exists {
				return val, nil
			}
		}
	case address >= 5000 && address <= 5999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			if val, exists := currentRecord.LocalFloats[address]; exists {
				return val, nil
			}
		}
	case address >= 6000 && address <= 6999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			if val, exists := currentRecord.LocalBools[address]; exists {
				return val, nil
			}
		}

	// Variables Temporales
	case address >= 7000 && address <= 7999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			if val, exists := currentRecord.TempInts[address]; exists {
				return val, nil
			}
		}
	case address >= 8000 && address <= 8999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			if val, exists := currentRecord.TempFloats[address]; exists {
				return val, nil
			}
		}
	case address >= 9000 && address <= 9999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			if val, exists := currentRecord.TempBools[address]; exists {
				return val, nil
			}
		}

	// Constantes
	case address >= 10000 && address <= 10999:
		if val, exists := emm.ConstantInts[address]; exists {
			return val, nil
		}
	case address >= 11000 && address <= 11999:
		if val, exists := emm.ConstantFloats[address]; exists {
			return val, nil
		}
	case address >= 12000 && address <= 12999:
		if val, exists := emm.ConstantBools[address]; exists {
			return val, nil
		}
	case address >= 13000 && address <= 13999:
		if val, exists := emm.ConstantStrings[address]; exists {
			return val, nil
		}
	}

	return nil, fmt.Errorf("dirección %d no encontrada en memoria", address)
}

// SetValue establece un valor en cualquier segmento de memoria
func (emm *ExecutionMemoryMap) SetValue(address int, value interface{}) error {
	switch {
	// Variables Globales
	case address >= 1000 && address <= 1999:
		emm.GlobalInts[address] = value.(int)
	case address >= 2000 && address <= 2999:
		emm.GlobalFloats[address] = value.(float64)
	case address >= 3000 && address <= 3999:
		emm.GlobalBools[address] = value.(bool)

	// Variables Locales
	case address >= 4000 && address <= 4999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			currentRecord.LocalInts[address] = value.(int)
		}
	case address >= 5000 && address <= 5999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			currentRecord.LocalFloats[address] = value.(float64)
		}
	case address >= 6000 && address <= 6999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			currentRecord.LocalBools[address] = value.(bool)
		}

	// Variables Temporales
	case address >= 7000 && address <= 7999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			currentRecord.TempInts[address] = value.(int)
		}
	case address >= 8000 && address <= 8999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			currentRecord.TempFloats[address] = value.(float64)
		}
	case address >= 9000 && address <= 9999:
		if len(emm.ActivationStack) > 0 {
			currentRecord := emm.ActivationStack[len(emm.ActivationStack)-1]
			currentRecord.TempBools[address] = value.(bool)
		}

	default:
		return fmt.Errorf("no se puede escribir en la dirección %d", address)
	}

	return nil
}

// PushActivationRecord crea un nuevo registro de activación
func (emm *ExecutionMemoryMap) PushActivationRecord(functionName string, returnAddress int) {
	record := &ActivationRecord{
		FunctionName:  functionName,
		LocalInts:     make(map[int]int),
		LocalFloats:   make(map[int]float64),
		LocalBools:    make(map[int]bool),
		TempInts:      make(map[int]int),
		TempFloats:    make(map[int]float64),
		TempBools:     make(map[int]bool),
		ReturnAddress: returnAddress,
		Parameters:    make(map[int]interface{}),
	}

	emm.ActivationStack = append(emm.ActivationStack, record)
}

// PopActivationRecord elimina el registro de activación actual
func (emm *ExecutionMemoryMap) PopActivationRecord() *ActivationRecord {
	if len(emm.ActivationStack) == 0 {
		return nil
	}

	record := emm.ActivationStack[len(emm.ActivationStack)-1]
	emm.ActivationStack = emm.ActivationStack[:len(emm.ActivationStack)-1]

	return record
}

// LoadConstants carga la tabla de constantes en memoria
func (emm *ExecutionMemoryMap) LoadConstants(constantTable map[string]int) {
	// Implementar carga de constantes desde la tabla
	// Por ejemplo: emm.ConstantInts[10000] = 2
}
