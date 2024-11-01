package rest

import (
	"fmt"
	"log"
	
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/usecases"
)

type EventHandler interface {
	CreateEvent(c *fiber.Ctx) error
}

type eventHandler struct {
	service usecases.EventUseCase
}

func (p *eventHandler) CreateEvent(c *fiber.Ctx) error {
	var req requests.CreateEventRequest
	log.Println("Request",req)
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	fmt.Printf("CreateEvent request: %v\n", req)
	event, err := p.service.CreateEvent(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(event)
}


func NewEventHandler(service usecases.EventUseCase) EventHandler {
	return &eventHandler{
		service: service,
	}
}