package config

import "time"

const (
	defaultServerPort               = "8080"
	defaultServerRWTimeout          = 10 * time.Second
	defaultServerMaxHeaderMegabytes = 1
)

type ServerConfig struct {
	Port               string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

type Config struct {
	Server ServerConfig
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:               defaultServerPort,
			ReadTimeout:        defaultServerRWTimeout,
			WriteTimeout:       defaultServerRWTimeout,
			MaxHeaderMegabytes: defaultServerMaxHeaderMegabytes,
		},
	}
}
