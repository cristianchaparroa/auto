package config

import (
	"fmt"

	"github.com/cristianchaparroa/auto/connection"
)

// Config contains all configurations
type Config struct {
	PathModels string `yaml:"scan"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Database   string `yaml:"database"`
	Driver     string `yaml:"driver"`
	Schema     string `yaml:"schema"`
	Orcl       string `yaml:"orcl"`
}

func (c *Config) String() string {
	return fmt.Sprintf("Scan:%s, Host:%s, Port:%d, User:%s, Password:%s, Database:%s, Driver:%s, Schema:%s, Orcl:%s",
		c.PathModels, c.Host, c.Port, c.User, c.Password, c.Database, c.Driver, c.Schema, c.Orcl)
}

// NewDatabaseConfig retrives just the database configuration from Global
// configuration
func NewDatabaseConfig(c *Config) *connection.Config {
	dbConfig := &connection.Config{}
	dbConfig.Host = c.Host
	dbConfig.Port = c.Port
	dbConfig.User = c.User
	dbConfig.Password = c.Password
	dbConfig.Database = c.Database
	dbConfig.Driver = c.Driver
	dbConfig.Schema = c.Schema
	dbConfig.Orcl = c.Orcl

	return dbConfig
}
