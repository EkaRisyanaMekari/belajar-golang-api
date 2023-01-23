package todo

import "time"

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description" binding:"required"`
	Status      bool   `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
