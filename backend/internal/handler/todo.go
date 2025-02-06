package handler

import (
	"encoding/json"
	"net/http"

	"go-cloud-run/internal/model"
	"go-cloud-run/internal/repository"
)

// TodoHandler はTODOのHTTPリクエストを処理する構造体
type TodoHandler struct {
	repo *repository.TodoRepository
}

// NewTodoHandler は新しいTodoHandlerを作成する
func NewTodoHandler(repo *repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

// HandleTodos は/api/todosへのリクエストを処理する
func (h *TodoHandler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTodos(w, r)
	case http.MethodPost:
		h.createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleTodo は/api/todoへのリクエストを処理する
func (h *TodoHandler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	todoID := r.URL.Query().Get("id")
	if todoID == "" {
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getTodo(w, todoID)
	case http.MethodPut:
		h.updateTodo(w, r, todoID)
	case http.MethodDelete:
		h.deleteTodo(w, todoID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) getTodo(w http.ResponseWriter, id string) {
	todo, err := h.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) updateTodo(w http.ResponseWriter, r *http.Request, id string) {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) deleteTodo(w http.ResponseWriter, id string) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
