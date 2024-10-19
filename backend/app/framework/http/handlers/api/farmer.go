package api

import (
	"fmt"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/presenters"
	"github.com/BenMeredithConsult/locagri-apps/app/application"
	"github.com/BenMeredithConsult/locagri-apps/app/application/farmer"
	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	"github.com/gofiber/fiber/v3"
)

type farmerHandler struct {
	service gateways.FarmerService
}

func NewFarmerHandler(db *database.Adapter,
	producer gateways.EventProducer,
	storageSrv gateways.StorageService,
	cacheSrv gateways.CacheService) *farmerHandler {
	return &farmerHandler{
		service: farmer.NewFarmerService(farmer.NewFarmerRepo(db),
			cacheSrv, storageSrv),
	}
}

func (h *farmerHandler) GetProfile() fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := int(c.Locals("user").(map[string]interface{})["id"].(float64))
		data, err := h.service.GetProfile(userId)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.SuccessResponse(data))
	}
}

func (h *farmerHandler) UpdateIDInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := int(c.Locals("user").(map[string]interface{})["id"].(float64))
		// Parse formdata to request
		file, err := c.FormFile("photo")
		// 2MB
		if file.Size > 2000000 {
			return c.Status(400).JSON(presenters.ErrorResponse(fmt.Errorf("File size is too large")))
		}
		if err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		request := &requestdto.UploadIDRequest{
			Photo:       file,
			IdNumber:    c.FormValue("idNumber"),
			IdType:      c.FormValue("idType"),
			Nationality: c.FormValue("nationality"),
		}
		// if err := application.BodyParser(c, &request); err != nil {
		// 	return c.Status(400).JSON(presenters.ErrorResponse(err))
		// }
		data, err := h.service.UpdateIDInfo(userId, request)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.MessageResponse(data))
	}
}

func (h *farmerHandler) UpdateProfilePhoto() fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := int(c.Locals("user").(map[string]interface{})["id"].(float64))
		// Parse formdata to request
		file, err := c.FormFile("photo")
		// 1MB
		if file.Size > 1000000 {
			return c.Status(400).JSON(presenters.ErrorResponse(fmt.Errorf("File size is too large")))
		}
		if err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		request := &requestdto.FarmerPhotoRequest{
			Photo: file,
		}
		data, err := h.service.UpdateProfilePhoto(userId, request)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		// return c.JSON(fiber.Map{"message": "success", "file": "hello", "user": userId})

		return c.Status(200).JSON(presenters.MessageResponse(data))
	}
}

func (h *farmerHandler) UpdateUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := int(c.Locals("user").(map[string]interface{})["id"].(float64))
		request := new(requestdto.FarmerUpdateInfoRequest)
		if err := application.BodyParser(c, &request); err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		data, err := h.service.UpdateUserInfo(userId, request)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.SuccessResponse(data))
	}
}
