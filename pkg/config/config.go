package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	HttpServer `yaml:"http_server"`
	DataBase   `yaml:"database"`
	Salt       string `yaml:"salt"`
}

type HttpServer struct {
	Host        string        `yaml:"host" env:"HOST"`
	Port        string        `yaml:"port" env:"PORT"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT"`
}

type DataBase struct {
	Driver   string `yaml:"driver" env:"DB_HOST"`
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"db_name" env:"DB_NAME"`
}

func GetInstance() *Config {
	if instance == nil {
		once.Do(func() { // принимает в аргумент функцию, которая отработает один раз за вызов
			instance = loadConfig()
		})
	}

	return instance
}

func loadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH is not set in .env")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can not read config: %s", err.Error())
	}

	return &cfg
}
