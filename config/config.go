package config

import (
	"github.com/subosito/gotenv"
	"os"
)

type Config struct {
	app *app
	db  *db
}

type app struct {
	port    string
	version string
}

type db struct {
	databaseURL string
}

// New config...
func NewConfig() *Config {
	_ = gotenv.Load(".env")
	return &Config{
		db: &db{
			databaseURL: os.Getenv("DATABASE_URL"),
		},
		app: &app{
			port:    os.Getenv("APP_PORT"),
			version: os.Getenv("APP_VERSION"),
		},
	}
}

// Postgres Address
func (c *Config) DatabaseURL() string {
	if c != nil {
		return c.db.databaseURL
	}

	return ""
}

// App port
func (c *Config) AppPort() string {
	if c != nil {
		return c.app.port
	}

	return ""
}

// App version
func (c *Config) AppVersion() string {
	if c != nil {
		return c.app.version
	}

	return ""
}
