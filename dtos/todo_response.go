package dtos

import "go-iris/models"

type TodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
	UserID      string `json:"user_id"`
}

func CreateTodoResponse(todo models.Todo) TodoResponse {
	return TodoResponse{
		ID:          todo.ID.Hex(),
		Title:       todo.Title,
		IsCompleted: todo.IsCompleted,
		UserID:      todo.UserID.Hex(),
	}
}

func CreateTodoListResponse(todos []models.Todo) []TodoResponse {
	var response []TodoResponse

	for _, t := range todos {
		todo := CreateTodoResponse(t)
		response = append(response, todo)
	}

	return response
}
