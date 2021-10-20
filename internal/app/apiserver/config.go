package apiserver

type Config struct {
	BindAddr    string `toml:"bin_addr"`
	LogLevel    string `toml:"log_level"`
	DataBaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
