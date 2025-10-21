package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Connect() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		panic("MONGO_URI not found in environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ MongoDB connection failed:", err)
	}

	if err := c.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	client = c
	log.Println("✅ Connected to MongoDB")
}

func GetCollection(dbName, collName string) *mongo.Collection {
	if client == nil {
		log.Fatal("Mongo client not initialized")
	}
	return client.Database(dbName).Collection(collName)
}
