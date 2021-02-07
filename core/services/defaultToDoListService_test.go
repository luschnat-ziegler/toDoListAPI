/*
 * package: services
 * --------------------
 * Includes service implementation(s) (as defined in package ports)
 */

package services

import (
	"github.com/golang/mock/gomock"
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	ports2 "github.com/luschnat-ziegler/toDoListAPI/core/ports"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
	"github.com/luschnat-ziegler/toDoListAPI/testUtils/mocks/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

var mockToDoListRepository *ports.MockToDoListRepository
var defaultToDoListService ports2.ToDoListService

/*
 * function: setupToDoListServiceTest
 * --------------------
 * Sets up variables for tests and returns teardown function.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: a function to clean up test variables.
 */

func setupToDoListServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockToDoListRepository = ports.NewMockToDoListRepository(ctrl)
	defaultToDoListService = NewToDoListService(mockToDoListRepository)
	return func() {
		defaultToDoListService = nil
		defer ctrl.Finish()
	}
}

/*
 * function: Test_DefaultToDoListService_GetAllLists_should_return_lists_returned_by_repo_method
 * --------------------
 * Tests if pointer to slice of domain.ToDoList from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_GetAllLists_should_return_lists_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockToDoLists := []domain.ToDoList{
		{
			Id:          primitive.ObjectID{},
			Name:        "mock list",
			Description: nil,
			Tasks: []domain.Task{
				{
					Id:          "test_id",
					Name:        "test task name",
					Description: nil,
				},
			},
		},
	}

	mockToDoListRepository.EXPECT().GetAll().Return(&mockToDoLists, nil).Times(1)

	lists, err := defaultToDoListService.GetAllLists()

	if err != nil {
		t.Errorf("Nil expected, error returned: %v", err.Code)
	}

	if !reflect.DeepEqual(*lists, mockToDoLists) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_GetAllLists_should_return_error_returned_by_repo_method
 * --------------------
 * Tests if pointer to errs.AppError from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_GetAllLists_should_return_error_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockAppError := errs.NewInternalError("test error")
	mockToDoListRepository.EXPECT().GetAll().Return(nil, mockAppError)

	_, err := defaultToDoListService.GetAllLists()

	if err == nil {
		t.Error("Error expected, nil returned")
		return
	}

	if !reflect.DeepEqual(*err, *mockAppError) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_SaveList_should_return_list_returned_by_repo_method
 * --------------------
 * Tests if pointer to domain.ToDoList from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_SaveList_should_return_list_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockToDoList := domain.ToDoList{
		Id:          primitive.ObjectID{},
		Name:        "mock list",
		Description: nil,
		Tasks: []domain.Task{
			{
				Id:          "test_id",
				Name:        "test task name",
				Description: nil,
			},
		},
	}

	mockToDoListRepository.EXPECT().Save(mockToDoList).Return(&mockToDoList, nil)

	list, err := defaultToDoListService.SaveList(mockToDoList)

	if err != nil {
		t.Errorf("Nil expected, error returned: %v", err.Code)
	}

	if !reflect.DeepEqual(*list, mockToDoList) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_SaveList_should_return_error_returned_by_repo_method
 * --------------------
 * Tests if pointer to errs.AppError from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_SaveList_should_return_error_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockToDoList := domain.ToDoList{
		Id:          primitive.ObjectID{},
		Name:        "mock list",
		Description: nil,
		Tasks: []domain.Task{
			{
				Id:          "test_id",
				Name:        "test task name",
				Description: nil,
			},
		},
	}
	mockAppError := errs.NewInternalError("test error")
	mockToDoListRepository.EXPECT().Save(mockToDoList).Return(nil, mockAppError)

	_, err := defaultToDoListService.SaveList(mockToDoList)

	if err == nil {
		t.Error("Error expected, nil returned")
		return
	}

	if !reflect.DeepEqual(*err, *mockAppError) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_GetOneListById_should_return_list_returned_by_repo_method
 * --------------------
 * Tests if pointer to domain.ToDoList from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_GetOneListById_should_return_list_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockToDoList := domain.ToDoList{
		Id:          primitive.ObjectID{},
		Name:        "mock list",
		Description: nil,
		Tasks: []domain.Task{
			{
				Id:          "test_id",
				Name:        "test task name",
				Description: nil,
			},
		},
	}

	mockToDoListRepository.EXPECT().GetOneById("test_id").Return(&mockToDoList, nil)

	list, err := defaultToDoListService.GetOneListById("test_id")

	if err != nil {
		t.Errorf("Nil expected, error returned: %v", err.Code)
	}

	if !reflect.DeepEqual(*list, mockToDoList) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_SaveList_should_return_error_returned_by_repo_method
 * --------------------
 * Tests if pointer to errs.AppError from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_GetOneListById_should_return_error_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockAppError := errs.NewInternalError("test error")
	mockToDoListRepository.EXPECT().GetOneById("test_id").Return(nil, mockAppError).Times(1)

	_, err := defaultToDoListService.GetOneListById("test_id")

	if err == nil {
		t.Error("Error expected, nil returned")
		return
	}

	if !reflect.DeepEqual(*err, *mockAppError) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_UpdateOneListById_should_return_list_returned_by_repo_method
 * --------------------
 * Tests if pointer to domain.ToDoList from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_UpdateOneListById_should_return_list_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockToDoList := domain.ToDoList{
		Id:          primitive.ObjectID{},
		Name:        "mock list",
		Description: nil,
		Tasks: []domain.Task{
			{
				Id:          "test_id",
				Name:        "test task name",
				Description: nil,
			},
		},
	}

	mockToDoListRepository.EXPECT().UpdateOneById("test_id", mockToDoList).Return(&mockToDoList, nil)

	list, err := defaultToDoListService.UpdateOneListById("test_id", mockToDoList)

	if err != nil {
		t.Errorf("Nil expected, error returned: %v", err.Code)
	}

	if !reflect.DeepEqual(*list, mockToDoList) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_UpdateOneListById_should_return_error_returned_by_repo_method
 * --------------------
 * Tests if pointer to errs.AppError from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_UpdateOneListById_should_return_error_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockToDoList := domain.ToDoList{
		Id:          primitive.ObjectID{},
		Name:        "mock list",
		Description: nil,
		Tasks: []domain.Task{
			{
				Id:          "test_id",
				Name:        "test task name",
				Description: nil,
			},
		},
	}
	mockAppError := errs.NewInternalError("test error")
	mockToDoListRepository.EXPECT().UpdateOneById("test_id", mockToDoList).Return(nil, mockAppError)

	_, err := defaultToDoListService.UpdateOneListById("test_id", mockToDoList)

	if err == nil {
		t.Error("Error expected, nil returned")
		return
	}

	if !reflect.DeepEqual(*err, *mockAppError) {
		t.Error("Data does not match mock return")
	}
}

/*
 * function: Test_DefaultToDoListService_DeleteList_should_return_nil_if_repo_method_returns_nil
 * --------------------
 * Tests if nil is returned by method if called repository method returns nil
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_DeleteList_should_return_nil_if_repo_method_returns_nil(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockToDoListRepository.EXPECT().DeleteOneById("test_id").Return(nil)
	err := defaultToDoListService.DeleteListById("test_id")

	if err != nil {
		t.Error("Error returned, nil expected")
	}
}

/*
 * function: Test_DefaultToDoListService_DeleteList_should_return_error_returned_by_repo_method
 * --------------------
 * Tests if pointer to errs.AppError from repository method is returned unmodified by service method.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_DefaultToDoListService_DeleteList_should_return_error_returned_by_repo_method(t *testing.T) {
	teardown := setupToDoListServiceTest(t)
	defer teardown()

	mockAppError := errs.NewInternalError("test error")
	mockToDoListRepository.EXPECT().DeleteOneById("test_id").Return(mockAppError)
	err := defaultToDoListService.DeleteListById("test_id")

	if err == nil {
		t.Error("Nil returned, error expected")
		return
	}

	if !reflect.DeepEqual(*err, *mockAppError) {
		t.Error("Data does not match mock return")
	}
}
