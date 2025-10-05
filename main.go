package main

import (
	"flag"
	"log"
	"taskflow/infrastructure"
)

func main() {
	log.Print("Called to start server")

	config := parseConfig()
	log.Printf("Config file: %s", config)

	infrastructure.Init(config)
}

func parseConfig() string {
	configFilePath := flag.String("config", "config.toml", "Path to config file")
	flag.Parse()
	return *configFilePath
}
