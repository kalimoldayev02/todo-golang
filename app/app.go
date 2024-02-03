package app

import (
	"database/sql"
	"fmt"
	"log"
	"todo-golang/app/controllers"
	"todo-golang/app/repositories"
	"todo-golang/app/services"
	"todo-golang/pkg/config"
	"todo-golang/pkg/routes"
)

func Run(cfg *config.Config) {
	// DB
	db, err := initDb(cfg)
	if err != nil {
		log.Printf(err.Error())
	}

	// Repository
	repo := repositories.NewRepository(db)

	// Service
	service := services.NewService(repo)

	// Controller
	controller := controllers.NewController(service)

	app := routes.NewRouter()
	routes.PublicRoutes(app, controller)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.HttpServer.Port)))
}

func initDb(c *config.Config) (*sql.DB, error) {
	var dsn string
	cfg := c.DataBase

	switch cfg.Driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.User, cfg.DBName, cfg.Password)
	}

	return sql.Open(cfg.Driver, dsn)
}
