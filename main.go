package main

import (
	"log/slog"
	"os"
	"taskflow/infrastructure"
)

func main() {
	slog.Info("called to start server")

	config := parseConfig()

	infrastructure.Init(config)
}

func parseConfig() string {
	env := os.Getenv("env")
	if env == "" {
		panic("env variable is required")
	}

	slog.Info("environment", "env", env)

	if env == "test" || env == "testing" || env == "dev" || env == "development" {
		return "config/config.local.toml"
	}

	if env == "prod" || env == "production" {
		return "config/config.prod.toml"
	}

	panic("invalid env variable")
}
