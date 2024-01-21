package app

import (
	"fmt"
	"log"
	"todo-golang/internal/handler"
	"todo-golang/internal/repository"
	"todo-golang/internal/router"
	"todo-golang/pkg/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	repository.NewRespository(db)

	// service

	// HTTP Server
	app := router.NewRouter()
	initRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.HttpServer.Port)))
}

func initRoutes(app *fiber.App) {
	h := handler.New()
	h.InitRoutes(app)
}

func intiDb(c *config.Config) (*gorm.DB, error) {
	var dsn string
	cfg := c.DataBase

	switch cfg.Driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.User, cfg.TableName, cfg.Password)
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	return nil, fmt.Errorf("can not init DB")
}
