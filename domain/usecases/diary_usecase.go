package usecases

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/internal/adapters/pg"
)

type DiaryUseCase interface {
	CreateDiary(ctx context.Context, req *requests.CreateDiaryRequest, files []*multipart.FileHeader) (*models.Diary, error)
	GetAllDiary(ctx context.Context) ([]*models.Diary, error)
	GetDiaryByID(ctx context.Context, id string) (*models.Diary, error)
	GetDiaryByUserID(ctx context.Context, userID string) ([]*models.Diary, error)
	GetDiaryByDate(ctx context.Context, date string) (*models.Diary, error)
	UpdateDiary(ctx context.Context, req *requests.UpdateDiaryRequest, date string, files []*multipart.FileHeader) error
}

type diaryService struct {
	diaryRepo repositories.DiaryRepositories
	config    *configs.Config
}

func (d *diaryService) CreateDiary(ctx context.Context, req *requests.CreateDiaryRequest, files []*multipart.FileHeader) (*models.Diary, error) {
	var imageURLs []string
	for _, fileHeader := range files {
		if fileHeader == nil {
			continue
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close() // ✅ Ensure file is closed after function exits

		url, err := pg.UploadImageToSupabaseV2(file, fileHeader.Filename, d.config.SUPABASE_BUCKET, d.config)
		if err != nil {
			return nil, fmt.Errorf("failed to upload file: %w", err)
		}

		imageURLs = append(imageURLs, url)
	}

	// Save image URLs in request
	req.Images = imageURLs
	return d.diaryRepo.Create(ctx, req)
}

func (d *diaryService) GetAllDiary(ctx context.Context) ([]*models.Diary, error) {
	return d.diaryRepo.GetAll(ctx)
}
func (d *diaryService) GetDiaryByID(ctx context.Context, id string) (*models.Diary, error) {
	return d.diaryRepo.GetByID(ctx, id)
}

func (d *diaryService) GetDiaryByUserID(ctx context.Context, userID string) ([]*models.Diary, error) {
	return d.diaryRepo.GetByUserID(ctx, userID)
}

func (d *diaryService) GetDiaryByDate(ctx context.Context, date string) (*models.Diary, error) {
	return d.diaryRepo.GetByDate(ctx, date)
}

func (d *diaryService) UpdateDiary(ctx context.Context, req *requests.UpdateDiaryRequest, date string, files []*multipart.FileHeader) error {
	imageURLs := req.ImagesURL
	for _, fileHeader := range files {
		if fileHeader == nil {
			continue
		}

		file, err := fileHeader.Open()
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		url, err := pg.UploadImageToSupabaseV2(file, fileHeader.Filename, d.config.SUPABASE_BUCKET, d.config)
		if err != nil {
			return fmt.Errorf("failed to upload file: %w", err)
		}

		imageURLs = append(imageURLs, url)
	}
	req.Images = imageURLs
	return d.diaryRepo.Update(ctx, req, date)
}

func ProvideDiaryService(diaryRepo repositories.DiaryRepositories, config *configs.Config) DiaryUseCase {
	return &diaryService{
		diaryRepo: diaryRepo,
		config:    config,
	}
}
