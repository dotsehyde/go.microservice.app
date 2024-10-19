package appconst

import (
	"context"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	"github.com/BenMeredithConsult/locagri-apps/ent"
)

var db *ent.Client

type repo struct {
	ctx context.Context
}

func NewAppConstRepo(adaptor *database.Adapter) gateways.AppConstRepo {
	db = adaptor.DB
	return &repo{ctx: context.Background()}
}

func (r *repo) SelectLanguages() ([]*ent.Language, error) {
	return db.Language.Query().All(r.ctx)
}

func (r *repo) SelectCountries() ([]*ent.Country, error) {
	return db.Country.Query().All(r.ctx)
}

func (r *repo) SelectNationalities() ([]*ent.Nationality, error) {
	return db.Nationality.Query().All(r.ctx)
}
