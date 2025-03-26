package usecases

import (
	"context"

	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/configs"
)

type TagUseCase interface {
	CreateTag(ctx context.Context, req *requests.CreateTagRequest) (*models.Tag, error)
	GetByIDTag(ctx context.Context, id string) (*models.Tag, error)
	GetByUserIDTag(ctx context.Context, id string) ([]*models.Tag, error)
	DeleteTag(ctx context.Context, id string) error
	UpdateTag(ctx context.Context, req *requests.UpdateTagRequest , id string)(error)
}

type tagService struct {
	tagRepo repositories.TagRepositories
	config  *configs.Config
}

// GetByUserID implements TagUseCase.
func (t *tagService) GetByUserIDTag(ctx context.Context, id string) ([]*models.Tag, error) {
	return t.tagRepo.GetByUserID(ctx, id)
}

// GetByID implements TagUseCase.	
func (t *tagService) GetByIDTag(ctx context.Context, id string) (*models.Tag, error) {
	return t.tagRepo.GetByID(ctx, id)
}

// Create implements TagUseCase.	
func (t *tagService) CreateTag(ctx context.Context, req *requests.CreateTagRequest) (*models.Tag, error) {
	return t.tagRepo.Create(ctx, req)
}

func (t *tagService) DeleteTag(ctx context.Context, id string) error {
	return t.tagRepo.Delete(ctx, id)
}

func (t *tagService) UpdateTag(ctx context.Context, req *requests.UpdateTagRequest, id string) error {
	return t.tagRepo.Update(ctx, req, id)
}

func ProvideTagService(tagRepo repositories.TagRepositories, config *configs.Config) TagUseCase {
	return &tagService{
		tagRepo: tagRepo,
		config:    config,
	}
}