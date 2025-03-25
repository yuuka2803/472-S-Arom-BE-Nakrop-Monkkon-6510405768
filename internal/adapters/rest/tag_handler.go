package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/usecases"
)

type TagHandler interface {
	CreateTag(c *fiber.Ctx) error
	GetByIDTag(c *fiber.Ctx) error
	GetByUserIDTag(c *fiber.Ctx) error
	UpdateTag(c *fiber.Ctx) error
	DeleteTag(c *fiber.Ctx) error
}

type tagHandler struct {
	service usecases.TagUseCase
}

func (p *tagHandler) CreateTag(c *fiber.Ctx) error {
	var req requests.CreateTagRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	tag, err := p.service.CreateTag(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(tag)
}

func (p *tagHandler) GetByIDTag(c *fiber.Ctx) error {
	id := c.Params("id")
	tag, err := p.service.GetByIDTag(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Tag not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(tag)
}

func (p *tagHandler) GetByUserIDTag(c *fiber.Ctx) error {
	id := c.Params("id")
	tags, err := p.service.GetByUserIDTag(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Tag not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(tags)
}

func (p *tagHandler) UpdateTag(c *fiber.Ctx) error {
	var req requests.UpdateTagRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	id := c.Params("id")
	err := p.service.UpdateTag(c.Context(), &req, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Tag with ID %s not found", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tag updated",
	})
}

func (p *tagHandler) DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	err := p.service.DeleteTag(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Tag with ID %s not found", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tag deleted",
	})
}

func NewTagHandler(service usecases.TagUseCase) TagHandler {
	return &tagHandler{
		service: service,
	}
}
