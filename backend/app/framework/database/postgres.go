package database

import (
	"fmt"
	"log"

	"github.com/BenMeredithConsult/locagri-apps/config"
	"github.com/BenMeredithConsult/locagri-apps/ent"
)

func postgresConnector(conf *config.DBConf) *ent.Client {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Name,
		conf.Password,
		conf.SSLMode,
	)
	client, err := ent.Open(conf.Driver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	return client
}
