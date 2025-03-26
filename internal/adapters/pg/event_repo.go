package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type EventPGRepository struct {
	db *sqlx.DB
}

// UpdateDate implements repositories.EventRepositories.
func (e *EventPGRepository) Update(ctx context.Context, req *requests.UpdateEventRequest, id string) (*models.Event, error) {
	var event models.Event

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, fmt.Errorf("failed to load Bangkok timezone: %v", err)
	}

	// แปลงเวลาจาก string -> time.Time
	startTimeLocal, err := time.ParseInLocation("2006-01-02T15:04:05Z", req.Start, location)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start time: %v", err)
	}

	// แปลงเป็น UTC ก่อนเก็บ
	startTimeUTC := startTimeLocal.UTC()

	// แปลงเวลาสิ้นสุด
	endTimeLocal, err := time.ParseInLocation("2006-01-02T15:04:05Z", req.End, location)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end time: %v", err)
	}

	// แปลงเป็น UTC ก่อนเก็บ
	endTimeUTC := endTimeLocal.UTC()
	
	err = e.db.QueryRowContext(ctx, `UPDATE "EVENT" 
		SET "Event_Title" = $1, "Event_Start" = $2, "Event_End" = $3, 
			"Event_Reminder" = $4, "Event_Description" = $5  
		WHERE "Event_Id" = $6 
		RETURNING 
			"Event_Id", "Event_Title", "Event_Description", 
			"Event_Start", "Event_End", "Event_Type", "Event_Complete", 
			"Event_Tag", "Event_Email", "Event_Reminder", "User_Id"
	`, req.Title, startTimeUTC, endTimeUTC, req.Reminder, req.Description, id).
		Scan(
			&event.Id, &event.Title, &event.Description, &event.Start, &event.End,
			&event.Type, &event.Completed, &event.Tag, &event.Notification, &event.Reminder, &event.UserId,
		)
	
	if err != nil {
		return nil, fmt.Errorf("event with id %s not found", id)
	}

	return &event, nil
}


// Updatestatus implements repositories.EventRepositories.
func (e *EventPGRepository) UpdateStatus(ctx context.Context, req *requests.UpdateStatusEventRequest, id string) error {
	_, err := e.db.ExecContext(ctx, `UPDATE "EVENT" SET "Event_Complete" = $1 WHERE "Event_Id" = $2`, req.Completed, id)
	if err != nil {
		return err
	}
	return nil
}

// GetByUserID implements repositories.EventRepositories.
func (e *EventPGRepository) GetByUserID(ctx context.Context, id string) ([]*models.Event, error) {
	var events []*models.Event
	err := e.db.SelectContext(ctx, &events, `SELECT * FROM "EVENT" WHERE "User_Id" = $1`, id)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetAll implements repositories.EventRepositories.
func (e *EventPGRepository) GetAll(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	err := e.db.SelectContext(ctx, &events, `SELECT * FROM "EVENT"`)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetByID implements repositories.EventRepositories.
func (e *EventPGRepository) GetByID(ctx context.Context, id string) (*models.Event, error) {
	var event models.Event
	
	err := e.db.QueryRowContext(ctx, `SELECT 
	"Event_Id", "Event_Title", "Event_Description", 
	"Event_Start", "Event_End", "Event_Type", 
	"Event_Complete", "Event_Tag", "Event_Email", 
	"Event_Reminder", "User_Id"
	FROM "EVENT" WHERE "Event_Id" = $1`, id).
	Scan(
		&event.Id, &event.Title, &event.Description, &event.Start, &event.End,
		&event.Type, &event.Completed, &event.Tag, &event.Notification, &event.Reminder, &event.UserId,
	)
	
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func NewEventPGRepository(db *sqlx.DB) repositories.EventRepositories {
	return &EventPGRepository{
		db: db,
	}
}

func (e *EventPGRepository) Create(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	var event models.Event

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, fmt.Errorf("failed to load Bangkok timezone: %v", err)
	}

	// แปลงเวลาจาก string -> time.Time
	startTimeLocal, err := time.ParseInLocation("2006-01-02T15:04:05Z", req.Start, location)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start time: %v", err)
	}

	// แปลงเป็น UTC ก่อนเก็บ
	startTimeUTC := startTimeLocal.UTC()

	// แปลงเวลาสิ้นสุด
	endTimeLocal, err := time.ParseInLocation("2006-01-02T15:04:05Z", req.End, location)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end time: %v", err)
	}

	// แปลงเป็น UTC ก่อนเก็บ
	endTimeUTC := endTimeLocal.UTC()

	// Insert ลง Database
	err = e.db.QueryRowxContext(ctx, `INSERT INTO "EVENT" (
		"Event_Title", 
		"Event_Description", 
		"Event_Start", 
		"Event_End", 
		"Event_Tag",
		"Event_Email",
		"Event_Reminder",
		"User_Id"
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING
		"Event_Id",
		"Event_Title",
		"Event_Description",
		"Event_Start",
		"Event_End",
		"Event_Type",
		"Event_Complete",
		"Event_Tag",
		"Event_Email",
		"Event_Reminder",
		"User_Id"`,
		req.Title, req.Description, startTimeUTC, endTimeUTC, req.Tag, req.Notification, req.Reminder, req.UserId,
	).Scan(
		&event.Id, &event.Title, &event.Description, &event.Start, &event.End,
		&event.Tag, &event.Completed, &event.Type, &event.Notification, &event.Reminder, &event.UserId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert event: %v", err)
	}

	return &event, nil
}

