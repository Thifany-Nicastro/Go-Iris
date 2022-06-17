package controllers

import (
	"context"
	"go-iris/config"
	"go-iris/models"
	"go-iris/services"
	"log"

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

// func (c *TodoController) HandleError(ctx iris.Context, err error) {
// 	if iris.IsErrPath(err) {
// 		// to ignore any "schema: invalid path" you can check the error type
// 		// and don't stop the execution.
// 		ctx.WriteString(err.Error())
// 		return // continue.
// 	}

// 	// ctx.WriteString(err.Error())
// 	// ctx.StopExecution()
// 	ctx.StopWithError(iris.StatusBadGateway, errors.New("bb"))
// }

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *TodoController) HandleHTTPError(err mvc.Err, statusCode mvc.Code) ErrorResponse {
	/* OR
	err := ctx.GetErr()
	code := ctx.GetStatusCode()
	*/
	code := int(statusCode)
	msg := ""
	if err != nil {
		msg = err.Error()
	} else {
		msg = iris.StatusText(code)
	}

	return ErrorResponse{
		Code:    code,
		Message: msg,
	}
}

func (c *TodoController) HandleError(ctx iris.Context, err error) {
	// ctx.StopWithError(iris.StatusBadGateway, errors.New("abc"))

	ctx.StopWithJSON(
		iris.StatusBadGateway,
		c.HandleHTTPError(err, iris.StatusBadGateway),
	)
}
