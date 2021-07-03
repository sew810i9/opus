package configs

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port            string `envconfig:"PORT"`
	MysqlConnection string `envconfig:"MYSQL_CONNECTION"`
}

func Init() *Config {
	var s Config
	err := envconfig.Process("opus", &s)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &s
}
