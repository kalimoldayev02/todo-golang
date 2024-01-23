package app

import (
	"database/sql"
	"fmt"
	"log"
	"todo-golang/internal/handler"
	"todo-golang/internal/repository"
	"todo-golang/internal/router"
	"todo-golang/internal/service"
	"todo-golang/pkg/config"

	"github.com/gofiber/fiber/v2"
)

const (
	postgresDriver = "postgres"
)

func Run(cfg *config.Config) {
	// db
	db, err := intiDb(cfg)
	if err != nil {
		log.Printf(err.Error())
	}

	// repository
	repo := repository.NewRespository(db)

	// service
	service := service.NewService(repo)

	// HTTP Server
	app := router.NewRouter()
	initRoutes(app, service)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.HttpServer.Port)))
}

func initRoutes(app *fiber.App, service *service.Service) {
	h := handler.NewHandler(service)

	api := app
	api.Group("/api").Get("/users", h.GetUsers)
}

func intiDb(c *config.Config) (*sql.DB, error) {
	var dsn string
	cfg := c.DataBase

	switch cfg.Driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.User, cfg.TableName, cfg.Password)
	}

	return sql.Open(cfg.Driver, dsn)
}
