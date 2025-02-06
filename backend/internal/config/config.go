package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config はアプリケーションの全設定を保持する構造体
type Config struct {
	DatabaseURL string
	Port        string
}

// Load は環境変数から設定を読み込む
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	if config.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	return config
}
