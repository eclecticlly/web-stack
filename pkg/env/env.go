package env

import (
	"os"

	"github.com/joho/godotenv"
)

// GetEnv with sensible default
func GetEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

// SetupEnvFile if present
func SetupEnvFile() {
	env := GetEnv("APP_ENV", "")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
}
