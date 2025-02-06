package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config はアプリケーションの全設定を保持する構造体
type Config struct {
	DatabaseURL string
	Port        string
}

// Load は環境変数から設定を読み込む
func Load() *Config {
	// プロジェクトルートの.envファイルを探す
	envPath := filepath.Join(".", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: .env file not found at %s", envPath)
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
