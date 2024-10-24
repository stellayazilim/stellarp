package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	postgresPassword string
	postgresUser     string
	postgresHost     string
	postgresPort     string
	postgresDatabase string
}

func (c Config) GetPostgresPassword() string {
	return c.postgresPassword
}

func (c Config) GetPostgresUser() string {
	return c.postgresUser
}

func (c Config) GetPostgresHost() string {
	return c.postgresHost
}

func (c Config) GetPostgresPort() string {
	return c.postgresPort
}

func (c Config) GetPostgresDatabase() string {
	return c.postgresDatabase
}

func loadEnv() error {
	dir, _ := os.Getwd()
	return godotenv.Load(dir + "/.env")
}

func newConfig() *Config {
	return &Config{
		postgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		postgresUser:     os.Getenv("POSTGRES_USER"),
		postgresHost:     os.Getenv("POSTGRES_HOST"),
		postgresPort:     os.Getenv("POSTGRES_PORT"),
		postgresDatabase: os.Getenv("POSTGRES_DATABASE"),
	}
}
