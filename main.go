package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// Task represents a task in the system
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Response represents a generic response structure
type Response struct {
	Message string `json:"message"`
}

// Global database connection pool
var db *sql.DB

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// すべてのオリジンを許可（開発環境用）
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 許可するHTTPメソッド
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 許可するヘッダー
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	// データベース接続の初期化
	// dbURL := os.Getenv("DATABASE_URL")
	// if dbURL != "" {
	// 	var err error
	// 	db, err = sql.Open("postgres", dbURL)
	// 	if err != nil {
	// 		log.Fatal("Database connection failed:", err)
	// 	}
	// 	defer db.Close()

	// 	// データベース接続のテスト
	// 	if err := db.Ping(); err != nil {
	// 		log.Fatal("Database ping failed:", err)
	// 	}
	// 	log.Println("Database connected successfully")
	// }

	// .env ファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// データベース接続の初期化
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	log.Println("Attempting to connect to database...")

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}
	defer db.Close()

	// コネクションプールの設定
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// データベース接続のテスト
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Successfully connected to database")

	// ルーティングの設定
	// http.HandleFunc("/", handleHome)
	// http.HandleFunc("/api/hello", handleHello)
	// http.HandleFunc("/api/tasks", handleTasks)
	// http.HandleFunc("/api/task", handleTask) // 個別のタスク操作
	// http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/", enableCORS(handleHome))
	http.HandleFunc("/api/hello", enableCORS(handleHello))
	http.HandleFunc("/api/tasks", enableCORS(handleTasks))
	http.HandleFunc("/api/task", enableCORS(handleTask))
	http.HandleFunc("/health", enableCORS(handleHealth))
	http.HandleFunc("/api/todos", enableCORS(handleTodos))
	http.HandleFunc("/api/todo", enableCORS(handleTodo))

	// ポート番号の取得（Cloud Run対応）
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Go on Cloud Run!"))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello from Go!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	// データベース接続がない場合はエラーを返す
	if db == nil {
		http.Error(w, "Database connection not available", http.StatusServiceUnavailable)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTask(w, taskID)
	case http.MethodPut:
		updateTask(w, r, taskID)
	case http.MethodDelete:
		deleteTask(w, taskID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// func createTask(w http.ResponseWriter, r *http.Request) {
// 	var task Task
// 	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err := db.QueryRow(
// 		"INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id",
// 		task.Title, task.Description,
// 	).Scan(&task.ID)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusCreated)
//		json.NewEncoder(w).Encode(task)
//	}
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO tasks (title, description, created_at, updated_at) 
        VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
        RETURNING id, created_at, updated_at`

	err := db.QueryRow(query, task.Title, task.Description).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func getTask(w http.ResponseWriter, id string) {
	var task Task
	query := `
        SELECT id, title, description, created_at, updated_at 
        FROM tasks 
        WHERE id = $1`

	err := db.QueryRow(query, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request, id string) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `
        UPDATE tasks 
        SET title = $1, description = $2, updated_at = CURRENT_TIMESTAMP 
        WHERE id = $3 
        RETURNING id, title, description, created_at, updated_at`

	err := db.QueryRow(query, task.Title, task.Description, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, id string) {
	query := "DELETE FROM tasks WHERE id = $1"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Todo represents a todo item in the system
type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "Database connection not available", http.StatusServiceUnavailable)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	todoID := r.URL.Query().Get("id")
	if todoID == "" {
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTodo(w, todoID)
	case http.MethodPut:
		updateTodo(w, r, todoID)
	case http.MethodDelete:
		deleteTodo(w, todoID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func getTodos(w http.ResponseWriter, r *http.Request) {
	// データベース接続の確認
	if db == nil {
		http.Error(w, "Database connection not available", http.StatusServiceUnavailable)
		return
	}

	query := `
        SELECT id, text, created_at, updated_at 
        FROM todos 
        ORDER BY created_at DESC`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []Todo{} // 空の配列で初期化
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Text, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")

	// エラーチェック
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSONエンコード
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO todos (text) 
        VALUES ($1) 
        RETURNING id, text, created_at, updated_at`

	err := db.QueryRow(query, todo.Text).
		Scan(&todo.ID, &todo.Text, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request, id string) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `
        UPDATE todos 
        SET text = $1 
        WHERE id = $2 
        RETURNING id, text, created_at, updated_at`

	err := db.QueryRow(query, todo.Text, id).
		Scan(&todo.ID, &todo.Text, &todo.CreatedAt, &todo.UpdatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, id string) {
	query := "DELETE FROM todos WHERE id = $1"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func getTodo(w http.ResponseWriter, id string) {
	var todo Todo
	query := `
        SELECT id, text, created_at, updated_at 
        FROM todos 
        WHERE id = $1`

	err := db.QueryRow(query, id).
		Scan(&todo.ID, &todo.Text, &todo.CreatedAt, &todo.UpdatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
