package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/BenMeredithConsult/locagri-apps/config"
	"github.com/BenMeredithConsult/locagri-apps/ent"
	"github.com/BenMeredithConsult/locagri-apps/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Adapter struct {
	DB *ent.Client
}

func NewDB(option ...*config.DBConf) *Adapter {
	conf := config.DB()
	if len(option) > 0 {
		conf = option[0]
	}
	switch conf.Driver {
	case "mysql":
		return &Adapter{DB: mysqlConnector(conf)}
	case "postgres":
		return &Adapter{DB: postgresConnector(conf)}
	}
	return nil
}

func MigrateDB(db *Adapter) error {
	ctx := context.Background()
	if err := db.DB.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}
	return nil
}

func Connect() *sql.DB {
	conf := config.DB()
	switch conf.Driver {
	case "postgres":
		psDSN := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Name,
			conf.Password,
			conf.SSLMode,
		)
		db, err := sql.Open(conf.Driver, psDSN)
		if err != nil {
			log.Panicln(err.Error())
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
			log.Panicln(err.Error())
		}
		return db
	}
}
func ConnectServer() *sql.DB {
	conf := config.DB()
	switch conf.Driver {
	case "postgres":
		psDSN := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s sslmode=%s",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Password,
			conf.SSLMode,
		)
		db, err := sql.Open(conf.Driver, psDSN)
		if err != nil {
			log.Panicln(err.Error())
		}
		return db
	default:
		mysqlDSN := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/?multiStatements=true",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
		)
		db, err := sql.Open(conf.Driver, mysqlDSN)
		if err != nil {
			log.Panicln(err.Error())
		}
		return db
	}
}

type Language struct {
	name string
	code string
}

type Nationality struct {
	name string
	code string
}

type Country struct {
	name string
	code string
}

// Seed the database with default languages
func Seeder(db *Adapter) error {
	client := db.DB
	ctx := context.Background()

	// Seed languages if they don't exist
	languageCount, err := client.Language.Query().Count(ctx)
	if err != nil {
		return fmt.Errorf("could not query languages: %v", err)
	}
	if languageCount == 0 {
		languages := []struct {
			name string
			code string
		}{
			{"English", "en"},
			{"French", "fr"},
		}

		languageBulk := make([]*ent.LanguageCreate, len(languages))
		for i, l := range languages {
			languageBulk[i] = client.Language.Create().SetCode(l.code).SetName(l.name)
		}
		if _, err := client.Language.CreateBulk(languageBulk...).Save(ctx); err != nil {
			return fmt.Errorf("could not seed languages: %v", err)
		}
	}

	// Seed countries if they don't exist
	countryCount, err := client.Country.Query().Count(ctx)
	if err != nil {
		return fmt.Errorf("could not query countries: %v", err)
	}
	if countryCount == 0 {
		countries := []struct {
			name string
			code string
		}{
			{"Ghana", "GH"},
			{"Cote d'Ivoire", "CI"},
		}

		countryBulk := make([]*ent.CountryCreate, len(countries))
		for i, c := range countries {
			countryBulk[i] = client.Country.Create().SetCode(c.code).SetName(c.name)
		}
		if _, err := client.Country.CreateBulk(countryBulk...).Save(ctx); err != nil {
			return fmt.Errorf("could not seed countries: %v", err)
		}
	}

	// Seed languages if they don't exist
	nationalityCount, err := client.Nationality.Query().Count(ctx)
	if err != nil {
		return fmt.Errorf("could not query languages: %v", err)
	}
	if nationalityCount == 0 {
		languages := []struct {
			name string
			code string
		}{
			{"Ghanaian", "GH"},
			{"Ivorian", "CI"},
		}

		nationalityBulk := make([]*ent.NationalityCreate, len(languages))
		for i, l := range languages {
			nationalityBulk[i] = client.Nationality.Create().SetCode(l.code).SetName(l.name)
		}
		if _, err := client.Nationality.CreateBulk(nationalityBulk...).Save(ctx); err != nil {
			return fmt.Errorf("could not seed nationalities: %v", err)
		}
	}

	return nil
}
