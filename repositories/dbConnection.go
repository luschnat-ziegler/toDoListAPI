/*
 * package: repositories
 * --------------------
 * Includes repository implementation(s) (as defined in ports)
 */

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
var collection *mongo.Collection

const (
	dbName         = "todo"
	collectionName = "lists"
)

/*
 * Function: connectDbClient
 * --------------------
 * Initiates a connection to mongoDB and sets the collection to the provided collectionName. Also instantiates a
 * context.Context (5s timeout) to be used with the client, as well as the respective context.CancelFunc.
 *
 * returns: An error or nil
 */

func connectDbClient() error {

	url, _ := os.LookupEnv("DB_URL")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	client, clientError := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if clientError != nil {
		logger.Error("Database init error: " + clientError.Error())
		return clientError
	}

	collection = client.Database(dbName).Collection(collectionName)

	return nil
}

/*
 * Function: disconnectClient
 * --------------------
 * Disconnects a mongo.Client.
 *
 * client: a pointer to the mongo.Client to be disconnected
 * ctx: a context.Context used with the client
 *
 * returns: nothing
 */

func disconnectClient(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		logger.Error("Error disconnecting from db client: " + err.Error())
	}
}
