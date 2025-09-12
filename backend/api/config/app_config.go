package appconfig

import envreader "github.com/ilkerciblak/buldum-app/shared/helper/env_reader"

type AppConfig struct {
	PORT string
}

func NewAppConfig() (*AppConfig, error) {
	port := envreader.GetStringOrDefault("PORT", "8000")

	return &AppConfig{
			PORT: port,
		},
		nil
}
