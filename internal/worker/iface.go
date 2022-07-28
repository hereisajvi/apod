package worker

import (
	"context"

	"github.com/chiefcake/apod/internal/model"
)

type PictureService interface {
	GetByDate(ctx context.Context, date string) (model.Picture, error)
}
