package config

import (
	"os"

	"github.com/BenMeredithConsult/locagri.worker.api/utils/env"
)

type (
	app struct {
		Name     string
		Version  string
		AppURL   string
		RabbitMQ string
	}
	sms struct {
		Sender  string
		Gateway string
	}

	mailer struct {
		Mailer      string
		Host        string
		Port        string
		Username    string
		Password    string
		Encryption  string
		FromAddress string
		FromName    string
	}

	arkesel struct {
		APIKey string
		URL    string
	}
)

func App() *app {
	if os.Getenv("APP_ENV") == "production" {
		return &app{
			Name:     os.Getenv("APP_NAME"),
			Version:  os.Getenv("APP_VERSION"),
			AppURL:   os.Getenv("APP_URL"),
			RabbitMQ: os.Getenv("RabbitMQ_HOST"),
		}
	}
	return &app{
		Name:     env.Get("APP_NAME", "Locagri Worker"),
		Version:  env.Get("APP_VERSION", "0.0.1"),
		AppURL:   env.Get("APP_URL", "http://127.0.0.1:80"),
		RabbitMQ: env.Get("RabbitMQ_HOST", "rabbitmq"),
	}
}

func Mailer() *mailer {
	if os.Getenv("APP_ENV") == "production" {
		return &mailer{
			Mailer:      os.Getenv("MAIL_MAILER"),
			Host:        os.Getenv("MAIL_HOST"),
			Port:        os.Getenv("MAIL_PORT"),
			Username:    os.Getenv("MAIL_USERNAME"),
			Password:    os.Getenv("MAIL_PASSWORD"),
			Encryption:  os.Getenv("MAIL_ENCRYPTION"),
			FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
			FromName:    os.Getenv("MAIL_FROM_NAME"),
		}
	}
	return &mailer{
		Mailer:      env.Get("MAIL_MAILER", "smtp"),
		Host:        env.Get("MAIL_HOST", ""),
		Port:        env.Get("MAIL_PORT", ""),
		Username:    env.Get("MAIL_USERNAME", ""),
		Password:    env.Get("MAIL_PASSWORD", ""),
		Encryption:  env.Get("MAIL_ENCRYPTION", ""),
		FromAddress: env.Get("MAIL_FROM_ADDRESS", "info@bookihub.com"),
		FromName:    env.Get("MAIL_FROM_NAME", "LocAgri"),
	}
}

func SMS() *sms {
	if os.Getenv("APP_ENV") == "production" {
		return &sms{
			Sender:  os.Getenv("SMS_SENDER"),
			Gateway: os.Getenv("SMS_GATEWAY"),
		}
	}
	return &sms{
		Sender:  env.Get("SMS_SENDER", "Asinyo"),
		Gateway: env.Get("SMS_GATEWAY", "arkesel"),
	}
}
func Arkesel() *arkesel {
	if os.Getenv("APP_ENV") == "production" {
		return &arkesel{
			APIKey: os.Getenv("ARKESEL_API_KEY"),
			URL:    os.Getenv("ARKESEL_URL"),
		}
	}
	return &arkesel{
		APIKey: env.Get("ARKESEL_API_KEY", ""),
		URL:    env.Get("ARKESEL_URL", ""),
	}
}
