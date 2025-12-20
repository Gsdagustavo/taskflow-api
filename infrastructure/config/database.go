package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

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
	slog.Info("connecting to database")

	conn, err := sql.Open(cfg.Database.Driver, cfg.Database.GetDSN())
	if err != nil {
		return errors.Join(errors.New("failed to connect to database"), err)
	}

	err = conn.Ping()
	if err != nil {
		return errors.Join(errors.New("failed to ping database"), err)
	}

	cfg.Database.Conn = conn
	slog.Info("connected to database successfully")
	return nil
}
