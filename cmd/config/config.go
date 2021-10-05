package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/config"
)

type HTTPConf struct {
	Addr string `envconfig:"HTTP_ADDR" default:"0.0.0.0:8080"`
}

type Config struct {
	HTTP  HTTPConf
	MySQL config.MySQLConfig

	SecretKey string `envconfig:"SECRET_KEY" default:"secret"`
	Origin    string `envconfig:"ORIGIN"`
}

func NewConfig() (*Config, error) {
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	if err != nil {
		return nil, errors.Wrap(err, "load config fail")
	}
	return &cnf, nil
}
