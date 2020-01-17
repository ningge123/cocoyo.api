package models

import (
	"database/sql"
	"time"
)

type Tag struct {
	Color     string
	CreatedAt time.Time
	DeletedAt time.Time
	Icon      sql.NullString
	ID        int
	Name      string
	Slug      string
	UpdatedAt time.Time
}