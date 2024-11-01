package repositories

import (
	"context"

	models "github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type EventRepositories interface {
	Create(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error)
}
