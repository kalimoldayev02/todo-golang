package main

import (
	"todo-golang/internal/app"
	"todo-golang/pkg/config"
)

func main() {
	// db := driverfactory.Make()
	// result := db.Add("SELECT *").Add("FROM users").Execute().FetchAll()

	config := config.GetInstance()
	app.Run(config)
}
