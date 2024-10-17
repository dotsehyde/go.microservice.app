package gateways

import (
	"sample/app/adapters/presenters"
	"sample/app/domain/requestdto"
)

type (
	AuthService interface {
		Login(request requestdto.LoginRequest) (*presenters.LoginResponse, error)
		Register(request requestdto.RegisterRequest) (*presenters.RegisterResponse, error)
		AllUsers() (presenters.PaginatedResponse, error)
	}
)
