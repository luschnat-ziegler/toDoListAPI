/*
 * package: handlers
 * --------------------
 * Includes handler function definitions.
 */

package handlers

import (
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/toDoListAPI/testUtils/dummies"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
 * function: Test_GetInfo_should_write_provided_info_to_JSON_in_response_body_and_code_200_to_header
 * --------------------
 * Tests if method writes correct JSON to response body as well as status code 200
 *
 * t: a pointer to testing.T to meet test function signature requirements.
 *
 * Returns: nothing
 */

func Test_GetInfo_should_write_provided_info_to_JSON_in_response_body_and_code_200_to_header(t *testing.T) {
	router = mux.NewRouter()
	defer func(){router = nil}()

	router.HandleFunc("/", GetInfo)

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected code 200, got %v instead", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	if resBody != dummies.DummyExpectedInfoJSON {
		t.Error("Response body does not match")
	}
}