package driverfactory

import (
	"todo-golang/pkg/config"
	"todo-golang/pkg/system/database"
	"todo-golang/pkg/system/database/postgres"
)

var instance = make(map[string]database.ActiveRecord)

// Генерация Driver'a
func Make() database.ActiveRecord {
	dbConfig := config.GetInstance().DataBase

	if _, ok := instance[dbConfig.Driver]; !ok {
		var driver database.ActiveRecord

		driverData := database.DriverData{
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			DBName:   dbConfig.DBName,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			Driver:   dbConfig.Driver,
		}

		switch dbConfig.Driver {
		case "postgres":
			driver = postgres.NewPostgres(driverData)
		}

		instance[dbConfig.Driver] = driver
	}

	return instance[dbConfig.Driver]
}
