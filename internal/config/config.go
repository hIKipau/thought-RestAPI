package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"prod"`
	DatabaseURL string `yaml:"database_url" env:"DATABASE_URL"`
	HTTPServer  `yaml:"http_server" env:"HTTP_SERVER"`
	Version     string `yaml:"version" env:"VERSION"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
}

func Load() (*Config, error) {
	const op = "internal/config/Load"

	err := godotenv.Load()
	if err != nil {
		return &Config{}, fmt.Errorf("%s: Cant loading .env file, Error: %s", op, err.Error())
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return &Config{}, fmt.Errorf("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, fmt.Errorf("%s: CONFIG_PATH does not exist path: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return &Config{}, fmt.Errorf("%s: Cant read config, Error: %s", op, err.Error())
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return &Config{}, fmt.Errorf("%s: Cant read env, Error: %s", op, err.Error())
	}

	return &cfg, nil
}
