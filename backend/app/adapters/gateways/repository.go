package gateways

import (
	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/ent"
)

type (
	AuthRepo interface {
		SelectByPhone(phone string) (*ent.User, error)
		SelectById(id int) (*ent.User, error)
		ExistByPhone(phone string) (bool, error)
		Insert(req *requestdto.CreateUserRequest) (*ent.User, error)
	}
	FarmerRepo interface {
		SelectByPhone(phone string) (*ent.User, error)
		SelectById(id int) (*ent.User, error)
		UpdateIDInfo(id int, nationality string, idNumber string, idPhoto string, idType string) error
		UpdateUserInfo(id int, req *requestdto.FarmerUpdateInfoRequest) (*ent.User, error)
		UpdateProfilePhoto(id int, photo string) error
	}
	AppConstRepo interface {
		SelectLanguages() ([]*ent.Language, error)
		SelectCountries() ([]*ent.Country, error)
		SelectNationalities() ([]*ent.Nationality, error)
	}
)
