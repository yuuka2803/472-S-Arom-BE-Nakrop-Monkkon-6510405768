package mock_repos

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type MockDiaryRepository struct {
	mock.Mock
}

func (m *MockDiaryRepository) Create(ctx context.Context, req *requests.CreateDiaryRequest) (*models.Diary, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*models.Diary), args.Error(1)
}

func (m *MockDiaryRepository) GetAll(ctx context.Context) ([]*models.Diary, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Diary), args.Error(1)
}

func (m *MockDiaryRepository) GetByID(ctx context.Context, id string) (*models.Diary, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Diary), args.Error(1)
}

func (m *MockDiaryRepository) GetByUserID(ctx context.Context, userID string) ([]*models.Diary, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*models.Diary), args.Error(1)
}

func (m *MockDiaryRepository) GetByDate(ctx context.Context, date string) (*models.Diary, error) {
	args := m.Called(ctx, date)
	return args.Get(0).(*models.Diary), args.Error(1)
}

func (m *MockDiaryRepository) Update(ctx context.Context, req *requests.UpdateDiaryRequest, date string) error {
	args := m.Called(ctx, req, date)
	return args.Error(0)
}
