package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	SQLClient   *sql.DB
	RedisClient *redis.Client
	MongoClient *mongo.Client
)

func Init() {
	var err error

	// 1. MySQL
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/guvi_db?parseTime=true"
	}
	SQLClient, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open MySQL:", err)
	}
	if err = SQLClient.Ping(); err != nil {
		log.Println("Warning: MySQL not connected. Ensure MySQL is running and database 'guvi_db' exists.", err)
	} else {
		fmt.Println("MySQL Connected")
	}

	// 2. Redis
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisPassword := os.Getenv("REDIS_PASSWORD") // Load password

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // Set password
		DB:       0,
	})
	if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
		log.Println("Warning: Redis not connected. Ensure Redis is running.", err)
	} else {
		fmt.Println("Redis Connected")
	}

	// 3. MongoDB
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
