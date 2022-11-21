package config

import (
	"os"
	"strconv"
	"sync"
)

var config *Config
var doOnce sync.Once

type GetConfigFn func() *Config

func GetConfig() *Config {

	doOnce.Do(func() {
		config = &Config{
			Environment:  os.Getenv("ENV"),
			DbConnString: os.Getenv("DB_CONNSTRING"),
		}

		restPort := os.Getenv("REST_PORT")

		intRestPort, err := strconv.Atoi(restPort)
		if err != nil {
			panic(err)
		}

		config.RestPort = intRestPort
	})

	return config
}
