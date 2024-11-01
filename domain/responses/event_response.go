package responses

import "github.com/jackc/pgx/v5/pgtype"

type Event struct {
	Id          pgtype.UUID `json:"id" db:"id"`
	Name        pgtype.Text `json:"name" db:"name"`
	Description pgtype.Text `json:"description" db:"description"`
}
