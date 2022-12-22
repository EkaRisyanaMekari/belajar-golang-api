package todo

import "time"

type Todo struct {
	ID          int
	Description string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
