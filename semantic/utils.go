package semantic

import (
	"strconv"
	"strings"
)

func InferTypeFromConstant(val string) string {
	// Intenta convertir a int
	if _, err := strconv.Atoi(val); err == nil {
		return "int"
	}

	// Intenta convertir a float
	if _, err := strconv.ParseFloat(val, 64); err == nil {
		return "float"
	}

	// Verifica si es booleano
	if val == "true" || val == "false" {
		return "bool"
	}

	// Si empieza y termina con comillas, es string
	if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
		return "string"
	}

	// Tipo desconocido
	return "error"
}
