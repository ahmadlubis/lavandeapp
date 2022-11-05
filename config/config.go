package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	MysqlConfig MysqlConfig
	AuthConfig  AuthConfig
}

func NewConfig() *Config {
	_ = gotenv.Load() // load .env if needed

	var config Config
	err := envdecode.Decode(&config)
	if err != nil {
		panic(err)
	}

	return &config
}

type AuthConfig struct {
	JWTSecretKey string `env:"JWT_SECRET_KEY"`
}
