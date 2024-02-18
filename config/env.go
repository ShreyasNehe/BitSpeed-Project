package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvironmentData struct {
	DBUrl   string `validate:"required"`
	EnvPort string `validate:"required"`
}

var EnvVariables EnvironmentData

func InitializeEnv() {
	_ = godotenv.Load("./Config.env")
	EnvVariables = EnvironmentData{
		DBUrl:   os.Getenv("PSQL_DB_URL"),
		EnvPort: os.Getenv("PORT"),
	}

	v := validator.New()
	err := v.Struct(EnvVariables)
	if err != nil {
		log.Panic(err)
	}
}
