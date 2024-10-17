package auth

import (
	"errors"
	"log/slog"
	"sample/app/adapters/gateways"
	"sample/app/domain/model"
	"sample/app/domain/requestdto"
)

type repo struct {
}

func NewAuthRepo() gateways.AuthRepo {
	return &NewRepo{}
}

var data = []model.User{}

func (r *repo) Save(user model.User) (model.User, error) {
	data = append(data, user)
	return user, nil
}

func (r *repo) Get(req requestdto.LoginRequest) (model.User, error) {
	for _, user := range data {
		if user.Email == req.Email && user.Password == req.Password {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

type NewRepo struct {
}

func (r *NewRepo) Save(user model.User) (model.User, error) {
	slog.Info("NewRepo Save")
	return user, nil
}

func (r *NewRepo) Get(req requestdto.LoginRequest) (model.User, error) {
	slog.Info("NewRepo Get")
	return model.User{}, errors.New("user not found")
}

func (r *NewRepo) GetAll() ([]model.User, error) {
	slog.Info("NewRepo GetAll")
	return data, nil
}
