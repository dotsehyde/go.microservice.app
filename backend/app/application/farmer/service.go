package farmer

import (
	"fmt"
	"strings"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/presenters"
	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/config"
)

type service struct {
	repo        gateways.FarmerRepo
	cache       gateways.CacheService
	storage     gateways.StorageService
	storagePath string
}

func NewFarmerService(repo gateways.FarmerRepo,
	cacheSrv gateways.CacheService,
	storageSrv gateways.StorageService) gateways.FarmerService {
	return &service{
		repo:        repo,
		cache:       cacheSrv,
		storage:     storageSrv,
		storagePath: "public/farmer",
	}
}

// Update ID information
func (s *service) UpdateIDInfo(id int, req *requestdto.UploadIDRequest) (string, error) {
	u, err := s.repo.SelectById(id)
	if err != nil {
		return "", err
	}
	filePath, err := s.storage.UploadFile(fmt.Sprintf("%s/ID", s.storagePath), req.Photo)
	if err != nil {
		return "", err
	}
	if err := s.repo.UpdateIDInfo(id, req.Nationality, req.IdNumber, filePath, req.IdType); err != nil {
		go s.storage.ExecuteTask(strings.Replace(filePath, config.App().AppURL, "public", 1), "delete_file")
		return "", err
	}
	go s.storage.ExecuteTask(strings.Replace(u.IDPhoto, config.App().AppURL, "public", 1), "delete_file")
	return "ID documents updated successfully", err
}

// Update profile photo
func (s *service) UpdateProfilePhoto(id int, req *requestdto.FarmerPhotoRequest) (string, error) {
	u, err := s.repo.SelectById(id)
	if err != nil {
		return "", err
	}

	//Delete if no photo is uploaded
	if req.Photo == nil {
		go s.storage.ExecuteTask(strings.Replace(u.ProfilePhoto, config.App().AppURL, "public", 1), "delete_file")
		return "Profile photo deleted successfully", nil
	}
	var filePath string
	if u.ProfilePhoto != "" {
		// Update profile photo
		go s.storage.ExecuteTask(strings.Replace(u.ProfilePhoto, config.App().AppURL, "public", 1), "delete_file")
		// Upload new photo
		filePath, err = s.storage.UploadFile(fmt.Sprintf("%s/profile", s.storagePath), req.Photo)
		if err != nil {
			return "", err
		}
	} else {
		// Upload new photo
		filePath, err = s.storage.UploadFile(fmt.Sprintf("%s/profile", s.storagePath), req.Photo)
		if err != nil {
			return "", err
		}
	}
	if err := s.repo.UpdateProfilePhoto(id, filePath); err != nil {
		return "", err
	}
	return "Profile photo updated successfully", err
}

// Get profile information
func (s *service) GetProfile(id int) (*presenters.FarmerUser, error) {
	data, err := s.repo.SelectById(id)
	if err != nil {
		return nil, err
	}
	return presenters.NewFramerProfile(data), nil
}

// Update user information
func (s *service) UpdateUserInfo(id int, req *requestdto.FarmerUpdateInfoRequest) (*presenters.FarmerUser, error) {
	data, err := s.repo.UpdateUserInfo(id, req)
	if err != nil {
		return nil, err
	}
	return presenters.NewFramerProfile(data), nil
}
