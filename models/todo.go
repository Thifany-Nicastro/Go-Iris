package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	IsCompleted bool               `bson:"is_completed"`
	UserID      primitive.ObjectID `bson:"user_id"`
}
