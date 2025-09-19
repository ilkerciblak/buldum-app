package appconfig

import (
	envreader "github.com/ilkerciblak/buldum-app/shared/helper/env_reader"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	PORT      string
	DB_URL    string
	DB_DRIVER string
}

func NewAppConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		return &AppConfig{}, err
	}

	port := envreader.GetStringOrDefault("APP_PORT", "8000")
	db_url := envreader.GetStringOrDefault("CONN_STR", "")
	driver := envreader.GetStringOrDefault("DB_DRIVER", "postgres")

	return &AppConfig{
			PORT:      port,
			DB_URL:    db_url,
			DB_DRIVER: driver,
		},
		nil
}
