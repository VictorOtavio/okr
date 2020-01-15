package config

import (
	"os"
)

// GetEnv busca por uma variável de ambiente declarada ou aplica um valor padrão
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
