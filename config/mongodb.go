package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	username := os.Getenv("MONGODB_USER")
	pswd := os.Getenv("MONGODB_PASSWORD")
	host := os.Getenv("MONGODB_HOST")
	port := os.Getenv("MONGODB_PORT")
	dbname := os.Getenv("MONGODB_DB")
	mongodburl := fmt.Sprintf("mongodb://%v:%v@%s:%v/%v", username, pswd, host, port, dbname)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodburl).SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("books")
	return db, cancel, nil
}
