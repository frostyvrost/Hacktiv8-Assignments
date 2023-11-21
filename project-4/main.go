package main

import (
	"project-4/app"
	"project-4/database"
)

func main() {
	database.StartDB()
	app.StartServer()
}
