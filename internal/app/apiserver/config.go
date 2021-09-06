package apiserver

import "http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bin_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
