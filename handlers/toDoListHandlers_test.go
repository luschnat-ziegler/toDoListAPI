/*
 * package: handlers
 * --------------------
 * Includes handler function definitions.
 */

package handlers

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/testUtils/dummies"
	"github.com/luschnat-ziegler/toDoListAPI/testUtils/mocks/ports"
	"net/http"
	"net/http/httptest"
	"testing"
)

var th ToDoListHandlers
var mockDefaultToDoListService *ports.MockToDoListService
var router *mux.Router

/*
 * function: setupToDoListHandlersTest
 * --------------------
 * Sets up variables for tests and returns teardown function.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: a function to clean up test variables.
 */

func setupToDoListHandlersTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockDefaultToDoListService = ports.NewMockToDoListService(ctrl)
	th = ToDoListHandlers{mockDefaultToDoListService}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

/*
 * function: Test_DoListHandlers_GetAll_should_write_lists_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 200 if service method returns
 * pointer to slice of domain.ToDoList and nil.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_GetAll_should_write_lists_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos", th.GetAll)

	dummyLists := []domain.ToDoList{
		dummies.DummyListValid,
	}
	mockDefaultToDoListService.EXPECT().GetAllLists().Return(&dummyLists, nil)

	request, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected code 200, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	if resBody != "["+dummies.DummyListValidAsJSON+"]" {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_GetAll_should_write_error_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as corresponding status code if service method
 * returns nil and a pointer to an errs.AppError.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_GetAll_should_write_error_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos", th.GetAll)
	mockDefaultToDoListService.EXPECT().GetAllLists().Return(nil, dummies.DummyInternalError)

	request, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected code 500, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyInternalErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Save_should_write_list_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 201 if service method returns
 * pointer to domain.ToDoList and nil.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Save_should_write_list_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos", th.Save)
	mockDefaultToDoListService.EXPECT().SaveList(dummies.DummyListValid).Return(&dummies.DummyListValidWithIds, nil)

	request, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer([]byte(dummies.DummyValidSaveListRequestAsJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected code 201, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyListValidWithIdsAsJson {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Save_should_write_error_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as corresponding status code if service method
 * returns nil and a pointer to an errs.AppError.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Save_should_write_error_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos", th.Save)
	mockDefaultToDoListService.EXPECT().SaveList(dummies.DummyListValid).Return(nil, dummies.DummyInternalError)

	request, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer([]byte(dummies.DummyValidSaveListRequestAsJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected code 500, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyInternalErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Save_should_write_error_400_to_json_body_if_JSON_invalid
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 400 if JSON in request body
 * cannot be parsed.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Save_should_write_error_400_to_json_body_if_JSON_invalid(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos", th.Save)

	request, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer([]byte(dummies.DummyRequestInvalidJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected code 400, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyBadRequestErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Save_should_write_validation_error_to_json_body_if_validation_fails
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 400 if validation fails and returns
 * an errs.ValidationError
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Save_should_write_validation_error_to_json_body_if_validation_fails(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos", th.Save)

	request, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer([]byte(dummies.DummyInvalidSaveListRequestAsJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected code 400, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyValidationErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_GetOne_should_write_list_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 200 if service method returns
 * pointer to domain.ToDoList and nil.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_GetOne_should_write_list_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.GetOne)
	mockDefaultToDoListService.EXPECT().GetOneListById("test_id").Return(&dummies.DummyListValidWithIds, nil)

	request, _ := http.NewRequest(http.MethodGet, "/todos/test_id", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected code 200, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyListValidWithIdsAsJson {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_GetOne_should_write_error_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as the correct error code if service method returns
 * nil and a pointer to an errs.AppError
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_GetOne_should_write_error_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.GetOne)
	mockDefaultToDoListService.EXPECT().GetOneListById("test_id").Return(nil, dummies.DummyInternalError)

	request, _ := http.NewRequest(http.MethodGet, "/todos/test_id", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected code 500, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyInternalErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Update_should_write_list_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 200 if service method returns
 * pointer to domain.ToDoList and nil.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Update_should_write_list_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.Update)
	mockDefaultToDoListService.EXPECT().UpdateOneListById("test_id", dummies.DummyListValid).Return(&dummies.DummyListValidWithIds, nil)

	request, _ := http.NewRequest(http.MethodPut, "/todos/test_id", bytes.NewBuffer([]byte(dummies.DummyValidSaveListRequestAsJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected code 200, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyListValidWithIdsAsJson {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Update_should_write_error_returned_by_service_method_to_json_body
 * --------------------
 * Tests if method writes correct JSON to response body as well as the correct error code if service method returns
 * nil and a pointer to an errs.AppError
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Update_should_write_error_returned_by_service_method_to_json_body(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.Update)
	mockDefaultToDoListService.EXPECT().UpdateOneListById("test_id", dummies.DummyListValid).Return(nil, dummies.DummyInternalError)

	request, _ := http.NewRequest(http.MethodPut, "/todos/test_id", bytes.NewBuffer([]byte(dummies.DummyValidSaveListRequestAsJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected code 500, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyInternalErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Update_should_write_error_400_to_json_body_if_JSON_invalid
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 400 if JSON in request body
 * cannot be parsed.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Update_should_write_error_400_to_json_body_if_JSON_invalid(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.Update)

	request, _ := http.NewRequest(http.MethodPut, "/todos/test_id", bytes.NewBuffer([]byte(dummies.DummyRequestInvalidJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected code 400, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyBadRequestErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Update_should_write_validation_error_to_json_body_if_validation_fails
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 400 if validation fails and returns
 * an errs.ValidationError
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Update_should_write_validation_error_to_json_body_if_validation_fails(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.Update)

	request, _ := http.NewRequest(http.MethodPut, "/todos/test_id", bytes.NewBuffer([]byte(dummies.DummyInvalidSaveListRequestAsJSON)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected code 400, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyValidationErrorAsJSON {
		t.Error("Response body does not match")
	}
}

/*
 * function: Test_ToDoListHandlers_Delete_should_write_204_if_service_method_returns_no_error
 * --------------------
 * Tests if method writes code 204 to response header if service method returns nil.
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Delete_should_write_204_if_service_method_returns_no_error(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.Delete)
	mockDefaultToDoListService.EXPECT().DeleteList("test_id").Return(nil)

	request, _ := http.NewRequest(http.MethodDelete, "/todos/test_id", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected code 204, got %v instead", recorder.Code)
	}
}

/*
 * function: Test_ToDoListHandlers_Delete_should_write_error_returned_by_service
 * --------------------
 * Tests if method writes correct JSON to response body as well as the correct error code if service method returns
 * nil and a pointer to an errs.AppError
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_ToDoListHandlers_Delete_should_write_error_returned_by_service(t *testing.T) {
	teardown := setupToDoListHandlersTest(t)
	defer teardown()

	router.HandleFunc("/todos/{id}", th.Delete)
	mockDefaultToDoListService.EXPECT().DeleteList("test_id").Return(dummies.DummyInternalError)

	request, _ := http.NewRequest(http.MethodDelete, "/todos/test_id", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected code 500, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	if resBody != dummies.DummyInternalErrorAsJSON {
		t.Error("Response body does not match")
	}
}
