package requestdto

import "mime/multipart"

type (
	FarmerPhotoRequest struct {
		Photo *multipart.FileHeader `json:"photo" form:"photo" validate:"required|image|size:1MB"`
	}

	FarmerUpdateInfoRequest struct {
		FirstName string `json:"firstName" validate:"required|string"`
		LastName  string `json:"lastName" validate:"required|string"`
		City      string `json:"city" validate:"string"`
		Address   string `json:"address" validate:"string"`
		CountryId int    `json:"countryId" validate:"required|numeric"`
		LangCode  string `json:"langCode" validate:"required|string|in:en,fr"`
	}
)
