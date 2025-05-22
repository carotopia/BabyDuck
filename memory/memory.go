package memory

// Configuration defines the memory address ranges for each segment and type.
type Configuration struct {
	GlobalIntStart, GlobalIntEnd     int
	GlobalFloatStart, GlobalFloatEnd int
	GlobalBoolStart, GlobalBoolEnd   int

	LocalIntStart, LocalIntEnd     int
	LocalFloatStart, LocalFloatEnd int
	LocalBoolStart, LocalBoolEnd   int

	TempIntStart, TempIntEnd     int
	TempFloatStart, TempFloatEnd int
	TempBoolStart, TempBoolEnd   int

	ConstIntStart, ConstIntEnd     int
	ConstFloatStart, ConstFloatEnd int
	ConstBoolStart, ConstBoolEnd   int

	// Si necesitas string, agrégalo aquí...
}

// DefaultMemoryConfig sets the standard memory map for your compiler
var DefaultMemoryConfig = Configuration{
	GlobalIntStart:   1000, GlobalIntEnd:   1999,
	GlobalFloatStart: 2000, GlobalFloatEnd: 2999,
	GlobalBoolStart:  3000, GlobalBoolEnd:  3999,

	LocalIntStart:    4000, LocalIntEnd:    4999,
	LocalFloatStart:  5000, LocalFloatEnd:  5999,
	LocalBoolStart:   6000, LocalBoolEnd:   6999,

	TempIntStart:     7000, TempIntEnd:     7999,
	TempFloatStart:   8000, TempFloatEnd:   8999,
	TempBoolStart:    9000, TempBoolEnd:    9999,

	ConstIntStart:    10000, ConstIntEnd:    10999,
	ConstFloatStart:  11000, ConstFloatEnd:  11999,
	ConstBoolStart:   12000, ConstBoolEnd:   12999,
}

// MemoryManager keeps track of the next available address for each segment/type.
type MemoryManager struct {
	config Configuration

	nextGlobalInt, nextGlobalFloat, nextGlobalBool   int
	nextLocalInt, nextLocalFloat, nextLocalBool      int
	nextTempInt, nextTempFloat, nextTempBool         int
	nextConstInt, nextConstFloat, nextConstBool      int
}

// NewMemoryManager creates a memory manager with the given config.
func NewMemoryManager(config Configuration) *MemoryManager {
	return &MemoryManager{
		config: config,

		nextGlobalInt:   config.GlobalIntStart,
		nextGlobalFloat: config.GlobalFloatStart,
		nextGlobalBool:  config.GlobalBoolStart,

		nextLocalInt:    config.LocalIntStart,
		nextLocalFloat:  config.LocalFloatStart,
		nextLocalBool:   config.LocalBoolStart,

		nextTempInt:     config.TempIntStart,
		nextTempFloat:   config.TempFloatStart,
		nextTempBool:    config.TempBoolStart,

		nextConstInt:    config.ConstIntStart,
		nextConstFloat:  config.ConstFloatStart,
		nextConstBool:   config.ConstBoolStart,
	}
}

// Métodos para Globales
func (m *MemoryManager) NextGlobalInt() int {
	addr := m.nextGlobalInt
	m.nextGlobalInt++
	if addr > m.config.GlobalIntEnd {
		panic("Global int memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextGlobalFloat() int {
	addr := m.nextGlobalFloat
	m.nextGlobalFloat++
	if addr > m.config.GlobalFloatEnd {
		panic("Global float memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextGlobalBool() int {
	addr := m.nextGlobalBool
	m.nextGlobalBool++
	if addr > m.config.GlobalBoolEnd {
		panic("Global bool memory overflow")
	}
	return addr
}

// Métodos para Locales
func (m *MemoryManager) NextLocalInt() int {
	addr := m.nextLocalInt
	m.nextLocalInt++
	if addr > m.config.LocalIntEnd {
		panic("Local int memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextLocalFloat() int {
	addr := m.nextLocalFloat
	m.nextLocalFloat++
	if addr > m.config.LocalFloatEnd {
		panic("Local float memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextLocalBool() int {
	addr := m.nextLocalBool
	m.nextLocalBool++
	if addr > m.config.LocalBoolEnd {
		panic("Local bool memory overflow")
	}
	return addr
}

// Métodos para Temporales
func (m *MemoryManager) NextTempInt() int {
	addr := m.nextTempInt
	m.nextTempInt++
	if addr > m.config.TempIntEnd {
		panic("Temporal int memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextTempFloat() int {
	addr := m.nextTempFloat
	m.nextTempFloat++
	if addr > m.config.TempFloatEnd {
		panic("Temporal float memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextTempBool() int {
	addr := m.nextTempBool
	m.nextTempBool++
	if addr > m.config.TempBoolEnd {
		panic("Temporal bool memory overflow")
	}
	return addr
}

// Métodos para Constantes
func (m *MemoryManager) NextConstInt() int {
	addr := m.nextConstInt
	m.nextConstInt++
	if addr > m.config.ConstIntEnd {
		panic("Constant int memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextConstFloat() int {
	addr := m.nextConstFloat
	m.nextConstFloat++
	if addr > m.config.ConstFloatEnd {
		panic("Constant float memory overflow")
	}
	return addr
}
func (m *MemoryManager) NextConstBool() int {
	addr := m.nextConstBool
	m.nextConstBool++
	if addr > m.config.ConstBoolEnd {
		panic("Constant bool memory overflow")
	}
	return addr
}

// Puedes agregar métodos para strings, arreglos, etc. siguiendo el mismo patrón.

