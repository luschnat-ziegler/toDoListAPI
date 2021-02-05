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

func (ah *ToDoListHandlers) GetAll(w http.ResponseWriter, _ *http.Request) {
	lists, err := ah.Service.GetAllLists()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, lists)
}

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

func (ah *ToDoListHandlers) GetOne(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	getListResponse, appErr := ah.Service.GetOneListById(id)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, getListResponse)
}

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

func (ah *ToDoListHandlers) Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	appErr := ah.Service.DeleteListById(id)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
