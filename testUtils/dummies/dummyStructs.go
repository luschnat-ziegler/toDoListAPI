/*
 * package: dummies
 * --------------------
 * Provides dummy data for unit testing
 */

package dummies

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DummyListValid = domain.ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "Dummy List Name",
	Description: nil,
	Tasks: []domain.Task{
		{
			Id:          "",
			Name:        "Dummy Task 1",
			Description: nil,
		},
		{
			Id:          "",
			Name:        "Dummy Task 2",
			Description: nil,
		},
	},
}

var objectId, _ = primitive.ObjectIDFromHex("601be448b9b5e15374b1e842")
var DummyListValidWithIds = domain.ToDoList{
	Id:          objectId,
	Name:        "Dummy List Name",
	Description: nil,
	Tasks: []domain.Task{
		{
			Id:          "1234",
			Name:        "Dummy Task 1",
			Description: nil,
		},
		{
			Id:          "3245",
			Name:        "Dummy Task 2",
			Description: nil,
		},
	},
}

var DummyListInvalid = domain.ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "",
	Description: nil,
	Tasks: []domain.Task{
		{
			Id:          "",
			Name:        "",
			Description: nil,
		},
		{
			Id:          "",
			Name:        "Dummy Task 2",
			Description: nil,
		},
	},
}

var DummyInternalError = errs.NewInternalError("internal error")
