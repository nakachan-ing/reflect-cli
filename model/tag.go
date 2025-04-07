package model

import (
	"time"
)

type Tag struct {
	ID        string     `json:"id"` // 例: t-001
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	Deleted   bool       `json:"deleted"`
}
