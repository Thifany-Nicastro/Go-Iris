package controllers

import (
	"errors"
	"fmt"
	"go-iris/dtos"
	"go-iris/models"
	"go-iris/services"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type TodoController struct {
	Service services.TodoService
}

func (c *TodoController) Get() []dtos.TodoResponse {
	todos := c.Service.GetTodos()

	return todos
}

func (c *TodoController) GetBy(id string) mvc.Result {
	todo, err := c.Service.FindTodo(id)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusNotFound,
			Object: map[string]any{
				"message": "Todo not found",
			},
		}
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: todo,
	}
}

func (c *TodoController) Post(request models.Todo) mvc.Result {
	id, err := c.Service.CreateTodo(request)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusBadRequest,
			Object: map[string]any{
				"message": err.Error(),
			},
		}
	}

	return mvc.Response{
		Code:   iris.StatusCreated,
		Object: id,
	}
}

func (c *TodoController) PutBy(id string, request models.Todo) mvc.Result {
	modifiedCount := c.Service.UpdateTodo(id, request)

	if modifiedCount == 0 {
		return mvc.Response{
			Code: iris.StatusNotFound,
			Object: map[string]any{
				"message": "Todo not found",
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusOK,
	}
}

func (c *TodoController) DeleteBy(id string) mvc.Result {
	deletedCount := c.Service.DeleteTodo(id)

	if deletedCount == 0 {
		return mvc.Response{
			Code: iris.StatusNotFound,
			Object: map[string]any{
				"message": "Todo not found",
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusNoContent,
	}
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *TodoController) HandleError(ctx iris.Context, err error) {
	code := iris.StatusInternalServerError
	message := "Ops! Something went wrong"

	if errors.As(err, &validator.ValidationErrors{}) {
		code = iris.StatusUnprocessableEntity
		message = err.Error()
	}

	fmt.Println(err.Error())

	ctx.StopWithJSON(
		code,
		ErrorResponse{
			Code:    code,
			Message: message,
		},
	)
}
