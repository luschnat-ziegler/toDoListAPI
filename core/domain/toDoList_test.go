package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"testing"
)

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

	dummyListWithZeroTaskIds.AssignTaskIDs()
	for i := range dummyListWithZeroTaskIds.Tasks {
		if dummyListWithZeroTaskIds.Tasks[i].Id == "" {
			t.Error("Expected new uuid string, got zero value instead")
		}
	}

	listCopy := dummyListWithNonZeroTaskIds
	listCopy.AssignTaskIDs()
	for i := range listCopy.Tasks {
		if dummyListWithZeroTaskIds.Tasks[i].Id == listCopy.Tasks[i].Id {
			t.Error("Expected new uuid string, got zero value instead")
		}
	}
}

func Test_ToDoList_Validate_should_return_nil_if_provided_with_valid_list(t *testing.T) {
	err := dummyListValid.Validate()
	if err != nil {
		t.Error("Expected nil, got validation error instead")
	}
}

func Test_ToDoList_Validate_should_return_matching_validation_error_with_invalid_list(t *testing.T) {
	err := dummyListInvalid.Validate()
	if err == nil {
		t.Error("Expected validation error, got nil instead")
		return
	}

	if err.Code != http.StatusBadRequest {
		t.Errorf("Expected error code 400, got %v instead.", err.Code)
	}

	if value, ok := err.InvalidFields["name"]; !ok {
		t.Error(`InvalidFields map is missing key "name".`)
	} else if value != "required" {
		t.Errorf(`InvalidFields map has wrong value associated with key "name". Expected "required", got %v instead.`, value)
	}

	if value, ok := err.InvalidFields["tasks[0].name"]; !ok {
		t.Error(`InvalidFields map is missing key "name".`)
	} else if value != "required" {
		t.Errorf(`InvalidFields map has wrong value associated with key "name". Expected "required", got %v instead.`, value)
	}
}

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

var dummyListValid = dummyListWithZeroTaskIds

var dummyListInvalid = ToDoList{
	Id:          primitive.ObjectID{},
	Name:        "",
	Description: nil,
	Tasks:       dummyTasksWithMissingName,
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

var dummyTasksWithMissingName = []Task{
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
}
