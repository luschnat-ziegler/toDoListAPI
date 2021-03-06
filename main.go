package main

import (
	"github.com/joho/godotenv"
	"github.com/luschnat-ziegler/toDoListAPI/logger"
	"github.com/luschnat-ziegler/toDoListAPI/server"
)

/*
 * function: init
 * --------------------
 * Loads .env file
 *
 * returns: nothing
 */

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Error("No .env file found")
	}
}

func main() {
	server.Start()
}
