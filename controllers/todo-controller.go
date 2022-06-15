package controllers

import (
	"context"
	"go-iris/config"
	"go-iris/models"
	"log"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
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

func (c *TodoController) Post(t models.Todo) int {
	todosCollection := config.GetCollection("todos")
	_, err := todosCollection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}
	// println("Received Todo: " + t.Title)

	return iris.StatusCreated
}

func (c *TodoController) Delete(t models.Todo) int {
	println("Received Todo: " + t.Title)

	return iris.StatusCreated
}

func (c *TodoController) GetBy(id string) int {
	println("Received Todo: " + id)

	return iris.StatusCreated
}
