package dtos

type TodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}

type TodoRequest struct {
	Title string `json:"title" validate:"required"`
}
