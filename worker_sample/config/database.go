package config

import (
	"os"

	"github.com/BenMeredithConsult/locagri.worker.api/utils/env"
)

type db struct {
	Driver   string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	KVStore  string
}

func DB() *db {
	var driver string
	if os.Getenv("APP_ENV") == "production" {
		driver = os.Getenv("DB_DRIVER")
	} else {
		driver = env.Get("DB_DRIVER", "sqlite")
	}
	switch driver {
	case "mysql":
		if os.Getenv("APP_ENV") == "production" {
			return &db{
				Driver:   driver,
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Name:     os.Getenv("DB_DATABASE"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				KVStore:  os.Getenv("KV_DB_STORE"),
			}
		}
		return &db{
			Driver:   driver,
			Host:     env.Get("DB_HOST", "127.0.0.1"),
			Port:     env.Get("DB_PORT", "3306"),
			Name:     env.Get("DB_DATABASE", "test_db"),
			Username: env.Get("DB_USERNAME", "root"),
			Password: env.Get("DB_PASSWORD", ""),
			KVStore:  env.Get("KV_DB_STORE", "mnt/cache/store"),
		}
	case "postgres":
		if os.Getenv("APP_ENV") == "production" {
			return &db{
				Driver:   driver,
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Name:     os.Getenv("DB_DATABASE"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				KVStore:  os.Getenv("KV_DB_STORE"),
			}

		}
		return &db{
			Driver:   driver,
			Host:     env.Get("DB_HOST", "127.0.0.1"),
			Port:     env.Get("DB_PORT", "5432"),
			Name:     env.Get("DB_DATABASE", "test_db"),
			Username: env.Get("DB_USERNAME", "postgres"),
			Password: env.Get("DB_PASSWORD", ""),
			KVStore:  env.Get("KV_DB_STORE", "mnt/cache/store"),
		}
	}
	return nil
}
