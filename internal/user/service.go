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

	existsErr := CheckUserExists(user)
	if existsErr != nil {
		return "", existsErr
	}

	hashedPassword, hashErr := utils.GenerateHash(user.Password)
	if hashErr != nil {
		return "", hashErr
	}
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	if user.Role == "" {
		user.Role = "user"
	}

	newUser, createErr := getUserCollection().InsertOne(ctx, user)
	if createErr != nil {
		return "", createErr
	}

	return newUser.InsertedID.(bson.ObjectID).Hex(), nil
}

func FindUserByKey(key string, value any) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	dbErr := getUserCollection().FindOne(ctx, bson.M{key: value}).Decode(&user)
	if dbErr != nil {
		if errors.Is(dbErr, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, dbErr
	}

	return &user, nil
}

func FindUserByEmail(email string) (*User, error) {
	return FindUserByKey("email", email)
}

func FindUserById(id string) (*User, error) {
	if !utils.IsValidObjectID(id) {
		return nil, errors.New("invalid ObjectId format")
	}
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return FindUserByKey("_id", objectId)
}

func ValidateUserAccess(userId, requestingUserId string) error {
	if requestingUserId != userId {
		return errors.New("unauthorized")
	}
	return nil
}

func GetUserProfileService(userId string) (*User, error) {
	user, err := FindUserById(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
