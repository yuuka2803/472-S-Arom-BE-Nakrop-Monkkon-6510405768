package mockrepos

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/models"
)


type MockTagRepo struct {
	mock.Mock
}

func (m *MockTagRepo) Create(ctx context.Context, req *requests.CreateTagRequest) (*models.Tag, error) {
	ret := m.Called(ctx, req)

	var r0 *models.Tag
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Tag)
	}
	r1 := ret.Error(1)

	return r0, r1
}

func (m *MockTagRepo) GetByID(ctx context.Context, id string) (*models.Tag, error) {
	ret := m.Called(ctx, id)

	var r0 *models.Tag
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Tag)
	}
	r1 := ret.Error(1)

	return r0, r1
}

func (m *MockTagRepo) GetByUserID(ctx context.Context, id string) ([]*models.Tag, error) {
	ret := m.Called(ctx, id)

	var r0 []*models.Tag
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*models.Tag)
	}
	r1 := ret.Error(1)

	return r0, r1
}

func (m *MockTagRepo) Delete(ctx context.Context, id string) error {
	ret := m.Called(ctx, id)

	return ret.Error(0)
}

func (m *MockTagRepo) Update(ctx context.Context, req *requests.UpdateTagRequest, id string) error {
	ret := m.Called(ctx, req, id)

	return ret.Error(0)
}


