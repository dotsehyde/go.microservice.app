package requests

import (
	"net/http"
	"strconv"

	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/config"
	"github.com/SeyramWood/valid"
)

func validConfig() *valid.Config {
	return &valid.Config{

		DB: &valid.Database{
			Host: config.DB().Host,
			Port: func() int {
				dbPort, _ := strconv.Atoi(config.DB().Port)
				if dbPort > 0 {
					return dbPort
				}
				return 3306
			}(),
			Name:     config.DB().Name,
			Username: config.DB().Username,
			Password: config.DB().Password,
			Driver:   config.DB().Driver,
		},
	}
}

// func ValidateAccount() fiber.Handler {
// 	return func(c fiber.Ctx) error {
// 		request := new(requestdto.AccountRequest)
// 		if err := application.BodyParser(c, request); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
// 		}
// 		if err := valid.New(validConfig(strings.ToLower(strings.ReplaceAll(c.Get("X-Tenant-Domain"), " ", "")))).ValidateStruct(request); err != nil {
// 			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(err))
// 		}
// 		return c.Next()
// 	}
// }

// Auth Validation
func ValidateCreateUser(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requestdto.CreateUserRequest)).ValidateRequest(next)
}

func ValidateLogin(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requestdto.LoginAndRestRequest)).ValidateRequest(next)
}

func ValidateVerifyOTP(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requestdto.VerifyOTPRequest)).ValidateRequest(next)
}

func ValidateIDUpload(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requestdto.UploadIDRequest)).ValidateRequest(next)
}

// Farmer Validation
func ValidateFarmerUpdateInfo(next http.Handler) http.Handler {
	return valid.New().RequestStruct(new(requestdto.FarmerUpdateInfoRequest)).ValidateRequest(next)
}

func ValidateFarmerPhoto(next http.Handler) http.Handler {
	return valid.New().RequestStruct(new(requestdto.FarmerPhotoRequest)).ValidateRequest(next)
}

// func ValidateMission(next http.Handler) http.Handler {
// 	return valid.New().RequestStruct(new(requestdto.MissionRequest)).ValidateRequest(next)
// }
// func ValidateApplication(next http.Handler) http.Handler {
// 	return valid.New().RequestStruct(new(requestdto.ApplicationRequest)).ValidateRequest(next)
// }
// func ValidatePersonalTest(next http.Handler) http.Handler {
// 	return valid.New().RequestStruct(new(requestdto.PersonalTestRequest)).ValidateRequest(next)
// }
