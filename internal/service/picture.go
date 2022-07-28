package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/chiefcake/apod/internal/model"
	"github.com/chiefcake/apod/internal/storage/postgres"
)

const layout = "2006-01-02"

var ErrInvalidDate = errors.New("provided date cannot exceed the current date")

// Proxy contains service logic for retrieveing pictures from the APOD API or database.
type Picture struct {
	client     APODClient
	repository PictureRepository
}

// NewPicture is a constructor for Picture.
func NewPicture(client APODClient, repository PictureRepository) *Picture {
	return &Picture{
		client:     client,
		repository: repository,
	}
}

// GetByDate retrieves picture from the APOD API or database by a specific date.
func (a Picture) GetByDate(ctx context.Context, date string) (model.Picture, error) {
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		return model.Picture{}, errors.Wrap(err, "could not parse provided date")
	}

	now := time.Now()

	if dateTime.After(now) {
		return model.Picture{}, ErrInvalidDate
	}

	picture, err := a.repository.GetByDate(ctx, dateTime)
	if err != nil && !errors.Is(err, postgres.ErrNoPicture) {
		return model.Picture{}, errors.Wrap(err, "could not get picture by provided date")
	}

	if errors.Is(err, postgres.ErrNoPicture) {
		apod, err := a.client.GetPictureOfTheDay(ctx, dateTime.Format(layout))
		if err != nil {
			return model.Picture{}, errors.Wrap(err, "could not get picture of the day from the api")
		}

		picture = model.Picture{
			ID:             uuid.New(),
			Copyright:      apod.Copyright,
			Date:           dateTime,
			Explanation:    apod.Explanation,
			HDURL:          apod.HDURL,
			MediaType:      apod.MediaType,
			ServiceVersion: apod.ServiceVersion,
			Title:          apod.Title,
			URL:            apod.URL,
		}

		err = a.repository.Create(ctx, picture)
		if err != nil {
			return model.Picture{}, errors.Wrap(err, "could not save picture of the day")
		}
	}

	return picture, nil
}

// List retrieves all stored pictures from a database.
func (a Picture) List(ctx context.Context) ([]model.Picture, error) {
	pictures, err := a.repository.List(ctx)
	if err != nil {
		return []model.Picture{}, errors.Wrap(err, "could not get list of pictures")
	}

	return pictures, nil
}
