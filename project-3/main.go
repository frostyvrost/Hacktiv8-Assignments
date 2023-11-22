package main

import (
	"project-3/app"
	"project-3/database"
)

func main() {
	database.StartDB()
	app.StartServer()
}
