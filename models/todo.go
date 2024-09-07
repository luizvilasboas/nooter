package models

type Todo struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" validate:"required" db:"title"`
	Details string `json:"details" validate:"required" db:"details"`
	Done    bool   `json:"done" validate:"required" db:"done"`
}

type TodoRequest struct {
	Title   string `json:"title" validate:"required"`
	Details string `json:"details" validate:"required"`
	Done    bool   `json:"done" validate:"required"`
}
