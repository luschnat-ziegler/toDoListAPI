package mocks

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DummyListValid = domain.ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "Dummy List Name",
	Description: nil,
	Tasks:       []domain.Task{
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

var DummyListValidAsJSON = `{"id":"000000000000000000000000","name":"Dummy List Name","description":null,"tasks":[{"id":"","name":"Dummy Task 1","description":null},{"id":"","name":"Dummy Task 2","description":null}]}`

var objectId, _ = primitive.ObjectIDFromHex("601be448b9b5e15374b1e842")

var DummyListValidWithIds = domain.ToDoList{
	Id:          objectId,
	Name:        "Dummy List Name",
	Description: nil,
	Tasks:       []domain.Task{
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

var DummyListValidWithIdsAsJson = `{"id":"601be448b9b5e15374b1e842","name":"Dummy List Name","description":null,"tasks":[{"id":"1234","name":"Dummy Task 1","description":null},{"id":"3245","name":"Dummy Task 2","description":null}]}`
var DummyInvalidJSON = `{id":"601be448b9b5e15374b1e842","name":"Dummy List Name","description":null,"tasks":[{"id":"1234","name":"Dummy Task 1","description":null},{"id":"3245","name":"Dummy Task 2","description":null}]}`
var DummyListInvalid = domain.ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "",
	Description: nil,
	Tasks:       []domain.Task{
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

var DummyAppError = errs.NewInternalError("internal error")
var DummyAppErrorAsJSON = `{"message":"internal error"}`

var DummyBadRequestErrorAsJSON = `{"message":"Body parsing error"}`

var ValidSaveListRequest = `{"name":"Dummy List Name", "description":null, "tasks":[{"name":"Dummy Task 1","description":null},{"name":"Dummy Task 2","description":null}]}`
var InvalidSaveListRequest = `{"name":"", "description":null, "tasks":[{"name":"Dummy Task 1","description":null},{"name":"Dummy Task 2","description":null}]}`

var DummyValidationErrorAsJSON = `{"invalid_fields":{"name":"required"}}`