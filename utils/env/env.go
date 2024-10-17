package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func Get(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func Setup() {
	if os.Getenv("APP_ENV") != "production" {
		envFile := ".env"
		var err error
		Env, err = godotenv.Read(envFile)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
