package handlers

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/toDoListAPI/mocks/ports"
	"testing"
)

var th ToDoListHandlers
var mockDefaultToDoListService *ports.MockToDoListService
var router *mux.Router

func setupToDoListHandlersTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockDefaultToDoListService := ports.NewMockToDoListService(ctrl)
	th = ToDoListHandlers{mockDefaultToDoListService}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}
