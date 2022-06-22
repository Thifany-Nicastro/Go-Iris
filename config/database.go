package config

import (
	"context"
	"go-iris/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getURI() string {
	return utils.GetEnvVar("MONGODB_URI")
}

func Connect() *mongo.Client {
	var err error

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getURI()))
	if err != nil {
		log.Fatal("No connection")
	}

	return mongoClient
}
