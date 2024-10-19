package gateways

import (
	"mime/multipart"
	"time"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/presenters"
	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/ent"
)

type (
	StorageService interface {
		UploadFile(dir string, f *multipart.FileHeader) (string, error)
		UploadCSVFile(dir string, f *multipart.FileHeader) (string, error)
		UploadResizeImage(dir string, f *multipart.FileHeader, width, height int) (string, error)
		UploadFiles(dir string, files []*multipart.FileHeader) ([]string, error)
		Disk(disk string) StorageService
		ExecuteTask(data any, taskType string)
		Listen()
		Done()
		Close()
	}
	EventProducer interface {
		Queue(queue string, payload any)
	}
	CacheService interface {
		Set(key string, value any, ttl time.Duration) error
		Get(key string, obj any) error
		Has(key string) bool
		Delete(key string) error
	}
	AppConstService interface {
		GetLanguages() ([]*ent.Language, error)
		GetCountries() ([]*ent.Country, error)
		GetNationalities() ([]*ent.Nationality, error)
	}

	AuthService interface {
		VerifyOTP(phone string, otp string) (*presenters.UserAuthPresenter, error)
		SendOTP(phone string) (string, error)
		ResetAccount(phone string) (string, error)
		Login(phone string) (string, error)
		Register(req *requestdto.CreateUserRequest) (string, error)
		RefreshToken(refreshToken string) (*presenters.UserAuthPresenter, error)
	}

	FarmerService interface {
		GetProfile(id int) (*presenters.FarmerUser, error)
		UpdateUserInfo(id int, req *requestdto.FarmerUpdateInfoRequest) (*presenters.FarmerUser, error)
		UpdateProfilePhoto(id int, req *requestdto.FarmerPhotoRequest) (string, error)
		UpdateIDInfo(id int, req *requestdto.UploadIDRequest) (string, error)
	}
)
