package usecases

import (
	"context"

	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error)
}

type eventService struct {
	eventRepo repositories.EventRepositories
	config    *configs.Config
}

// CreateEvent implements EventUseCase.
func (e *eventService) CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	return e.eventRepo.Create(ctx, req)
}

func ProvideEventService(eventRepo repositories.EventRepositories, config *configs.Config) EventUseCase {
	return &eventService{
		eventRepo: eventRepo,
		config:    config,
	}
}
