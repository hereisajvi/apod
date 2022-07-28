package model

import (
	"time"

	"github.com/google/uuid"
)

type Picture struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Copyright      string    `json:"copyright" db:"copyright"`
	Date           time.Time `json:"date" db:"date"`
	Explanation    string    `json:"explanation" db:"explanation"`
	HDURL          string    `json:"hdurl" db:"hdurl"`
	MediaType      string    `json:"media_type" db:"media_type"`
	ServiceVersion string    `json:"service_version" db:"service_version"`
	Title          string    `json:"title" db:"title"`
	URL            string    `json:"url" db:"url"`
}
