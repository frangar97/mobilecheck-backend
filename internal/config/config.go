package config

import (
	"os"
)

type config struct {
	Port        string
	DatabaseUrl string
}

func InitConfig() (*config, error) {
	/*err := godotenv.Load()

	if err != nil {
		return nil, err
	}*/

	cfg := config{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}

	return &cfg, nil
}
