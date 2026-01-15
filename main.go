package main

import (
	"flag"
	"log/slog"
	"taskflow/infrastructure"
)

func main() {
	slog.Info("called to start server")

	config := parseConfig()

	infrastructure.Init(config)
}

func parseConfig() string {
	envFlag := flag.String("env", "", "environment")
	flag.Parse()

	if envFlag == nil {
		panic("env variable is required")
	}

	env := *envFlag

	slog.Info("environment", "env", env)

	if env == "test" || env == "testing" || env == "dev" || env == "development" {
		return "config/config.local.toml"
	}

	if env == "prod" || env == "production" {
		return "config/config.prod.toml"
	}

	panic("invalid env variable")
}
