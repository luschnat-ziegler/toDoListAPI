package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/toDoListAPI/core/services"
	"github.com/luschnat-ziegler/toDoListAPI/handlers"
	"github.com/luschnat-ziegler/toDoListAPI/logger"
	"github.com/luschnat-ziegler/toDoListAPI/repositories"
	"log"
	"net/http"
	"os"
)

func Start() {
	sanityCheck()
	logger.Info("Application started...")

	toDoListRepositoryDB := repositories.NewToDoListRepositoryDB()
	th := handlers.ToDoListHandlers{Service: services.NewToDoListService(toDoListRepositoryDB)}

	router := mux.NewRouter()
	router.HandleFunc("/todos", th.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/todos", th.Save).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", th.GetOne).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", th.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/todos/{id}", th.Update).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func sanityCheck() {
	envVars := []string{
		"DB_URL",
	}
	for _, envVar := range envVars {
		_, ok := os.LookupEnv(envVar)
		if !ok {
			log.Fatal(fmt.Sprintf("Environment variable %s not set in .env. Terminating application...", envVar))
		}
	}
}
