package once

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {

		var err error
		config = &Config{
			Server: os.Getenv("SERVER_URL"),
		}
		config.Port, err = strconv.ParseInt(os.Getenv("PORT"), 10, 0)
		if err != nil {
			config.Port = 8080
		}
		log.Println("init config")
	})
	return config
}
