package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Listen   string `envconfig:"LISTEN,required"`
	MysqlDsn string `envconfig:"MYSQL_DSN,required"`
}

func Init() *Config {
	var s Config
	err := envconfig.Process("opus", &s)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &s
}
