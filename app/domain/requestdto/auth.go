package requestdto

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required|email"`
		Password string `json:"password" validate:"required|min:6"`
	}
	RegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required|email"`
		Password string `json:"password" validate:"required|min:6"`
	}
)
