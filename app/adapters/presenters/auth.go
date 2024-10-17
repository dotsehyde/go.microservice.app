package presenters

type (
	LoginResponse struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}
	RegisterResponse struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}
)
