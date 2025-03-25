package rest

import (

	"github.com/gofiber/fiber/v2"

	"github.com/kritpi/arom-web-services/domain/usecases"
)

type TagHandler interface{
	CreateTag(c *fiber.Ctx) error
	GetByIDTag(c *fiber.Ctx) error
	GetByUserIDTag(c *fiber.Ctx) error
	UpdateTag(c *fiber.Ctx) error
	DeleteTag(c *fiber.Ctx) error
}

type tagHandler struct {
	service usecases.TagUseCase
}

