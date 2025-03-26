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

}
