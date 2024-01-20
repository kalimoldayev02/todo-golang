package main

import (
	"fmt"
	"todo-golang/pkg/config"
)

func main() {
	config := config.GetInstance().HttpServer
	fmt.Println(config)
}
