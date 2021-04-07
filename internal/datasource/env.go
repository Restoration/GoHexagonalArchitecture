package datasource

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// イニシャライズ
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// GetEnv 環境変数取得用関数
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
