/*
 * Package: domain_test
 * --------------------
 * Includes test of domain model type methods.
 * Note: Excluded from package domain in order to prevent circular imports.
 */

package domain_test

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/testUtils/dummies"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"testing"
)

/*
 * Function: Test_ToDoList_ResetID_should_set_or_reset_ID_to_zero_value
 * --------------------
 * Tests functionality of ToDoList.ResetID by checking if zero value is inserted in field id of dummies.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 */

func Test_ToDoList_ResetID_should_set_or_reset_ID_to_zero_value(t *testing.T) {

	zeroID := primitive.ObjectID{}

	dummyListWithZeroId := domain.ToDoList{
		Id:          primitive.ObjectID{},
		Name:        "Dummy List Name",
		Description: nil,
		Tasks: []domain.Task{
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
		},
	}

	dummyListWithZeroId.ResetID()
	if dummyListWithZeroId.Id != zeroID {
		t.Error("Id does not match zero value")
	}

	dummyListWithNonZeroId := domain.ToDoList{
		Id:          primitive.NewObjectID(),
		Name:        "Dummy List Name",
		Description: nil,
		Tasks: []domain.Task{
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
		},
	}

	dummyListWithNonZeroId.ResetID()
	if dummyListWithNonZeroId.Id != zeroID {
		t.Error("Id does not match zero value")
	}
}

/*
 * Function: Test_ToDoList_AssignTaskIDs_should_set_or_reset_all_task_ids_to_new_values
 * --------------------
 * Tests functionality of ToDoList.AssignTaskIDs by checking if new ids are inserted in field id of
 * dummies' slice of domain.Task.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 */

func Test_ToDoList_AssignTaskIDs_should_set_or_reset_all_task_ids_to_new_values(t *testing.T) {

	dummyListNonZeroTaskIds := domain.ToDoList{
		Id:          primitive.ObjectID{},
		Name:        "Dummy List Name",
		Description: nil,
		Tasks: []domain.Task{
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
		},
	}
	initialIds := [2]string{"1234", "2345"}

	dummyListNonZeroTaskIds.AssignTaskIDs()
	for i := range dummyListNonZeroTaskIds.Tasks {
		if dummyListNonZeroTaskIds.Tasks[i].Id == initialIds[i] {
			t.Error("Expected new uuid string, got zero value instead")
		}
	}

	var dummyListZeroTaskIds = domain.ToDoList{
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

	dummyListZeroTaskIds.AssignTaskIDs()

	for i := range dummyListZeroTaskIds.Tasks {
		if dummyListZeroTaskIds.Tasks[i].Id == "" {
			t.Error("Expected new uuid string, got zero value instead")
		}
	}
}

/*
 * Function: Test_ToDoList_Validate_should_return_nil_if_provided_with_valid_list
 * --------------------
 * Tests functionality of ToDoList.Validate by calling method on valid dummy list.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 */

func Test_ToDoList_Validate_should_return_nil_if_provided_with_valid_list(t *testing.T) {
	err := dummies.DummyListValid.Validate()
	if err != nil {
		t.Error("Expected nil, got validation error instead")
	}
}

/*
 * Function: Test_ToDoList_Validate_should_return_matching_validation_error_with_invalid_list
 * --------------------
 * Tests functionality of ToDoList.Validate by calling method on invalid dummy list and comparing
 * the resulting errs.ValidationErrors field with the correct values
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 */

func Test_ToDoList_Validate_should_return_matching_validation_error_with_invalid_list(t *testing.T) {
	err := dummies.DummyListInvalid.Validate()
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
