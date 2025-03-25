package usecases

import (
	"context"

	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/configs"
)

type TagUseCase interface {
	Create(ctx context.Context, req *requests.CreateTagRequest) (*models.Tag, error)
	GetByID(ctx context.Context, id string) (*models.Tag, error)
	GetByUserID(ctx context.Context, id string) ([]*models.Tag, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, req *requests.UpdateTagRequest , id string)(error)
}

type tagService struct {
	tagRepo repositories.TagRepositories
	config  *configs.Config
}

// GetByUserID implements TagUseCase.
func (t *tagService) GetByUserID(ctx context.Context, id string) ([]*models.Tag, error) {
	return t.tagRepo.GetByUserID(ctx, id)
}

// GetByID implements TagUseCase.	
func (t *tagService) GetByID(ctx context.Context, id string) (*models.Tag, error) {
	return t.tagRepo.GetByID(ctx, id)
}

// Create implements TagUseCase.	
func (t *tagService) Create(ctx context.Context, req *requests.CreateTagRequest) (*models.Tag, error) {
	return t.tagRepo.Create(ctx, req)
}

func (t *tagService) Delete(ctx context.Context, id string) error {
	return t.tagRepo.Delete(ctx, id)
}

func (t *tagService) Update(ctx context.Context, req *requests.UpdateTagRequest, id string) error {
	return t.tagRepo.Update(ctx, req, id)
}