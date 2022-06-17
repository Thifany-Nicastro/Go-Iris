package controllers

import (
	"context"
	"errors"
	"fmt"
	"go-iris/config"
	"go-iris/models"
	"go-iris/services"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoController struct {
	/* dependencies */
}

func (c *TodoController) Get() []models.Todo {
	var todos []models.Todo

	todosCollection := config.GetCollection("todos")
	cursor, err := todosCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &todos); err != nil {
		log.Fatal(err)
	}

	return todos
}

func (c *TodoController) GetBy(id string) mvc.Result {
	var todo models.Todo

	todosCollection := config.GetCollection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	err := todosCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&todo)
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

func (c *TodoController) Post(todo models.Todo) mvc.Result {
	id, err := services.CreateTodo(todo)
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

func (c *TodoController) PutBy(id string, t models.Todo) mvc.Result {
	todo := bson.M{"title": t.Title}

	todosCollection := config.GetCollection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := todosCollection.UpdateOne(context.TODO(),
		bson.M{"_id": objId},
		bson.M{"$set": todo},
	)
	if err != nil {
		log.Fatal(err)
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: result,
	}
}

func (c *TodoController) DeleteBy(id string) mvc.Result {
	todosCollection := config.GetCollection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, _ := todosCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if result.DeletedCount == 0 {
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
