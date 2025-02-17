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
	DBPort     int    `env:"DB_PORT"`
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	User       string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
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
