/*
 * package: server
 * --------------------
 * Includes server functionalities and application wiring
 */

package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/toDoListAPI/core/services"
	"github.com/luschnat-ziegler/toDoListAPI/handlers"
	"github.com/luschnat-ziegler/toDoListAPI/logger"
	"github.com/luschnat-ziegler/toDoListAPI/repositories"
	"net/http"
	"os"
)

/*
 * function: Start
 * --------------------
 * Sets up routing as well as repositories, services and handlers with
 * their dependencies. Starts the server listening for requests.
 *
 * returns: nothing
 */

func Start() {
	if sanityCheck() {
		logger.Info("Application started...")

		toDoListRepositoryDB := repositories.NewToDoListRepositoryDB()
		th := handlers.ToDoListHandlers{Service: services.NewToDoListService(toDoListRepositoryDB)}

		router := mux.NewRouter()
		router.HandleFunc("/", handlers.GetInfo).Methods(http.MethodGet)
		router.HandleFunc("/todos", th.GetAll).Methods(http.MethodGet)
		router.HandleFunc("/todos", th.Save).Methods(http.MethodPost)
		router.HandleFunc("/todos/{id}", th.GetOne).Methods(http.MethodGet)
		router.HandleFunc("/todos/{id}", th.Update).Methods(http.MethodPut)
		router.HandleFunc("/todos/{id}", th.Delete).Methods(http.MethodDelete)

		if err := http.ListenAndServe(":8000", router); err != nil {
			logger.Error("Error starting server: " + err.Error())
		}
	}
}

/*
 * function: sanityCheck
 * --------------------
 * Checks for existence of needed environment variables. In case of missing variable logs error indicating missing
 * variable
 *
 * returns: bool; true if check passes, false otherwise.
 */

func sanityCheck() bool {
	envVars := []string{
		"DB_URL",
	}
	for _, envVar := range envVars {
		_, ok := os.LookupEnv(envVar)
		if !ok {
			logger.Error(fmt.Sprintf("Environment variable %s not set in .env. Terminating application...", envVar))
			return false
		}
	}
	return true
}
