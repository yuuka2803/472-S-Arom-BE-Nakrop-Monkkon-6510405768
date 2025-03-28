package mock_repos

import (
	"context"

	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

// GetAll implements repositories.EventRepositories.
func (m *MockEventRepository) GetAll(ctx context.Context) ([]*models.Event, error) {
	panic("unimplemented")
}

// GetByID implements repositories.EventRepositories.
func (m *MockEventRepository) GetByID(ctx context.Context, id string) (*models.Event, error) {
	panic("unimplemented")
}

// GetByUserID implements repositories.EventRepositories.
func (m *MockEventRepository) GetByUserID(ctx context.Context, id string) ([]*models.Event, error) {
	panic("unimplemented")
}

// UpdateStatus implements repositories.EventRepositories.
func (m *MockEventRepository) UpdateStatus(ctx context.Context, req *requests.UpdateStatusEventRequest, id string) error {
	panic("unimplemented")
}

func (m *MockEventRepository) Create(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventRepository) Update(ctx context.Context, req *requests.UpdateEventRequest, id string) (*models.Event, error) {
	args := m.Called(ctx, req, id)
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventRepository) SendReminderEmailBeforeEvent(event *models.Event, user *models.User, reminder string) error {
	args := m.Called(event, user, reminder)
	return args.Error(0)
}
