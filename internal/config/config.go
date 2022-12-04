package config

import (
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	ListenAddr   string        `toml:"listenaddr"`
	LogLevel     string        `toml:"log_level"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

func NewConfig() *Config {
	return &Config{
		ListenAddr:   "127.0.0.1:8080",
		LogLevel:     "info",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (c *Config) ReadConfig() error {
	_, err := toml.DecodeFile("config.toml", c)
	return err
}
