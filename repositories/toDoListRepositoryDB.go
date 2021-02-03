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
	if err := connectDbClient(); err != nil {
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
	if err := connectDbClient(); err != nil {
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

func (toDoListRepositoryDB ToDoListRepositoryDB) UpdateOneById(id string, newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	if err := connectDbClient(); err != nil {
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

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"name":        newList.Name,
			"description": newList.Description,
			"tasks":       newList.Tasks,
		},
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}

	if res.MatchedCount == 0 {
		return nil, errs.NewNotFoundError("No documents with id " + id)
	}

	newList.Id = objectId
	return &newList, nil
}

func (toDoListRepositoryDB ToDoListRepositoryDB) Save(newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	if err := connectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	var newToDoList domain.ToDoList

	result, err := collection.InsertOne(ctx, newToDoList)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database error")
	}

	newList.Id = result.InsertedID.(primitive.ObjectID)
	return &newList, nil
}


func (toDoListRepositoryDB ToDoListRepositoryDB) DeleteOneById(id string) (*int64, *errs.AppError) {
	if err := connectDbClient(); err != nil {
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

	collection := client.Database(dbName).Collection(collectionName)

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database error")
	}

	if result.DeletedCount == 0 {
		return nil, errs.NewNotFoundError("ID does not match a document")
	}

	return &result.DeletedCount, nil
}

func NewToDoListRepositoryDB() ToDoListRepositoryDB {
	return ToDoListRepositoryDB{}
}