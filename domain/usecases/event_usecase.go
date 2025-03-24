package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"gopkg.in/gomail.v2"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error)
	GetAllEvent(ctx context.Context) ([]*models.Event, error)
	GetByIDEvent(ctx context.Context, id string) (*models.Event, error)
	GetByUserIDEvent(ctx context.Context, id string) ([]*models.Event, error)
	UpdateDateEvent(ctx context.Context, req *requests.UpdateEventRequest, id string) error
}

type eventService struct {
	eventRepo repositories.EventRepositories
	userRepo  repositories.UserRepositories
	config    *configs.Config
}

// SendEmail ใช้ส่งอีเมล
func (e *eventService) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.config.EMAIL_FROM)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.config.SMTP_HOST, e.config.SMTP_PORT, e.config.EMAIL_FROM, e.config.EMAIL_PASSWORD)

	// ส่งอีเมล
	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	return nil
}

// CreateEvent implements EventUseCase.
func (e *eventService) CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	// สร้าง Event
	event, err := e.eventRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	// เช็คว่า Event_Email เป็น true หรือไม่
	if event.Notification {
		// ดึงข้อมูลผู้ใช้จาก UserRepository
		user, err := e.userRepo.GetUserByUserID(ctx, &requests.SendEmailRequest{
			ID: req.UserId, // ใช้ UserId เพื่อดึงข้อมูล
		})
		if err != nil {
			return nil, err
		}

		// ตรวจสอบว่ามีอีเมลของผู้ใช้
		if user.Email != "" {
			subject := fmt.Sprintf("Arom Notification: %s", event.Title)

			// จัดรูปแบบวันที่
			startFormatted := event.Start.Format("2 Jan 2006, 15:04:05")
			endFormatted := event.End.Format("2 Jan 2006, 15:04:05")

			body := fmt.Sprintf(`
				<html>
				<head><title>Arom Event Notification</title></head>
				<body>
					<h1>%s</h1>
					<p>%s</p>
					<p><strong>Event starts at:</strong> %s</p>
					<p><strong>Event ends at:</strong> %s</p>
				</body>
				</html>
			`, event.Title, event.Description, startFormatted, endFormatted)

			// ส่งอีเมลไปยังผู้ใช้
			err := e.SendEmail(user.Email, subject, body)
			if err != nil {
				return nil, err
			}
		}
	}

	return event, nil
}

// GetByUserIDEvent implements EventUseCase.
func (e *eventService) GetByUserIDEvent(ctx context.Context, id string) ([]*models.Event, error) {
	return e.eventRepo.GetByUserID(ctx, id)
}

// GetAllEvent implements EventUseCase.
func (e *eventService) GetAllEvent(ctx context.Context) ([]*models.Event, error) {
	return e.eventRepo.GetAll(ctx)
}

// GetByIDEvent implements EventUseCase.
func (e *eventService) GetByIDEvent(ctx context.Context, id string) (*models.Event, error) {
	return e.eventRepo.GetByID(ctx, id)
}

// UpdateDateEvent implements EventUseCase.
func (e *eventService) UpdateDateEvent(ctx context.Context, req *requests.UpdateEventRequest, id string) error {
	return e.eventRepo.Updatestatus(ctx, req, id)
}

func ProvideEventService(eventRepo repositories.EventRepositories, userRepo repositories.UserRepositories, config *configs.Config) EventUseCase {
	return &eventService{
		eventRepo: eventRepo,
		userRepo:  userRepo,
		config:    config,
	}
}
