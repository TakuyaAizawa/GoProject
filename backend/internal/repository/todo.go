package repository

import (
	"database/sql"

	"go-cloud-run/internal/model"
)

// TodoRepository はTODOのデータベース操作を扱う構造体
type TodoRepository struct {
	db *sql.DB
}

// NewTodoRepository は新しいTodoRepositoryを作成する
func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// GetAll は全てのTODOを取得する
func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	rows, err := r.db.Query("SELECT id, text, created_at, updated_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Text, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

// Create は新しいTODOを作成する
func (r *TodoRepository) Create(todo *model.Todo) error {
	query := `
		INSERT INTO todos (text, created_at, updated_at) 
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query, todo.Text).
		Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
}

// Get は指定されたIDのTODOを取得する
func (r *TodoRepository) Get(id string) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.QueryRow("SELECT id, text, created_at, updated_at FROM todos WHERE id = $1", id).
		Scan(&todo.ID, &todo.Text, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// Update は既存のTODOを更新する
func (r *TodoRepository) Update(id string, todo *model.Todo) error {
	query := `
		UPDATE todos 
		SET text = $1, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $2 
		RETURNING created_at, updated_at`

	return r.db.QueryRow(query, todo.Text, id).
		Scan(&todo.CreatedAt, &todo.UpdatedAt)
}

// Delete はTODOを削除する
func (r *TodoRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}
