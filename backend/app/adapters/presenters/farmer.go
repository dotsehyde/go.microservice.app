package presenters

import "github.com/BenMeredithConsult/locagri-apps/ent"

type (
	FarmerUser struct {
		ID           int    `json:"id"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		Phone        string `json:"phone"`
		Role         string `json:"role,omitempty"`
		Language     string `json:"language,omitempty"`
		Country      string `json:"country,omitempty"`
		ProfilePhoto string `json:"profilePhoto"`
		IDPhoto      string `json:"idPhoto,omitempty"`
		Nationality  string `json:"nationality,omitempty"`
		IDNumber     string `json:"idNumber"`
		IdType       string `json:"idType"`
		IsVerified   bool   `json:"isVerified"`
		IsBlocked    bool   `json:"isBlocked"`
		Reason       string `json:"reason,omitempty"`
		Address      string `json:"address"`
		City         string `json:"city"`
		CreatedAt    any    `json:"createdAt,omitempty"`
		UpdatedAt    any    `json:"updatedAt,omitempty"`
	}
)

func NewFramerProfile(user *ent.User) *FarmerUser {
	return &FarmerUser{
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
		IdType:       user.IDType,
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
