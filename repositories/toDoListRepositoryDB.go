/*
 * package: repositories
 * --------------------
 * Includes repository implementation(s) (as defined in package ports)
 */

package repositories

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
	"github.com/luschnat-ziegler/toDoListAPI/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoListRepositoryDB struct{}

/*
 * Method: ToDoListRepositoryDB.GetAll
 * --------------------
 * Retrieves all lists from the database.
 *
 * returns: a pointer to a slice of domain.ToDoList and nil on success.
 *          Otherwise, nil and a pointer to an errs.AppError are returned.
 */

func (toDoListRepositoryDB ToDoListRepositoryDB) GetAll() (*[]domain.ToDoList, *errs.AppError) {
	if err := connectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewInternalError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Error("Error querying database")
		return nil, errs.NewInternalError("Database Error: " + err.Error())
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
			return nil, errs.NewInternalError("Database Error")
		}
		output = append(output, toDoList)
	}

	return &output, nil
}

/*
 * Method: ToDoListRepositoryDB.GetOneById
 * --------------------
 * Retrieves one list from the database (by id).
 *
 * id: a string representation of a primitive.ObjectID associated with the requested list
 *
 * returns: a pointer to a domain.ToDoList and nil on success.
 *          Otherwise, nil and a pointer to an errs.AppError are returned.
 */

func (toDoListRepositoryDB ToDoListRepositoryDB) GetOneById(id string) (*domain.ToDoList, *errs.AppError) {
	if err := connectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewInternalError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Error parsing id: " + err.Error())
		return nil, errs.NewBadRequestError("ID is invalid")
	}

	var toDoList domain.ToDoList

	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&toDoList)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, errs.NewInternalError("Database Error")
	}
	return &toDoList, nil
}

/*
 * Method: ToDoListRepositoryDB.UpdateOneById
 * --------------------
 * Overwrites one list in the database (by id). Does not implement upserting.
 *
 * id: a string representation of a primitive.ObjectID associated with the list requested for update.
 * newList: the new domain.ToDoList to overwrite the existing resource with.
 *
 * returns: a pointer to a domain.ToDoList and nil on success.
 *          Otherwise, nil and a pointer to an errs.AppError are returned.
 */

func (toDoListRepositoryDB ToDoListRepositoryDB) UpdateOneById(id string, newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	if err := connectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewInternalError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Error parsing id: " + err.Error())
		return nil, errs.NewInternalError("Database Error")
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
		return nil, errs.NewInternalError("Database Error")
	}

	if res.MatchedCount == 0 {
		return nil, errs.NewNotFoundError("No documents matching id " + id)
	}

	newList.Id = objectId
	return &newList, nil
}

/*
 * Method: ToDoListRepositoryDB.Save
 * --------------------
 * Saves one new list in the database.
 *
 * newList: the new domain.ToDoList to be persisted.
 *
 * returns: a pointer to a domain.ToDoList (the new resource) and nil on success.
 *          Otherwise, nil and a pointer to an errs.AppError are returned.
 */

func (toDoListRepositoryDB ToDoListRepositoryDB) Save(newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	if err := connectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewInternalError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	result, err := collection.InsertOne(ctx, newList)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, errs.NewInternalError("Database error")
	}

	newList.Id = result.InsertedID.(primitive.ObjectID)
	return &newList, nil
}

/*
 * Method: ToDoListRepositoryDB.DeleteOnById
 * --------------------
 * Deletes one list from the database.
 *
 * id: a string representation of a primitive.ObjectID associated with the list requested for deletion.
 *
 * returns: nil on success or a pointer to an errs.AppError on failure
 */

func (toDoListRepositoryDB ToDoListRepositoryDB) DeleteOneById(id string) *errs.AppError {
	if err := connectDbClient(); err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return errs.NewInternalError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Error parsing id: " + err.Error())
		return errs.NewInternalError("Database Error")
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return errs.NewInternalError("Database error")
	}

	if result.DeletedCount == 0 {
		return errs.NewNotFoundError("No Documents matching ID " + id)
	}

	return nil
}

/*
 * Function: NewToDoListRepositoryDB
 * --------------------
 * Instantiates a new ToDoListRepositoryDB for dependency injection.
 *
 * returns: an instance of ToDoListRepositoryDB
 */

func NewToDoListRepositoryDB() ToDoListRepositoryDB {
	return ToDoListRepositoryDB{}
}
