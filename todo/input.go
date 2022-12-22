package todo

type TodoInput struct {
	ID          int    `json:"id"`
	Description string `json:"description" binding:"required"`
	Status      bool   `json:"status" binding:"required"`
}
