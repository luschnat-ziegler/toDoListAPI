package repositories

import (
	"context"
	"github.com/luschnat-ziegler/toDoListAPI/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var clientError error
var collection *mongo.Collection

const (
	dbName         = "todo"
	collectionName = "lists"
)

func connectDbClient() error {

	url, _ := os.LookupEnv("DB_URL")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	client, clientError = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if clientError != nil {
		logger.Error("Database init error: " + clientError.Error())
		return clientError
	}

	collection = client.Database(dbName).Collection(collectionName)

	return nil
}

func disconnectClient(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		logger.Error("Error disconnecting from db client: " + err.Error())
	}
}