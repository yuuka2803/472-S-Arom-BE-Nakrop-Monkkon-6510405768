package repositories

import (
	"context"

	models "github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type TagRepositories interface {
	Create(ctx context.Context, req *requests.CreateTagRequest) (*models.Tag, error)
	GetByID(ctx context.Context, id string) (*models.Tag, error)
	GetByUserID(ctx context.Context, id string) ([]*models.Tag, error)
	Update(ctx context.Context, req *requests.UpdateTagRequest , id string)(error)
	Delete(ctx context.Context, id string) error
}
