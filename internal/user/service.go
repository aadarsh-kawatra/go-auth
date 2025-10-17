package user

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"crud/internal/db"
)

func FindUserByEmail(email string) (*User, error) {
	mongoDbName := os.Getenv("MONGO_DB")
	if mongoDbName == "" {
		panic("MONGO_DB not found in env")
	}

	userCollection := db.GetCollection(mongoDbName, "users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	dbErr := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if dbErr != nil {
		if errors.Is(dbErr, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, dbErr
	}

	return &user, nil
}
