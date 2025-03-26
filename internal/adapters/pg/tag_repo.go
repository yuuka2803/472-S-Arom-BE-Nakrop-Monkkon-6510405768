package pg

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type TagPGRepository struct {
	db *sqlx.DB
}

// Create implements repositories.TagRepositories.
func (t *TagPGRepository) Create(ctx context.Context, req *requests.CreateTagRequest) (*models.Tag, error) {
	var tag models.Tag
	err := t.db.QueryRowxContext(ctx, `
		INSERT INTO "TAG" (
			"Tag_Name", 
			"User_Id"
		) 
		VALUES ($1, $2) 
		RETURNING 
			"Tag_ID", 
			"Tag_Name", 
			"User_Id"
	`, 
		req.Name, req.UserID).Scan(&tag.Id, &tag.Name, &tag.UserId)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}


// GetByID implements repositories.TagRepositories.
func (t *TagPGRepository) GetByID(ctx context.Context, id string) (*models.Tag, error) {
	var tag models.Tag
	err := t.db.GetContext(ctx, &tag, `SELECT * FROM "TAG" WHERE "Tag_ID" = $1`, id)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetByUserID implements repositories.TagRepositories.
func (t *TagPGRepository) GetByUserID(ctx context.Context, id string) ([]*models.Tag, error) {
	var tags []*models.Tag
	println(id)
	err := t.db.SelectContext(ctx, &tags, `SELECT * FROM "TAG" WHERE "User_Id" = $1`, id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *TagPGRepository) Delete(ctx context.Context, id string) error {
	_, err := t.db.ExecContext(ctx, `DELETE FROM "TAG" WHERE "Tag_ID" = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TagPGRepository) Update(ctx context.Context, req *requests.UpdateTagRequest, id string) error {
	_, err := t.db.ExecContext(ctx, `UPDATE "TAG" SET "Tag_Name" = $1 WHERE "Tag_ID" = $2`, req.Name,id)
	if err != nil {
		return err
	}
	return nil
}

func NewTagPGRepository(db *sqlx.DB) repositories.TagRepositories {
	return &TagPGRepository{
		db: db,
	}
}