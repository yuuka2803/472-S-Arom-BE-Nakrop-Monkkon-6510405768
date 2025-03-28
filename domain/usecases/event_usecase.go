package usecases

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/internal/infrastrutures/mailer"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error)
	GetAllEvent(ctx context.Context) ([]*models.Event, error)
	GetByIDEvent(ctx context.Context, id string) (*models.Event, error)
	GetByUserIDEvent(ctx context.Context, id string) ([]*models.Event, error)
	UpdateEvent(ctx context.Context, req *requests.UpdateEventRequest, id string) (*models.Event, error)
	UpdateStatusEvent(ctx context.Context, req *requests.UpdateStatusEventRequest, id string) error
}

type eventService struct {
	eventRepo repositories.EventRepositories
	userRepo  repositories.UserRepositories
	mailer    mailer.Mailer
	config    *configs.Config
}

// UpdateEvent implements EventUseCase.
func (e *eventService) UpdateEvent(ctx context.Context, req *requests.UpdateEventRequest, id string) (*models.Event, error) {
	// อัปเดต Event และใช้ค่าใหม่
	updatedEvent, err := e.eventRepo.Update(ctx, req, id)
	if err != nil {
		return nil, err
	}

	// ดึงข้อมูลผู้ใช้จาก UserRepository
	user, err := e.userRepo.GetUserByUserID(ctx, &requests.SendEmailRequest{
		ID: req.UserId, // ใช้ UserId เพื่อดึงข้อมูล
	})
	if err != nil {
		return nil, err
	}

	// เช็คว่า Event_Email เป็น true และ User มี Email
	if updatedEvent.Notification && user.Email != "" {
		subject := fmt.Sprintf("Update Notification: %s", updatedEvent.Title)

		location, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			log.Fatalf("Error loading time location: %v", err)
		}

		// แปลงเวลาให้เป็นโซนที่กำหนดก่อน Format
		startFormatted := updatedEvent.Start.In(location).Format("2 Jan 2006, 15:04:05")
		endFormatted := updatedEvent.End.In(location).Format("2 Jan 2006, 15:04:05")

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
		`, updatedEvent.Title, updatedEvent.Description, startFormatted, endFormatted)

		// ส่งอีเมลไปยังผู้ใช้
		if err := e.mailer.SendEmail(user.Email, subject, body); err != nil {
			return nil, err
		}
	}

	// ตรวจสอบว่า updatedEvent.Reminder ไม่ใช่ "None"
	if updatedEvent.Reminder != "None" {
		log.Println("Reminder function called with reminderAt:", updatedEvent.Reminder)
		// เรียกใช้ฟังก์ชัน SendReminderEmailBeforeEvent
		err := e.SendReminderEmailBeforeEvent(updatedEvent, user, updatedEvent.Reminder)
		if err != nil {
			return nil, err
		}
	}

	return updatedEvent, nil // คืนค่า updatedEvent แทน nil
}

// SendReminderEmailBeforeEvent ใช้ส่งอีเมลก่อนเวลาเริ่ม Event
func (e *eventService) SendReminderEmailBeforeEvent(event *models.Event, user *models.User, durationStr string) error {
	// ถ้า durationStr เป็น "None" ให้ return nil
	if durationStr == "None" {
		log.Println("Reminder is set to none. Skipping email.")
		return nil
	}

	// แปลง durationStr เป็น time.Duration
	durationBeforeEvent, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Printf("Error parsing duration: %v", err)
		return err
	}

	// โหลดโซนเวลาของกรุงเทพฯ
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatalf("Error loading time location: %v", err)
	}

	// แปลง event.Start เป็นเวลาตามโซนที่ต้องการ
	eventStartTime := event.Start.In(location)

	// แปลงเวลาปัจจุบันให้เป็นโซนเวลาที่ต้องการ
	currentTime := time.Now().In(location)

	// หาช่วงเวลาที่ต้องการส่งก่อน Event
	timeUntilEventStart := time.Until(eventStartTime) // คำนวณระยะเวลาจากเวลาปัจจุบันถึงเวลาของ Event
	timeToWait := timeUntilEventStart - durationBeforeEvent

	// ตั้งเวลาให้ส่งอีเมลตามระยะเวลาที่กำหนด
	timer := time.NewTimer(timeToWait)

	log.Printf("Reminder duration: %v, Time until event start: %v, Time to wait: %v",
		timeUntilEventStart, timeToWait, timeToWait)
	log.Println("Current time:", currentTime)
	log.Println("Event start time:", eventStartTime)

	go func() {
		<-timer.C // รอจนกว่าจะถึงเวลา
		// ส่งอีเมล
		subject := fmt.Sprintf("Reminder: %s", event.Title)

		// จัดรูปแบบวันที่
		startFormatted := event.Start.In(location).Format("2 Jan 2006, 15:04:05")
		endFormatted := event.End.In(location).Format("2 Jan 2006, 15:04:05")

		body := fmt.Sprintf(`
			<html>
			<head><title>Arom Event Reminder</title></head>
			<body>
				<h1>%s</h1>
				<p>%s</p>
				<p><strong>Event starts at:</strong> %s</p>
				<p><strong>Event ends at:</strong> %s</p>
			</body>
			</html>
		`, event.Title, event.Description, startFormatted, endFormatted)

		// ส่งอีเมลไปยังผู้ใช้
		err := e.mailer.SendEmail(user.Email, subject, body)
		if err != nil {
			log.Printf("Error sending reminder email: %v", err)
		}
	}()

	return nil
}

// CreateEvent implements EventUseCase.
func (e *eventService) CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	// สร้าง Event
	event, err := e.eventRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	// ดึงข้อมูลผู้ใช้จาก UserRepository
	user, err := e.userRepo.GetUserByUserID(ctx, &requests.SendEmailRequest{
		ID: req.UserId, // ใช้ UserId เพื่อดึงข้อมูล
	})
	if err != nil {
		return nil, err
	}

	// เช็คว่า Event_Email เป็น true หรือไม่
	if event.Notification && user.Email != "" {
		subject := fmt.Sprintf("Arom Notification: %s", event.Title)

		location, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			log.Fatalf("Error loading time location: %v", err)
		}

		// แปลงเวลาให้เป็นโซนที่กำหนดก่อน Format
		startFormatted := event.Start.In(location).Format("2 Jan 2006, 15:04:05")
		endFormatted := event.End.In(location).Format("2 Jan 2006, 15:04:05")

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
		err = e.mailer.SendEmail(user.Email, subject, body)
		if err != nil {
			return nil, err
		}
	}

	// ตรวจสอบว่า event.Reminder ไม่ใช่ "None"
	if event.Reminder != "None" {
		log.Println("Reminder function called with reminderAt:", event.Reminder)
		// เรียกใช้ฟังก์ชัน SendReminderEmailBeforeEvent
		err := e.SendReminderEmailBeforeEvent(event, user, event.Reminder)
		if err != nil {
			return nil, err
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

// UpdateStatusDateEvent implements EventUseCase.
func (e *eventService) UpdateStatusEvent(ctx context.Context, req *requests.UpdateStatusEventRequest, id string) error {
	return e.eventRepo.UpdateStatus(ctx, req, id)
}

func ProvideEventService(eventRepo repositories.EventRepositories, userRepo repositories.UserRepositories, config *configs.Config, mailer mailer.Mailer) EventUseCase {
	return &eventService{
		eventRepo: eventRepo,
		userRepo:  userRepo,
		config:    config,
		mailer:    mailer,
	}
}
