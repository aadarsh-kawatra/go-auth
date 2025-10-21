package user

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"crud/infrastructure/db"
	"crud/pkg/utils"
)

func getUserCollection() *mongo.Collection {
	mongoDbName := os.Getenv("MONGO_DB")
	if mongoDbName == "" {
		panic("MONGO_DB not found in env")
	}

	userCollection := db.GetCollection(mongoDbName, "users")
	return userCollection
}

func CheckUserExists(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser User
	existErr := getUserCollection().FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if existErr == nil {
		return errors.New("email already registered")
	}
	if !errors.Is(existErr, mongo.ErrNoDocuments) {
		return existErr
	}

	return nil
}

func CreateUser(user User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hashedPassword, hashErr := utils.GenerateHash(user.Password)
	if hashErr != nil {
		return "", hashErr
	}
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	newUser, createErr := getUserCollection().InsertOne(ctx, user)
	if createErr != nil {
		return "", createErr
	}

	return newUser.InsertedID.(bson.ObjectID).Hex(), nil
}

func FindUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	dbErr := getUserCollection().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if dbErr != nil {
		if errors.Is(dbErr, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, dbErr
	}

	return &user, nil
}
