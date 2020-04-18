package config

import (
	"context"
	"github.com/subosito/gotenv"
	"github.com/zhs/loggr"
	"os"
)

type Config struct {
	app *app
	db  *db
	ctx context.Context
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
	logger := loggr.New("@version", os.Getenv("APP_VERSION"))
	ctx := loggr.ToContext(context.TODO(), logger)

	return &Config{
		db: &db{
			databaseURL: os.Getenv("DATABASE_URL"),
		},
		app: &app{
			port:    os.Getenv("APP_PORT"),
			version: os.Getenv("APP_VERSION"),
		},
		ctx: ctx,
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

// Context ...
func (c *Config) Context() context.Context {
	if c != nil {
		return c.ctx
	}

	return context.Background()
}
