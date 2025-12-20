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
	configFilePath := flag.String("config", "config.toml", "Path to config file")
	flag.Parse()
	return *configFilePath
}
