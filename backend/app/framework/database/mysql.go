package database

import (
	"fmt"
	"log"

	"github.com/BenMeredithConsult/locagri-apps/config"
	"github.com/BenMeredithConsult/locagri-apps/ent"
)

func mysqlConnector(conf *config.DBConf) *ent.Client {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)
	client, err := ent.Open(conf.Driver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	return client
}
