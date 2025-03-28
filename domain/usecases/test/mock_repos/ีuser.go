package mock_repos

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type MockUserRepository struct {
	mock.Mock
}

// Mock for CreateUser
func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*models.User), args.Error(1)
}

// Mock for GetUserByUsername
func (m *MockUserRepository) GetUserByUsername(ctx context.Context, req *requests.LoginRequest) (*models.User, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*models.User), args.Error(1)
}

// Mock for GetUserByUserID
func (m *MockUserRepository) GetUserByUserID(ctx context.Context, req *requests.SendEmailRequest) (*models.User, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*models.User), args.Error(1)
}
