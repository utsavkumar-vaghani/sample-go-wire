package db

import (
	"context"
	"os"

	"example.com/sample-go-wire/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection() *mongo.Database {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB_URI")) 
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.GetLogger().Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.GetLogger().Fatal(err)
	}

	logger.GetLogger().Infof("connected to MongoDB on: %v", os.Getenv("MONGO_DB_URI"))

	return client.Database(os.Getenv("MONGO_DB_NAME"))
}
