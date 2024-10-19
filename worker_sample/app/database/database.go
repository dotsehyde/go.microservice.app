package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/BenMeredithConsult/locagri.worker.api/config"
	"github.com/BenMeredithConsult/locagri.worker.api/ent"
	"github.com/BenMeredithConsult/locagri.worker.api/utils/env"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Adapter struct {
	DB *ent.Client
}

func NewDB() *Adapter {
	var dBDriver string
	if os.Getenv("APP_ENV") == "production" {
		dBDriver = os.Getenv("DB_DRIVER")
	} else {
		dBDriver = env.Get("DB_DRIVER", "sqlite")
	}
	switch dBDriver {
	case "mysql":
		return &Adapter{DB: mysqlConnector(dBDriver)}
	case "postgres":
		return &Adapter{DB: postgresConnector(dBDriver)}
	}
	return nil
}

func Connect() *sql.DB {
	conf := config.DB()
	var DbSSLMode string
	if os.Getenv("APP_ENV") == "production" {
		DbSSLMode = os.Getenv("DB_SSLMODE")
	} else {
		DbSSLMode = env.Get("DB_SSLMODE", "disable")
	}
	switch conf.Driver {
	case "postgres":
		psDSN := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Name,
			conf.Password,
			DbSSLMode,
		)
		db, err := sql.Open(conf.Driver, psDSN)
		if err != nil {
			log.Fatalf(err.Error())
		}
		return db
	default:
		mysqlDSN := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name,
		)
		db, err := sql.Open(conf.Driver, mysqlDSN)
		if err != nil {
			log.Fatalf(err.Error())
		}
		return db
	}

}
