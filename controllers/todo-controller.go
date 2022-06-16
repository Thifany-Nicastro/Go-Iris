package controllers

import (
	"context"
	"go-iris/config"
	"go-iris/models"
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

func (c *TodoController) Post(t models.Todo) mvc.Result {
	todosCollection := config.GetCollection("todos")
	result, err := todosCollection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}

	return mvc.Response{
		Code:   iris.StatusCreated,
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

func (c *TodoController) PutBy(id string, t models.Todo) mvc.Result {
	todo := bson.M{"Title": t.Title}

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
