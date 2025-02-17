package cfg

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	DB  Database
	JWT JWT
}

type Database struct {
	DBPort     int    `env:"DB_PORT"  env-required:"true"`
	DBHost     string `env:"DB_HOST"  env-required:"true"`
	DBName     string `env:"DB_NAME"  env-required:"true"`
	User       string `env:"DB_USER"  env-required:"true"`
	DBPassword string `env:"DB_PASSWORD" env-required:"true"`
}

type JWT struct {
	SecretKey string `env:"JWT_SECRET"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}
	return &cfg, nil
}
