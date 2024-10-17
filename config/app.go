package config

import (
	"os"
	"sample/utils/env"
)

type (
	app struct {
		Name             string
		Version          string
		AppURL           string
		Key              string
		TokenName        string
		PORT             string
		FilesystemDriver string
		Redis            string
		RabbitMQ         string
	}
	serverConfig struct {
		Prefork           bool
		CaseSensitive     bool
		StrictRouting     bool
		StreamRequestBody bool
		EnablePrintRoutes bool
		Concurrency       int64
		ServerHeader      string
		AppName           string
	}
)

func App() *app {
	if os.Getenv("APP_ENV") == "production" {
		return &app{
			Name:             os.Getenv("APP_NAME"),
			Version:          os.Getenv("APP_VERSION"),
			AppURL:           os.Getenv("APP_URL"),
			Key:              os.Getenv("APP_KEY"),
			TokenName:        os.Getenv("API_TOKEN_NAME"),
			PORT:             os.Getenv("SERVER_PORT"),
			FilesystemDriver: os.Getenv("FILESYSTEM_DRIVER"),
			Redis:            os.Getenv("REDIS_HOST"),
			RabbitMQ:         os.Getenv("RabbitMQ_HOST"),
		}
	}
	return &app{
		Name:             env.Get("APP_NAME", "My First API"),
		Version:          env.Get("APP_VERSION", "0.0.1"),
		AppURL:           env.Get("APP_URL", "http://127.0.0.1:8500"),
		Key:              env.Get("APP_KEY", "secretKEY5465"),
		TokenName:        env.Get("API_TOKEN_NAME", "remember"),
		PORT:             env.Get("PORT", "8500"),
		FilesystemDriver: env.Get("FILESYSTEM_DRIVER", "local"),
		Redis:            env.Get("REDIS_HOST", "redis"),
		RabbitMQ:         env.Get("RabbitMQ_HOST", "rabbitmq"),
	}
}
func Server() *serverConfig {
	return &serverConfig{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		StreamRequestBody: true,
		EnablePrintRoutes: true,
		Concurrency:       256 * 2048,
		ServerHeader:      "FirstApi",
		AppName:           env.Get("APP_NAME", "FirstApi"),
	}
}
