package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/usecases"
	"github.com/kritpi/arom-web-services/domain/usecases/test/mock_repos"
	mock_mailer "github.com/kritpi/arom-web-services/internal/infrastrutures/mailer/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gopkg.in/gomail.v2"
)

type MockDialer struct {
	mock.Mock
}

func (m *MockDialer) DialAndSend(msg *gomail.Message) error {
	args := m.Called(msg)
	return args.Error(0)
}

func Test_eventService_CreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEventRepo := new(mock_repos.MockEventRepository)
	mockUserRepo := new(mock_repos.MockUserRepository)
	mockConfig := &configs.Config{}
	mockMailer := mock_mailer.NewMockMailer(ctrl)

	eventUsecase := usecases.ProvideEventService(mockEventRepo, mockUserRepo, mockConfig, mockMailer)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		t.Log("Running Success Test Case")
		mockEventRepo.ExpectedCalls = nil

		createEventRequest := &requests.CreateEventRequest{
			Title:        "Team Meeting",
			Description:  "Discuss project updates and next steps",
			Start:        "2025-04-01T10:00:00Z",
			End:          "2025-04-01T18:00:00Z",
			Tag:          uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
			Notification: true,
			Reminder:     "None",
			UserId:       uuid.MustParse("a69b79ad-fb64-498b-ae8e-1aa7c6b57e99"),
		}

		startTime, err := time.Parse(time.RFC3339, createEventRequest.Start)
		if err != nil {
			t.Fatalf("Failed to parse start time: %v", err)
		}

		endTime, err := time.Parse(time.RFC3339, createEventRequest.End)
		if err != nil {
			t.Fatalf("Failed to parse end time: %v", err)
		}

		eventID := uuid.New() // สร้าง UUID เดียว

		event := &models.Event{
			Id:           eventID, // ใช้ UUID เดียวกัน
			Title:        createEventRequest.Title,
			Description:  createEventRequest.Description,
			Start:        startTime,
			End:          endTime,
			Tag:          createEventRequest.Tag.String(),
			Type:         "event",
			Completed:    false,
			Notification: createEventRequest.Notification,
			Reminder:     createEventRequest.Reminder,
			UserId:       createEventRequest.UserId,
		}

		// ตั้งค่า mock behavior สำหรับ GetUserByUserID
		mockUserRepo.On("GetUserByUserID", ctx, &requests.SendEmailRequest{ID: createEventRequest.UserId, Email: ""}).Return(&models.User{}, nil).Once()

		// แก้ไข mock.MatchedBy ให้รับ requests.CreateEventRequest
		mockEventRepo.On("Create", ctx, mock.MatchedBy(func(req *requests.CreateEventRequest) bool {
			return req.Title == "Team Meeting" && req.UserId == createEventRequest.UserId
		})).Return(event, nil).Once()

		resultEvent, err := eventUsecase.CreateEvent(ctx, createEventRequest)
		assert.NoError(t, err)
		assert.NotNil(t, resultEvent)
		assert.Equal(t, event, resultEvent) // Verify the returned event matches the expected one
		mockEventRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t) // ตรวจสอบ mockUserRepo
	})
}

func Test_Send_Email_When_Notification_Enabled(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Mock Repositories
	mockEventRepo := new(mock_repos.MockEventRepository)
	mockUserRepo := new(mock_repos.MockUserRepository)
	mockMailer := mock_mailer.NewMockMailer(ctrl)
	mockConfig := &configs.Config{}

	// สร้าง EventService ด้วย mock dependencies
	eventUsecase := usecases.ProvideEventService(mockEventRepo, mockUserRepo, mockConfig, mockMailer)

	// กำหนดค่าการทดสอบ
	ctx := context.TODO()
	createEventRequest := &requests.CreateEventRequest{
		Title:        "Team Meeting",
		Description:  "Discuss project updates and next steps",
		Start:        "2025-04-01T10:00:00Z",
		End:          "2025-04-01T18:00:00Z",
		Tag:          uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
		Notification: true,
		Reminder:     "None",
		UserId:       uuid.MustParse("a69b79ad-fb64-498b-ae8e-1aa7c6b57e99"),
	}

	// แปลงเวลาเริ่มต้นและเวลาสิ้นสุด
	startTime, err := time.Parse(time.RFC3339, createEventRequest.Start)
	if err != nil {
		t.Fatalf("Failed to parse start time: %v", err)
	}

	endTime, err := time.Parse(time.RFC3339, createEventRequest.End)
	if err != nil {
		t.Fatalf("Failed to parse end time: %v", err)
	}

	eventID := uuid.New() // สร้าง UUID เดียว
	// สร้าง Event object
	event := &models.Event{
		Id:           eventID, // ใช้ UUID เดียวกัน
		Title:        createEventRequest.Title,
		Description:  createEventRequest.Description,
		Start:        startTime,
		End:          endTime,
		Tag:          createEventRequest.Tag.String(),
		Type:         "event",
		Completed:    false,
		Notification: createEventRequest.Notification,
		Reminder:     createEventRequest.Reminder,
		UserId:       createEventRequest.UserId,
	}

	// สร้าง User object
	user := &models.User{
		ID:    uuid.MustParse("a69b79ad-fb64-498b-ae8e-1aa7c6b57e99"),
		Email: "nutta.nut2327@gmail.com",
	}

	// Mock repositories
	mockEventRepo.On("Create", ctx, createEventRequest).Return(event, nil)
	mockUserRepo.On("GetUserByUserID", ctx, &requests.SendEmailRequest{ID: uuid.MustParse("a69b79ad-fb64-498b-ae8e-1aa7c6b57e99")}).Return(user, nil)

	// Mocking the SendEmail method
	mockMailer.EXPECT().SendEmail(user.Email, gomock.Eq("Arom Notification: Team Meeting"), gomock.Any()).Return(nil)

	// เรียกใช้ฟังก์ชัน CreateEvent
	_, err = eventUsecase.CreateEvent(ctx, createEventRequest)

	// Assertions
	assert.NoError(t, err) // ตรวจสอบว่าไม่มี error
}

func Test_Send_Email_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock Dependencies
	mockEventRepo := new(mock_repos.MockEventRepository)
	mockUserRepo := new(mock_repos.MockUserRepository)
	mockMailer := mock_mailer.NewMockMailer(ctrl)
	mockConfig := &configs.Config{}

	eventUsecase := usecases.ProvideEventService(mockEventRepo, mockUserRepo, mockConfig, mockMailer)

	ctx := context.TODO()
	createEventRequest := &requests.CreateEventRequest{
		Title:        "Team Meeting",
		Description:  "Discuss project updates and next steps",
		Start:        "2025-04-01T10:00:00Z",
		End:          "2025-04-01T18:00:00Z",
		Tag:          uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
		Notification: true,
		Reminder:     "None",
		UserId:       uuid.MustParse("a69b79ad-fb64-498b-ae8e-1aa7c6b57e99"),
	}

	startTime, err := time.Parse(time.RFC3339, createEventRequest.Start)
	require.NoError(t, err)

	endTime, err := time.Parse(time.RFC3339, createEventRequest.End)
	require.NoError(t, err)

	eventID := uuid.New()
	event := &models.Event{
		Id:           eventID,
		Title:        createEventRequest.Title,
		Description:  createEventRequest.Description,
		Start:        startTime,
		End:          endTime,
		Tag:          createEventRequest.Tag.String(),
		Type:         "event",
		Completed:    false,
		Notification: createEventRequest.Notification,
		Reminder:     createEventRequest.Reminder,
		UserId:       createEventRequest.UserId,
	}

	user := &models.User{
		ID:    createEventRequest.UserId,
		Email: "user@example.com",
	}

	// Mock Repositories
	mockEventRepo.On("Create", ctx, createEventRequest).Return(event, nil)
	mockUserRepo.On("GetUserByUserID", ctx, &requests.SendEmailRequest{ID: user.ID}).Return(user, nil)

	// Mock SendEmail ให้ล้มเหลว
	mockMailer.EXPECT().
		SendEmail(user.Email, gomock.Eq("Arom Notification: Team Meeting"), gomock.Any()).
		Return(errors.New("failed to connect to email server"))

	// Call CreateEvent
	_, err = eventUsecase.CreateEvent(ctx, createEventRequest)

	// Assert ต้องมี error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to connect to email server")
}

func Test_eventService_UpdateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepo := new(mock_repos.MockEventRepository)
	mockUserRepo := new(mock_repos.MockUserRepository)
	mockMailer := mock_mailer.NewMockMailer(ctrl)
	mockConfig := &configs.Config{}

	eventUsecase := usecases.ProvideEventService(mockEventRepo, mockUserRepo, mockConfig, mockMailer)
	ctx := context.Background()

	eventID := uuid.New()
	updateEventRequest := &requests.UpdateEventRequest{
		Title:        "Updated Meeting",
		Description:  "Updated project discussion",
		Start:        "2025-04-02T10:00:00Z",
		End:          "2025-04-02T12:00:00Z",
		Notification: true,
		Reminder:     "1h",
		UserId:       uuid.MustParse("a69b79ad-fb64-498b-ae8e-1aa7c6b57e99"),
	}

	startTime, _ := time.Parse(time.RFC3339, updateEventRequest.Start)
	endTime, _ := time.Parse(time.RFC3339, updateEventRequest.End)

	updatedEvent := &models.Event{
		Id:           eventID,
		Title:        updateEventRequest.Title,
		Description:  updateEventRequest.Description,
		Start:        startTime,
		End:          endTime,
		Notification: updateEventRequest.Notification,
		Reminder:     updateEventRequest.Reminder,
		UserId:       updateEventRequest.UserId,
	}

	user := &models.User{
		ID:    updateEventRequest.UserId,
		Email: "user@example.com",
	}

	t.Run("Success - Update Event and Send Email", func(t *testing.T) {
		// Mock อัปเดต Event สำเร็จ
		mockEventRepo.On("Update", ctx, updateEventRequest, eventID.String()).Return(updatedEvent, nil).Once()

		// Mock ดึงข้อมูลผู้ใช้สำเร็จ
		mockUserRepo.On("GetUserByUserID", ctx, &requests.SendEmailRequest{ID: updateEventRequest.UserId}).
			Return(user, nil).Once()

		// Mock ส่งอีเมลแจ้งเตือนสำเร็จ
		mockMailer.EXPECT().
			SendEmail(user.Email, gomock.Eq("Update Notification: Updated Meeting"), gomock.Any()).
			Return(nil)

		resultEvent, err := eventUsecase.UpdateEvent(ctx, updateEventRequest, eventID.String())

		assert.NoError(t, err)
		assert.NotNil(t, resultEvent)
		assert.Equal(t, updatedEvent, resultEvent)

		mockEventRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Fail - Update Event Success but Send Email Fails", func(t *testing.T) {
		mockEventRepo.On("Update", ctx, updateEventRequest, eventID.String()).Return(updatedEvent, nil).Once()
		mockUserRepo.On("GetUserByUserID", ctx, &requests.SendEmailRequest{ID: updateEventRequest.UserId}).
			Return(user, nil).Once()

		// Mock ส่งอีเมลแจ้งเตือน **ล้มเหลว**
		mockMailer.EXPECT().
			SendEmail(user.Email, gomock.Eq("Update Notification: Updated Meeting"), gomock.Any()).
			Return(errors.New("failed to send email"))

		resultEvent, err := eventUsecase.UpdateEvent(ctx, updateEventRequest, eventID.String())

		assert.Error(t, err)
		assert.Nil(t, resultEvent)
		assert.EqualError(t, err, "failed to send email")

		mockEventRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Fail - Update Event Fails", func(t *testing.T) {
		mockEventRepo.On("Update", ctx, updateEventRequest, eventID.String()).
			Return((*models.Event)(nil), errors.New("database error")).Once()
	
		resultEvent, err := eventUsecase.UpdateEvent(ctx, updateEventRequest, eventID.String())
	
		assert.Error(t, err)
		assert.Nil(t, resultEvent)
		assert.EqualError(t, err, "database error")
	
		mockEventRepo.AssertExpectations(t)
	})	

}
