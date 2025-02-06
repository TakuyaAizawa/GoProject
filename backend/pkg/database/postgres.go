package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// NewPostgresDB は新しいPostgreSQLデータベース接続を作成する
func NewPostgresDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// コネクションプールの設定
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// データベース接続のテスト
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Successfully connected to database")

	return db, nil
}
