package presenters

import "github.com/BenMeredithConsult/locagri-apps/ent"

type (
	UserAuthPresenter struct {
		User               any    `json:"user,omitempty"`
		Token              string `json:"token,omitempty"`
		TokenExpiry        int64  `json:"tokenExpiry,omitempty"`
		RefreshToken       string `json:"refreshToken,omitempty"`
		RefreshTokenExpiry int64  `json:"refreshTokenExpiry,omitempty"`
	}

	AuthUser struct {
		ID           int    `json:"id"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		Phone        string `json:"phone"`
		Role         string `json:"role,omitempty"`
		Language     string `json:"language,omitempty"`
		Country      string `json:"country,omitempty"`
		ProfilePhoto string `json:"profilePhoto,omitempty"`
		IDPhoto      string `json:"idPhoto,omitempty"`
		Nationality  string `json:"nationality,omitempty"`
		IDNumber     string `json:"idNumber"`
		IsVerified   bool   `json:"isVerified"`
		IsBlocked    bool   `json:"isBlocked"`
		Reason       string `json:"reason,omitempty"`
		Address      string `json:"address,omitempty"`
		City         string `json:"city,omitempty"`
		CreatedAt    any    `json:"createdAt,omitempty"`
		UpdatedAt    any    `json:"updatedAt,omitempty"`
	}

	AuthSession struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Phone string `json:"phone,omitempty"`
		Photo string `json:"photo,omitempty"`
	}
)

func NewAuthUser(user *ent.User) *AuthUser {
	return &AuthUser{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Phone:        user.Phone,
		Role:         user.Role.String(),
		Language:     user.Language.String(),
		Country:      user.Edges.Country.Name,
		ProfilePhoto: user.ProfilePhoto,
		IDPhoto:      user.IDPhoto,
		IDNumber:     user.IDNumber,
		Address:      user.Address,
		City:         user.City,
		Reason:       user.Reason,
		Nationality:  user.Nationality,
		IsVerified:   user.IsVerified,
		IsBlocked:    user.IsBlocked,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
