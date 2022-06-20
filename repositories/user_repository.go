package repositories

import (
	"context"
	"go-iris/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db         *mongo.Client
	collection *mongo.Collection
}

type UserRepository interface {
	All() []models.User
	FindById(id primitive.ObjectID) (models.User, error)
	Create(user models.User) (interface{}, error)
	Update(id primitive.ObjectID, fields primitive.M) int64
	Delete(id primitive.ObjectID) int64
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		db:         db,
		collection: db.Database("go-iris").Collection("users"),
	}
}

func (s *userRepository) All() []models.User {
	var users []models.User

	cursor, _ := s.collection.Find(context.TODO(), bson.D{{}})

	cursor.All(context.TODO(), &users)

	return users
}

func (s *userRepository) FindById(id primitive.ObjectID) (models.User, error) {
	var user models.User

	err := s.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return user, err
}

func (s *userRepository) Create(user models.User) (interface{}, error) {
	return s.collection.InsertOne(context.TODO(), user)
}

func (s *userRepository) Update(id primitive.ObjectID, fields primitive.M) int64 {
	result, _ := s.collection.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": fields},
	)

	return result.ModifiedCount
}

func (s *userRepository) Delete(id primitive.ObjectID) int64 {
	result, _ := s.collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	return result.DeletedCount
}
