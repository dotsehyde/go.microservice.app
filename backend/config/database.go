package config

import (
	"os"

	"github.com/BenMeredithConsult/locagri-apps/utils/env"
)

type (
	DBConf struct {
		Driver        string
		Host          string
		Port          string
		Name          string
		Username      string
		Password      string
		SSLMode       string
		Mongo         string
		MongoPassword string
	}
	TenantConf struct {
		Subdomain string
		Name      string
		DB        *DBConf
	}
)

func DB() *DBConf {
	var driver string
	if os.Getenv("APP_ENV") == "production" {
		driver = os.Getenv("DB_DRIVER")
	} else {
		driver = env.Get("DB_DRIVER", "sqlite")
	}
	switch driver {
	case "mysql":
		if os.Getenv("APP_ENV") == "production" {
			return &DBConf{
				Driver:        driver,
				Host:          os.Getenv("DB_HOST"),
				Port:          os.Getenv("DB_PORT"),
				Name:          os.Getenv("DB_DATABASE"),
				Username:      os.Getenv("DB_USERNAME"),
				Password:      os.Getenv("DB_PASSWORD"),
				Mongo:         os.Getenv("MONGO_HOST"),
				MongoPassword: os.Getenv("MONGO_PASSWORD"),
			}
		}
		return &DBConf{
			Driver:        driver,
			Host:          env.Get("DB_HOST", "127.0.0.1"),
			Port:          env.Get("DB_PORT", "3306"),
			Name:          env.Get("DB_DATABASE", "test_db"),
			Username:      env.Get("DB_USERNAME", "user"),
			Password:      env.Get("DB_PASSWORD", ""),
			Mongo:         env.Get("MONGO_HOST", "mongo"),
			MongoPassword: env.Get("MONGO_PASSWORD", "gesDB123"),
		}
	case "postgres":
		if os.Getenv("APP_ENV") == "production" {
			return &DBConf{
				Driver:   driver,
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Name:     os.Getenv("DB_DATABASE"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				SSLMode:  os.Getenv("DB_SSLMODE"),
				Mongo:    os.Getenv("MONGO_HOST"),
			}

		}
		return &DBConf{
			Driver:   driver,
			Host:     env.Get("DB_HOST", "127.0.0.1"),
			Port:     env.Get("DB_PORT", "5432"),
			Name:     env.Get("DB_DATABASE", "test_db"),
			Username: env.Get("DB_USERNAME", "postgres"),
			Password: env.Get("DB_PASSWORD", ""),
			SSLMode:  env.Get("DB_SSLMODE", "disable"),
			Mongo:    env.Get("MONGO_HOST", "mongo"),
		}
	}
	return nil
}
