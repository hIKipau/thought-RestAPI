package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"prod"`
	DatabaseURL string `yaml:"database_url" env:"DATABASE_URL"`
	HTTPServer  `yaml:"http_server" env:"HTTP_SERVER"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
}

func Load() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return &Config{}, errors.New("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, errors.New("CONFIG_PATH does not exist path: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return &Config{}, errors.New("error loading config: " + err.Error())
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return &Config{}, errors.New("error loading config: " + err.Error())
	}

	return &cfg, nil
}
