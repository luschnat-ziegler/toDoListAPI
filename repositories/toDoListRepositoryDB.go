package repositories

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
	"github.com/luschnat-ziegler/toDoListAPI/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoListRepositoryDB struct{}

func (toDoListRepositoryDB ToDoListRepositoryDB) GetAll() (*[]domain.ToDoList, *errs.AppError) {
	if err := ConnectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Error("Error querying database")
		return nil, errs.NewUnexpectedError("Database Error: " + err.Error())
	}

	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			logger.Error("Error closing cursor: " + err.Error())
		}
	}()

	var output []domain.ToDoList

	for cursor.Next(ctx) {
		var toDoList domain.ToDoList
		err := cursor.Decode(&toDoList)
		if err != nil {
			logger.Error("Error decoding database object: " + err.Error())
			return nil, errs.NewUnexpectedError("Database Error")
		}
		output = append(output, toDoList)
	}

	return &output, nil
}

func (toDoListRepositoryDB ToDoListRepositoryDB) GetOneById(id string) (*domain.ToDoList, *errs.AppError) {
	if err := ConnectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Error parsing id: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}

	var toDoList domain.ToDoList

	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&toDoList)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}
	return &toDoList, nil
}
