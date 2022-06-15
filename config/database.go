package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

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

func Connect() {
	var err error

	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(getURI()))
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := MongoClient.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// coll := client.Database("sample_mflix").Collection("movies")
	// title := "Back to the Future"
	// var result bson.M

	// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }

	// if err != nil {
	// 	panic(err)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%s\n", jsonData)
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := MongoClient.Database("go-iris").Collection(collectionName)

	return collection
}
