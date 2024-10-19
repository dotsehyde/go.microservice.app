package farmer

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

func NewFarmerRepo(adaptor *database.Adapter) gateways.FarmerRepo {
	db = adaptor.DB
	return &repo{ctx: context.Background()}
}

func (r *repo) SelectByPhone(phone string) (*ent.User, error) {
	return db.User.Query().Where(user.Phone(phone)).WithCountry().Only(r.ctx)
}

func (r *repo) SelectById(id int) (*ent.User, error) {
	return db.User.Query().Where(user.ID(id)).WithCountry().Only(r.ctx)
}

func (r *repo) UpdateIDInfo(id int, nationality string, idNumber string, idPhoto string, idType string) error {
	_, err := db.User.UpdateOneID(id).
		SetIDPhoto(idPhoto).
		SetIDType(idType).
		SetNationality(nationality).
		SetIDNumber(idNumber).
		Save(r.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateProfilePhoto(id int, photo string) error {
	_, err := db.User.UpdateOneID(id).
		SetProfilePhoto(photo).
		Save(r.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateUserInfo(id int, req *requestdto.FarmerUpdateInfoRequest) (*ent.User, error) {
	data, err := db.User.UpdateOneID(id).
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetAddress(req.Address).
		SetCity(req.City).
		SetCountryID(req.CountryId).
		SetLanguage(user.Language(req.LangCode)).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.SelectById(data.ID)

}
