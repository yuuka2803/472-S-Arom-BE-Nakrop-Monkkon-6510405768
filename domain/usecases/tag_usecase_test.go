package usecases_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/stretchr/testify/assert"

	"github.com/kritpi/arom-web-services/domain/usecases"
	mockrepos "github.com/kritpi/arom-web-services/domain/usecases/test/mock_repos"
)

func TestGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		mockTagRepo := &mockrepos.MockTagRepo{}
		mockTagRepo.On("GetByID", ctx, "1").Return(&models.Tag{}, nil)

		tagService := usecases.ProvideTagService(mockTagRepo, &configs.Config{})
		tag, err := tagService.GetByIDTag(ctx, "1")

		if err != nil {
			t.Errorf("Error was not expected: %v", err)
		}

		if reflect.TypeOf(tag) != reflect.TypeOf(&models.Tag{}) {
			t.Errorf("Expected type %v but got %v", reflect.TypeOf(&models.Tag{}), reflect.TypeOf(tag))
		}

		mockTagRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockTagRepo := &mockrepos.MockTagRepo{}
		mockTagRepo.On("GetByID", ctx, "1").Return(nil, nil)

		tagService := usecases.ProvideTagService(mockTagRepo, &configs.Config{})
		tag, err := tagService.GetByIDTag(ctx, "1")

		if err != nil {
			t.Errorf("Error was not expected: %v", err)
		}

		if tag != nil {
			t.Errorf("Expected nil but got %v", tag)
		}

		mockTagRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockTagRepo := &mockrepos.MockTagRepo{}
		mockTagRepo.On("GetByID", ctx, "1").Return(nil, assert.AnError)

		tagService := usecases.ProvideTagService(mockTagRepo, &configs.Config{})
		tag, err := tagService.GetByIDTag(ctx, "1")

		if err == nil {
			t.Errorf("Error was expected but got nil")
		}

		if tag != nil {
			t.Errorf("Expected nil but got %v", tag)
		}

		mockTagRepo.AssertExpectations(t)
	})
}
