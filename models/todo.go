package models

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title" validate:"required"`
	Details string `json:"details" validate:"required"`
	Done    bool   `json:"done" validate:"required"`
}

type TodoRequest struct {
	Title   string `json:"title" validate:"required"`
	Details string `json:"details" validate:"required"`
	Done    bool   `json:"done" validate:"required"`
}
