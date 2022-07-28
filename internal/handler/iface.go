package handler

import (
	"context"

	"github.com/chiefcake/apod/internal/model"
)

//go:generate mockgen -source=iface.go -destination=../mock/services.go -package=mock
type PictureService interface {
	GetByDate(ctx context.Context, date string) (model.Picture, error)
	List(ctx context.Context) ([]model.Picture, error)
}
