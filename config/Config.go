package config

import (
	"fmt"
	"io/ioutil"

	"github.com/cristianchaparroa/auto/connection"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
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

// NewConfig returns the config.yaml
func NewConfig(path string) (*Config, error) {
	msj := fmt.Sprintf("Loading config from:%s \n", path)
	log.Info(msj)
	c := &Config{}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
