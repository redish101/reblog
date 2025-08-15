package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Dev         bool
	Port        int
	DatabaseURL string
	SecretKey   string
	OwnerEmail  string
)

func Init() {
	godotenv.Load()

	Dev = getAsBool("DEV", false)
	Port = getAsInt("PORT", 3000)
	DatabaseURL = getAsString("DATABASE_URL", "postgres://user:password@localhost:5432/reblog?sslmode=disable")
	SecretKey = getAsString("SECRET_KEY", "reblog")
	OwnerEmail = getAsString("OWNER_EMAIL", "")
}

func getAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func getAsBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func getAsString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
