package responses

import "github.com/jackc/pgx/v5/pgtype"

type Tag struct {
	Id          pgtype.UUID `json:"id" db:"id"`
	Name        pgtype.Text `json:"name" db:"name"`
}
