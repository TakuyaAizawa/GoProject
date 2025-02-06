package repository

import (
	"database/sql"

	"go-cloud-run/internal/model"
)

// TaskRepository はタスクのデータベース操作を扱う構造体
type TaskRepository struct {
	db *sql.DB
}

// NewTaskRepository は新しいTaskRepositoryを作成する
func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// GetAll は全てのタスクを取得する
func (r *TaskRepository) GetAll() ([]model.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// Create は新しいタスクを作成する
func (r *TaskRepository) Create(task *model.Task) error {
	query := `
		INSERT INTO tasks (title, description, created_at, updated_at) 
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query, task.Title, task.Description).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

// Get は指定されたIDのタスクを取得する
func (r *TaskRepository) Get(id string) (*model.Task, error) {
	var task model.Task
	err := r.db.QueryRow("SELECT id, title, description, created_at, updated_at FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Update は既存のタスクを更新する
func (r *TaskRepository) Update(id string, task *model.Task) error {
	query := `
		UPDATE tasks 
		SET title = $1, description = $2, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $3 
		RETURNING created_at, updated_at`

	return r.db.QueryRow(query, task.Title, task.Description, id).
		Scan(&task.CreatedAt, &task.UpdatedAt)
}

// Delete はタスクを削除する
func (r *TaskRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
