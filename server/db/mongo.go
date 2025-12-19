package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo() {
	var err error
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println("Warning: MongoDB failed connect setup:", err)
	} else {
		if err = MongoClient.Ping(context.Background(), nil); err != nil {
			log.Println("Warning: MongoDB not reachable. Ensure MongoDB is running.", err)
		} else {
			fmt.Println("MongoDB Connected")
		}
	}
}
