package requestdto

import "mime/multipart"

type (
	CreateUserRequest struct {
		Phone     string `json:"phone" validate:"required|phone_with_code|unique:users.phone"`
		FirstName string `json:"firstName" validate:"required|ascii"`
		LastName  string `json:"lastName" validate:"required|ascii"`
		Role      string `json:"role" validate:"required|in:farmer,worker,supplier,officer|string"`
		LangCode  string `json:"langCode" validate:"required|string|in:en,fr"`
		CountryId int    `json:"countryId" validate:"required|numeric"`
	}
	UploadIDRequest struct {
		Nationality string                `json:"nationality" form:"nationality" validate:"required|ascii"`
		IdType      string                `json:"idType" form:"idType" validate:"required|string"`
		IdNumber    string                `json:"idNumber" form:"id_number" validate:"required|string"`
		Photo       *multipart.FileHeader `json:"photo" form:"photo" validate:"image|size:1MB"`
	}

	LoginAndRestRequest struct {
		Phone string `json:"phone" validate:"required|phone_with_code|exists:users.phone"`
	}
	VerifyOTPRequest struct {
		Phone string `json:"phone" validate:"required|phone_with_code"`
		Code  string `json:"code" validate:"required|string|len:6"`
	}
)
