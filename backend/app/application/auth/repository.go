package auth

import (
	"context"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	"github.com/BenMeredithConsult/locagri-apps/ent"
	"github.com/BenMeredithConsult/locagri-apps/ent/user"
)

var db *ent.Client

type repo struct {
	ctx context.Context
}

func NewAuthRepo(adaptor *database.Adapter) gateways.AuthRepo {
	db = adaptor.DB
	return &repo{ctx: context.Background()}
}

func (r *repo) ExistByPhone(phone string) (bool, error) {
	_, err := db.User.Query().Where(user.Phone(phone)).Only(r.ctx)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *repo) Insert(req *requestdto.CreateUserRequest) (*ent.User, error) {
	return db.User.Create().
		SetPhone(req.Phone).
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetRole(user.Role(req.Role)).
		SetLanguage(user.Language(req.LangCode)).
		SetCountryID(req.CountryId).
		Save(r.ctx)
}

func (r *repo) SelectByPhone(phone string) (*ent.User, error) {
	return db.User.Query().Where(user.Phone(phone)).WithCountry().Only(r.ctx)

}

func (r *repo) SelectById(id int) (*ent.User, error) {
	return db.User.Query().Where(user.ID(id)).WithCountry().Only(r.ctx)
}
