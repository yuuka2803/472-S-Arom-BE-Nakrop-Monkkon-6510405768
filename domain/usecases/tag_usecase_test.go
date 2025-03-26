package usecases_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/stretchr/testify/assert"

	"github.com/kritpi/arom-web-services/domain/usecases"
	"github.com/kritpi/arom-web-services/domain/usecases/test/mock_repos"
)

func TestGetByID(t *testing.T){
	MockTagRepo := &mockrepos.MockTagRepo{}
	ctx := context.Background()

	t.Run("Success", func(t *testing.T){
		MockTagRepo.On("GetByID", ctx, "1").Return(&models.Tag{}, nil)

		tagService := usecases.ProvideTagService(MockTagRepo, &configs.Config{})
		tag, err := tagService.GetByIDTag(ctx, "1")

		if err != nil {
			t.Errorf("Error was not expected: %v", err)
		}

		if reflect.TypeOf(tag) != reflect.TypeOf(&models.Tag{}) {
			t.Errorf("Expected type %v but got %v", reflect.TypeOf(&models.Tag{}), reflect.TypeOf(tag))
		}
		
		assert.Equal(t, tag, &models.Tag{})
		MockTagRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T){
		MockTagRepo.On("GetByID", ctx, "1").Return(nil, assert.AnError)

		tagService := usecases.ProvideTagService(MockTagRepo, &configs.Config{})
		tag, err := tagService.GetByIDTag(ctx, "1")

		if err == nil {
			t.Errorf("Error was expected but got nil")
		}

		if tag != nil {
			t.Errorf("Expected nil but got %v", tag)
		}
		MockTagRepo.AssertExpectations(t)
	})

}

func TestGetByUserID(t *testing.T){
	MockTagRepo := &mockrepos.MockTagRepo{}
	ctx := context.Background()

	t.Run("Success", func(t *testing.T){
		MockTagRepo.On("GetByUserID", ctx, "1").Return([]*models.Tag{}, nil)

		tagService := usecases.ProvideTagService(MockTagRepo, &configs.Config{})
		tags, err := tagService.GetByUserIDTag(ctx, "1")

		if err != nil {
			t.Errorf("Error was not expected: %v", err)
		}

		if reflect.TypeOf(tags) != reflect.TypeOf([]*models.Tag{}) {
			t.Errorf("Expected type %v but got %v", reflect.TypeOf([]*models.Tag{}), reflect.TypeOf(tags))
		}
		
		assert.Equal(t, tags, []*models.Tag{})
		MockTagRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T){
		MockTagRepo.On("GetByUserID", ctx, "1").Return(nil, assert.AnError)

		tagService := usecases.ProvideTagService(MockTagRepo, &configs.Config{})
		tags, err := tagService.GetByUserIDTag(ctx, "1")

		if err == nil {
			t.Errorf("Error was expected but got nil")
		}

		if tags != nil {
			t.Errorf("Expected nil but got %v", tags)
		}
		MockTagRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T){
	MockTagRepo := &mockrepos.MockTagRepo{}
	ctx := context.Background()

	t.Run("Success", func(t *testing.T){
		MockTagRepo.On("Delete", ctx, "1").Return(nil)

		tagService := usecases.ProvideTagService(MockTagRepo, &configs.Config{})
		err := tagService.DeleteTag(ctx, "1")

		if err != nil {
			t.Errorf("Error was not expected: %v", err)
		}

		MockTagRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T){
		MockTagRepo.On("Delete", ctx, "1").Return(assert.AnError)

		tagService := usecases.ProvideTagService(MockTagRepo, &configs.Config{})
		err := tagService.DeleteTag(ctx, "1")

		if err == nil {
			t.Errorf("Error was expected but got nil")
		}

		MockTagRepo.AssertExpectations(t)
	})
}
