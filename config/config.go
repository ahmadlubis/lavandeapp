package config

import (
	"fmt"
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	MysqlConfig MysqlConfig
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

type MysqlConfig struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	DB       string `env:"MYSQL_DATABASE"`
	Username string `env:"MYSQL_USERNAME"`
	Password string `env:"MYSQL_PASSWORD"`
}

func (config *MysqlConfig) BuildDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DB,
	)
}
