package user

import (
	"auction_golang_concurrency/configuration/logger"
	"auction_golang_concurrency/internal/entity/user_entity"
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{Collection: database.Collection("users")}
}

func (r *UserRepository) findUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userId}

	var user UserEntityMongo
	err := r.Collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		errMessage := fmt.Sprintf("Error on find user by id %v - error %v", userId, err)

		if err == mongo.ErrNoDocuments {
			errMessage = fmt.Sprintf("user with %s not found", userId)
			logger.Error(errMessage, err)
			return nil, internal_error.NewNotFoundError(errMessage)
		}

		logger.Error("Error on find user by id %v", err)
		return nil, internal_error.NewInternalError(errMessage)
	}

	return &user_entity.User{Id: user.Id, Name: user.Name}, nil
}
