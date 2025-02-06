package main

import (
	"log"
	"net/http"

	"go-cloud-run/internal/config"
	"go-cloud-run/internal/handler"
	"go-cloud-run/internal/middleware"
	"go-cloud-run/internal/repository"
	"go-cloud-run/pkg/database"
)

func main() {
	// 設定の読み込み
	cfg := config.Load()

	// データベース接続の初期化
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// リポジトリの初期化
	taskRepo := repository.NewTaskRepository(db)
	todoRepo := repository.NewTodoRepository(db)

	// ハンドラーの初期化
	taskHandler := handler.NewTaskHandler(taskRepo)
	todoHandler := handler.NewTodoHandler(todoRepo)
	healthHandler := handler.NewHealthHandler()

	// ルーティングの設定
	http.HandleFunc("/", middleware.EnableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Welcome to Go on Cloud Run!"))
	}))

	http.HandleFunc("/api/tasks", middleware.EnableCORS(taskHandler.HandleTasks))
	http.HandleFunc("/api/task", middleware.EnableCORS(taskHandler.HandleTask))
	http.HandleFunc("/api/todos", middleware.EnableCORS(todoHandler.HandleTodos))
	http.HandleFunc("/api/todo", middleware.EnableCORS(todoHandler.HandleTodo))
	http.HandleFunc("/health", middleware.EnableCORS(healthHandler.HandleHealth))

	// サーバーの起動
	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
