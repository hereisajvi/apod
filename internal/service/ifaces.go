package service

import (
	"context"
	"time"

	"github.com/chiefcake/apod/clients/apod"
	"github.com/chiefcake/apod/internal/model"
)

type PictureRepository interface {
	Create(ctx context.Context, picture model.Picture) error
	GetByDate(ctx context.Context, dateTime time.Time) (model.Picture, error)
	List(ctx context.Context) ([]model.Picture, error)
}

type APODClient interface {
	GetPictureOfTheDay(ctx context.Context, date string) (apod.Picture, error)
}
