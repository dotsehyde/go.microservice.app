package gateways

import (
	"sample/app/domain/model"
	"sample/app/domain/requestdto"
)

type (
	AuthRepo interface {
		Get(requestdto.LoginRequest) (model.User, error)
		Save(model.User) (model.User, error)
		GetAll() ([]model.User, error)
	}
)
