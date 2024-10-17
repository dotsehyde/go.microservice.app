package auth

import (
	"sample/app/adapters/gateways"
	"sample/app/adapters/presenters"
	"sample/app/domain/model"
	"sample/app/domain/requestdto"
)

type service struct {
	repo gateways.AuthRepo
}

func NewAuthService(repo gateways.AuthRepo) gateways.AuthService {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(request requestdto.RegisterRequest) (*presenters.RegisterResponse, error) {
	newUser := model.User{
		ID:       len(data) + 1,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	user, err := s.repo.Save(newUser)
	return &presenters.RegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}, err

}

func (s *service) Login(request requestdto.LoginRequest) (*presenters.LoginResponse, error) {
	user, err := s.repo.Get(request)
	return &presenters.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
	}, err
}

func (s *service) AllUsers() (presenters.PaginatedResponse, error) {
	users, err := s.repo.GetAll()
	return presenters.PaginatedResponse{
		Data:          users,
		TotalElements: len(users),
		HasMore:       false,
	}, err
}
