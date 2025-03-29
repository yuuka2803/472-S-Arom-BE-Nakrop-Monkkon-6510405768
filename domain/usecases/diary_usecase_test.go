package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/usecases"
	"github.com/kritpi/arom-web-services/domain/usecases/test/mock_repos"
	"github.com/stretchr/testify/assert"
)

func Test_diaryService_CreateDiary(t *testing.T) {
	mockDiaryRepo := new(mock_repos.MockDiaryRepository)
	diaryUsecase := usecases.ProvideDiaryService(mockDiaryRepo, nil)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		mockDiaryRepo.ExpectedCalls = nil

		diaryRequest := &requests.CreateDiaryRequest{
			Mood:        "Happy",
			Emotions:    []string{"Excited", "Grateful"},
			Description: "Had a great day at the park.",
			UserID:      "user123",
			Images:      []string{"image1.jpg", "image2.jpg"},
		}

		expectedDiary := &models.Diary{
			Mood:        "Happy",
			Emotions:    []string{"Excited", "Grateful"},
			Description: "Had a great day at the park.",
			UserID:      uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
		}

		mockDiaryRepo.On("Create", ctx, diaryRequest).Return(expectedDiary, nil).Once()

		diary, err := diaryUsecase.CreateDiary(ctx, diaryRequest, nil)

		assert.NoError(t, err)
		assert.NotNil(t, diary)
		assert.Equal(t, expectedDiary.Mood, diary.Mood)
		mockDiaryRepo.AssertExpectations(t)
	})	

	t.Run("Failure", func(t *testing.T) {
		mockDiaryRepo.ExpectedCalls = nil

		diaryRequest := &requests.CreateDiaryRequest{
			Mood: "Sad",
		}
		expectedError := errors.New("database error")

		mockDiaryRepo.On("Create", ctx, diaryRequest).Return((*models.Diary)(nil), expectedError).Once()

		_, err := diaryUsecase.CreateDiary(ctx, diaryRequest, nil)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockDiaryRepo.AssertExpectations(t)
	})

}

func Test_diaryService_UpdateDiary(t *testing.T) {
	mockDiaryRepo := new(mock_repos.MockDiaryRepository)
	ctx := context.Background()
	diaryUsecase := usecases.ProvideDiaryService(mockDiaryRepo, nil)
	t.Run("Success", func(t *testing.T) {
		mockDiaryRepo.ExpectedCalls = nil

		updateRequest := &requests.UpdateDiaryRequest{
			Mood:        "Calm",
			Emotions:    []string{"Relaxed"},
			Description: "Had a peaceful day.",
			Images:      []string{"updated_image.jpg"},
		}

		mockDiaryRepo.On("Update", ctx, updateRequest, "2025-03-26").Return(nil).Once()

		err := diaryUsecase.UpdateDiary(ctx, updateRequest, "2025-03-26", nil)
		assert.NoError(t, err)
		mockDiaryRepo.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockDiaryRepo.ExpectedCalls = nil

		updateRequest := &requests.UpdateDiaryRequest{
			Mood: "Anxious",
		}
		expectedError := errors.New("update failed")

		mockDiaryRepo.On("Update", ctx, updateRequest, "2025-03-26").Return(expectedError).Once()

		err := diaryUsecase.UpdateDiary(ctx, updateRequest, "2025-03-26", nil)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockDiaryRepo.AssertExpectations(t)
	})
}

