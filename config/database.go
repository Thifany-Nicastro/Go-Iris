package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var MongoClient *mongo.Client

func getURI() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	return uri
}

func Connect() *mongo.Client {
	var err error

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getURI()))
	if err != nil {
		log.Fatal("No connection :(")
	}
	// defer func() {
	// 	if err = mongoClient.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return mongoClient
}

// func GetCollection(db *mongo.Client, collectionName string) *mongo.Collection {
// 	collection := db.Database("go-iris").Collection(collectionName)

// 	return collection
// }
