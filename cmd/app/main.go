package main

import (
	"todo-golang/internal/app"
	"todo-golang/pkg/config"

	_ "github.com/lib/pq"
)

func main() {
	config := config.GetInstance()
	app.Run(config)
}
