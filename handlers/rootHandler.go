/*
 * package: handlers
 * --------------------
 * Includes handler function definitions.
 */

package handlers

import "net/http"

func GetInfo(w http.ResponseWriter, _ *http.Request) {

	apiInfo := map[string]string{
		"1. GET /todos":         "Returns an array of all todo lists",
		"2. POST /todos":        "Creates and saves new todo, returns the newly created resource",
		"3. GET /todos/{id}":    "Returns the todo list with the provided id, if existing",
		"4. PUT /todos/{id}":    "Overwrites the todo list with the provided id (if existing) with the provided new list.",
		"5. DELETE /todos/{id}": "Deletes the todo list with the provided id, if existing",
	}

	writeResponse(w, http.StatusOK, apiInfo)
}
