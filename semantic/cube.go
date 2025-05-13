package semantic

// Structure of 3 levels
// operator ->  left ->  right -> result
type SemanticCube map[string]map[string]map[string]string

// valid combinations between operators
var Cube = SemanticCube{
	// Operaciones aritméticas
	"+": {
		"int": {
			"int":   "int",
			"float": "float",
		},
		"float": {
			"int":   "float",
			"float": "float",
		},
	},
	"-": {
		"int": {
			"int":   "int",
			"float": "float",
		},
		"float": {
			"int":   "float",
			"float": "float",
		},
	},
	"*": {
		"int": {
			"int":   "int",
			"float": "float",
		},
		"float": {
			"int":   "float",
			"float": "float",
		},
	},
	"/": {
		"int": {
			"int":   "float",
			"float": "float",
		},
		"float": {
			"int":   "float",
			"float": "float",
		},
	},

	// Operaciones relacionales (<, >, !=)
	"<": {
		"int": {
			"int":   "bool",
			"float": "bool",
		},
		"float": {
			"int":   "bool",
			"float": "bool",
		},
	},
	">": {
		"int": {
			"int":   "bool",
			"float": "bool",
		},
		"float": {
			"int":   "bool",
			"float": "bool",
		},
	},
	"!=": {
		"int": {
			"int":   "bool",
			"float": "bool",
		},
		"float": {
			"int":   "bool",
			"float": "bool",
		},
	},
}
