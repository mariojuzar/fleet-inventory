package config

import (
	"github.com/jinzhu/configor"
	"time"
)

type (
	MainConfig struct {
		Server   ServerConfig
		Database DBConfig
	}

	ServerConfig struct {
		Port          uint          `env:"SERVER_PORT"`
		WriteTimeout  time.Duration `env:"SERVER_WRITE_TIMEOUT"`
		ReadTimeout   time.Duration `env:"SERVER_READ_TIMEOUT"`
		EnableSwagger bool          `env:"SERVER_ENABLE_SWAGGER" default:"false"`
	}

	DBConfig struct {
		DSN             string        `env:"DB_DSN"`
		RetryInterval   int           `env:"DB_RETRY_INTERVAL"`
		MaxIdleConn     int           `env:"DB_MAX_IDLE_CONN"`
		MaxConn         int           `env:"DB_MAX_CONN"`
		ConnMaxLifetime time.Duration `env:"DB_CONN_MAX_LIFETIME"`
	}
)

func ReadConfig() (*MainConfig, error) {
	config := &MainConfig{}
	err := configor.Load(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
