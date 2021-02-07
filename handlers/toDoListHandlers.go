/*
 * package: handlers
 * --------------------
 * Includes handler function definitions.
 */

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/core/ports"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
	"net/http"
)

type ToDoListHandlers struct {
	Service ports.ToDoListService
}

/*
 * Method: ToDoListHandlers.GetAll
 * --------------------
 * To be called when all lists are requested. Writes them to the response body as JSON and code 200 to the header.
 * If a pointer to an errs.AppError is returned by the service method, its message is written to the response body
 * and its Code to the header, instead.
 *
 * w, _: an http.ResponseWriter and a pointer to an http.Request needed to meet the handler function signature.
 *
 * returns: nothing
 */

func (ah *ToDoListHandlers) GetAll(w http.ResponseWriter, _ *http.Request) {
	lists, err := ah.Service.GetAllLists()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, lists)
}

/*
 * Method: ToDoListHandlers.Save
 * --------------------
 * To be called when a posted list is to be saved. Rejects invalid JSON bodies and lists failing validation and writes
 * the respective information as JSON to the response body as well as the error code to the header.
 * If a pointer to an errs.AppError is returned by the service method, its message is
 * written to the response body and its Code to the header.
 * On success, the newly created resource is written to the response body as JSON and code 201 to the header.
 *
 * w, r: an http.ResponseWriter and a pointer to an http.Request needed to meet the handler function signature.
 *
 * returns: nothing
 */

func (ah *ToDoListHandlers) Save(w http.ResponseWriter, r *http.Request) {

	var newList domain.ToDoList

	err := json.NewDecoder(r.Body).Decode(&newList)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Body parsing error").AsMessage())
		return
	}

	validationError := newList.Validate()
	if validationError != nil {
		writeResponse(w, validationError.Code, validationError.AsMessage())
		return
	}

	getListResponse, appErr := ah.Service.SaveList(newList)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, getListResponse)
}

/*
 * Method: ToDoListHandlers.GetOne
 * --------------------
 * To be called when one specific list is requested. Writes it to the response body as JSON and code 200 to the header.
 * If a pointer to an errs.AppError is returned by the service method, its message is written to the response body
 * and its Code to the header, instead.
 *
 * w, r: an http.ResponseWriter and a pointer to an http.Request needed to meet the handler function signature.
 *
 * returns: nothing
 */

func (ah *ToDoListHandlers) GetOne(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	getListResponse, appErr := ah.Service.GetOneListById(id)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, getListResponse)
}

/*
 * Method: ToDoListHandlers.Update
 * --------------------
 * To be called when one specific list is requested to be overwritten. Writes it to the response body as JSON and
 * code 200 to the header. If a pointer to an errs.AppError is returned by the service method, its message is
 * written to the response body and its Code to the header, instead.
 *
 * w, r: an http.ResponseWriter and a pointer to an http.Request needed to meet the handler function signature.
 *
 * returns: nothing
 */

func (ah *ToDoListHandlers) Update(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	var newList domain.ToDoList
	err := json.NewDecoder(r.Body).Decode(&newList)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Body parsing error").AsMessage())
		return
	}

	validationError := newList.Validate()
	if validationError != nil {
		writeResponse(w, validationError.Code, validationError.AsMessage())
		return
	}

	updatedList, appErr := ah.Service.UpdateOneListById(id, newList)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, updatedList)
}

/*
 * Method: ToDoListHandlers.Delete
 * --------------------
 * To be called when one specific list is requested to be deleted. Writes no response body and code 204 to the header.
 * If a pointer to an errs.AppError is returned by the service method, its message is written to the response body
 * and its Code to the header, instead.
 *
 * w, r: an http.ResponseWriter and a pointer to an http.Request needed to meet the handler function signature.
 *
 * returns: nothing
 */

func (ah *ToDoListHandlers) Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	appErr := ah.Service.DeleteListById(id)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

/*
 * Function: writeResponse
 * --------------------
 * Utility function for writing http responses with status code and JSON body.
 *
 * w: an http.ResponseWriter to be used for writing the response
 * code: an integer representing a status code
 * data: The data to be written, generic type
 *
 * returns: nothing
 */

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
