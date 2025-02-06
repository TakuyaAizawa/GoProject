package handler

import (
	"encoding/json"
	"net/http"
)

// Response は一般的なレスポンス構造体
type Response struct {
	Message string `json:"message"`
}

// HealthHandler はヘルスチェックリクエストを処理する構造体
type HealthHandler struct{}

// NewHealthHandler は新しいHealthHandlerを作成する
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HandleHealth は/healthへのリクエストを処理する
func (h *HealthHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
