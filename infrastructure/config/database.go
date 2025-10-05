package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Driver   string `toml:"driver"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	Conn     *sql.DB
}

func (d Database) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.User, d.Password, d.Host, d.Port, d.Name)
}

func NewDatabase() *Database {
	return &Database{}
}

func Connect(cfg *Config) error {
	log.Printf("Connecting to database %s", cfg.Database.Name)
	log.Printf("Driver: %s", cfg.Database.Driver)

	conn, err := sql.Open(cfg.Database.Driver, cfg.Database.GetDSN())
	if err != nil {
		return fmt.Errorf("error opening connection: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %s", err)
	}

	cfg.Database.Conn = conn
	log.Println("Database connection established successfully")
	return nil
}
