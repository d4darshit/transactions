package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	User      string `toml:"user"`
	Password  string `toml:"password"`
	Host      string `toml:"host"`
	Port      string `toml:"port"`
	Name      string `toml:"name"`
	Charset   string `toml:"charset"`
	ParseTime bool   `toml:"parseTime"`
	Loc       string `toml:"loc"`
}

type ServerConfig struct {
	Port string `toml:"port"`
}

type Config struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
}

var AppConfig Config

// LoadConfig loads configuration from the config.toml file.
func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := Config{}

	config.Database.User = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Name = os.Getenv("DB_NAME")
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Server.Port = os.Getenv("APP_PORT")

	AppConfig = config
}
