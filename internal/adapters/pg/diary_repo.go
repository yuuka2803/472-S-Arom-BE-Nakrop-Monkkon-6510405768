package pg

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/lib/pq"
)

type DiaryPGRepository struct {
	db *sqlx.DB
}

// Create implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) Create(ctx context.Context, req *requests.CreateDiaryRequest) (*models.Diary, error) {
	var diary models.Diary
	emotionsArray := "{" + strings.Join(req.Emotions, ",") + "}"
	imagesArray := "{" + strings.Join(req.Images, ",") + "}"
	var emotions pq.StringArray
	var images pq.StringArray
	err := d.db.QueryRowxContext(ctx, `INSERT INTO "DIARY" (
    "Diary_Date",
    "Diary_Emotions",
    "Diary_Mood",
    "Diary_Description",
	"Diary_Images",
    "User_Id"
)
VALUES ($1, $2, $3, $4, $5 ,$6)
RETURNING
    "Diary_Id",
    "Diary_Date",
    "Diary_Emotions",
    "Diary_Mood",
    "Diary_Description",
	"Diary_Images",
    "User_Id";
`,
		req.Date, emotionsArray, req.Mood, req.Description, imagesArray, req.UserID).Scan(&diary.Id, &diary.Date, &emotions, &diary.Mood, &diary.Description, &images, &diary.UserID)
	if err != nil {
		return nil, err
	}
	diary.Emotions = emotions
	diary.Images = images

	return &diary, nil

}
func (d *DiaryPGRepository) GetAll(ctx context.Context) ([]*models.Diary, error) {
	var diaries []*models.Diary
	err := d.db.SelectContext(ctx, &diaries, `SELECT * FROM "DIARY"`)
	if err != nil {
		return nil, err
	}
	return diaries, nil
}

// GetByDate implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) GetByDate(ctx context.Context, date string) (*models.Diary, error) {
	var diary models.Diary
	err := d.db.GetContext(ctx, &diary, `SELECT * FROM "DIARY" WHERE "Diary_Date" = $1`, date)
	if err != nil {
		return nil, err
	}
	return &diary, nil
}

// GetByID implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) GetByID(ctx context.Context, id string) (*models.Diary, error) {
	var diary models.Diary
	err := d.db.GetContext(ctx, &diary, `SELECT * FROM "DIARY" WHERE "Diary_Id" = $1`, id)
	if err != nil {
		return nil, err
	}
	return &diary, nil
}

// GetByUserID implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) GetByUserID(ctx context.Context, userID string) ([]*models.Diary, error) {
	var diaries []*models.Diary
	err := d.db.SelectContext(ctx, &diaries, `SELECT "Diary_Id", "Diary_Date", "Diary_Emotions", "Diary_Mood", "Diary_Description","Diary_Images" ,"User_Id" FROM "DIARY" WHERE "User_Id" = $1`, userID)
	if err != nil {
		return nil, err
	}
	return diaries, nil
}

// Update implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) Update(ctx context.Context, req *requests.UpdateDiaryRequest, date string) error {
	emotionsArrayU := "{" + strings.Join(req.Emotions, ",") + "}"
	imagesArray := "{" + strings.Join(req.Images, ",") + "}"
	_, err := d.db.ExecContext(ctx, `
    UPDATE "DIARY"
    SET "Diary_Emotions" = $1,
        "Diary_Mood" = $2,
        "Diary_Description" = $3,
        "Diary_Images" = $4
    WHERE "Diary_Date" = $5
    AND "User_Id" = $6`,
		emotionsArrayU, req.Mood, req.Description, imagesArray, date, req.UserID)

	if err != nil {
		return err
	}
	return nil
}

func NewDiaryPGRepository(db *sqlx.DB) repositories.DiaryRepositories {
	return &DiaryPGRepository{
		db: db,
	}
}
