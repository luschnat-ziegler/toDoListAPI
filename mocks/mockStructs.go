package mocks

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
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
