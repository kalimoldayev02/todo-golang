package main

import (
	"todo-golang/internal/app"
	"todo-golang/pkg/config"
)

func main() {
	config := config.GetInstance()
	app.Run(config)
}
