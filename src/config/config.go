package config

import (
	"os"

	"example.com/sample-go-wire/src/logger"
	"github.com/joho/godotenv"
)

func init() {
	if val := os.Getenv("ENV"); val != "" {
		fileName := "./src/config/." + val + ".env"
		err := godotenv.Load(fileName)
		if err != nil {
			logger.GetLogger().Warnf("error while loading config %v: ", fileName, err)
		}
	}

	err := godotenv.Load("./src/config/.env")
	if err != nil {
		logger.GetLogger().Warnf("error while loading config .env: %v", err)
	}
}
