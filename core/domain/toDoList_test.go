package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var dummyListWithNonZeroId = ToDoList{
	Id:          primitive.NewObjectID(),
	Name:        "Dummy List Name",
	Description: nil,
	Tasks:       dummyTasksWithNonZeroIds,
}

var dummyListWithZeroId = ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "Dummy List Name",
	Description: nil,
	Tasks:       dummyTasksWithNonZeroIds,
}

var dummyListWithNonZeroTaskIds = ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "Dummy List Name",
	Description: nil,
	Tasks:       dummyTasksWithNonZeroIds,
}

var dummyListWithZeroTaskIds = ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "Dummy List Name",
	Description: nil,
	Tasks:       dummyTasksWithZeroIds,
}

var dummyTasksWithNonZeroIds = []Task{
	{
		Id:          "1234",
		Name:        "Dummy Task 1",
		Description: nil,
	},
	{
		Id:          "2345",
		Name:        "Dummy Task 2",
		Description: nil,
	},
}

var dummyTasksWithZeroIds = []Task{
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
}

func Test_ToDoList_ResetID_should_set_or_reset_ID_to_zero_value(t *testing.T) {

	zeroID := primitive.ObjectID{}

	dummyListWithZeroId.ResetID()
	if dummyListWithZeroId.Id != zeroID {
		t.Error("Id does not match zero value")
	}

	dummyListWithNonZeroId.ResetID()
	if dummyListWithNonZeroId.Id != zeroID {
		t.Error("Id does not match zero value")
	}
}

func Test_ToDoList_AssignTaskIDs_should_set_or_reset_all_task_ids_to_new_values(t *testing.T) {
	
}
