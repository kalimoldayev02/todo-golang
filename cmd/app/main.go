package main

import (
	"todo-golang/app"
	"todo-golang/pkg/config"

	_ "github.com/lib/pq"
)

func main() {
	config := config.GetInstance()
	app.Run(config)
}
