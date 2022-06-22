package controllers

import (
	"errors"
	"fmt"
	"go-iris/dtos"
	"go-iris/services"
	"go-iris/utils"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type TodoController struct {
	Service services.TodoService
	// Ctx     iris.Context
	User string
}

func (c *TodoController) BeforeActivation(ctx iris.Context, b mvc.BeforeActivation) {
	user := utils.GetAuthenticatedUser(ctx)

	b.Dependencies().Register(user)
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
			Object: iris.Map{
				"message": "Todo not found",
			},
		}
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: todo,
	}
}

func (c *TodoController) Post(request dtos.TodoRequest) mvc.Result {
	// user := utils.GetAuthenticatedUser(c.Ctx)

	id, err := c.Service.CreateTodo(request, c.User)

	if err != nil {
		return mvc.Response{
			Code: iris.StatusBadRequest,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusCreated,
		Object: iris.Map{
			"message": "Todo #" + id + " created",
		},
	}
}

func (c *TodoController) PutBy(id string, request dtos.TodoRequest) mvc.Result {
	modifiedCount := c.Service.UpdateTodo(id, request)

	if modifiedCount == 0 {
		return mvc.Response{
			Code: iris.StatusNotFound,
			Object: iris.Map{
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
			Object: iris.Map{
				"message": "Todo not found",
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusNoContent,
	}
}

func (c *TodoController) PutCompleteBy(id string) mvc.Result {
	modifiedCount := c.Service.CompleteTodo(id)

	if modifiedCount == 0 {
		return mvc.Response{
			Code: iris.StatusNotFound,
			Object: iris.Map{
				"message": "Todo not found",
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusOK,
		Object: iris.Map{
			"message": "Todo completed!",
		},
	}
}

type ErrorResponse struct {
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
			Message: message,
		},
	)
}
