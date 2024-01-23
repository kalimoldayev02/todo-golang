package database

import (
	"todo-golang/pkg/config"
)

var instance map[string]DriverInterface

type DriverData struct {
	Host     string
	Port     string
	DBName   string
	Password string
}

type Driver struct {
	Driver DriverInterface
}

func Make() DriverInterface {
	dbConfig := config.GetInstance().DataBase

	if _, ok := instance[dbConfig.Driver]; ok == false {
		var dbInter DriverInterface
		driverData := DriverData{
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			DBName:   dbConfig.DBName,
			Password: dbConfig.Password,
		}

		switch dbConfig.Driver {
		case "postgres":
			dbInter = NewPostgres(driverData)
		}

		instance[dbConfig.Driver] = dbInter
	}

	return instance[dbConfig.Driver]
}
