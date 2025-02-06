package model

import "time"

// Todo はシステム内のTODOアイテムを表現する構造体
type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
