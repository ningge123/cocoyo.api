package models

import "time"

type Thread struct {
	BannedAt time.Time
	Cache []byte
	CreatedAt   time.Time
	DeletedAt   time.Time
	ExcellentAt time.Time
	FrozenAt    time.Time
	ID          int
	NodeID      int
	PinnedAt    time.Time
	PopularAt   time.Time
	PublishedAt time.Time
	Title       string
	UpdatedAt   time.Time
	UserID      int
}
