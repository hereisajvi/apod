package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/chiefcake/apod/internal/model"
	"github.com/chiefcake/apod/internal/storage"
)

var ErrNoPicture = errors.New("no picture")

// PictureRepository contains all methods for managing pictures in a database.
type PictureRepository struct {
	postgres *storage.Postgres
}

// NewPictureRepository is a constructor for PictureRepository.
func NewPictureRepository(postgres *storage.Postgres) *PictureRepository {
	return &PictureRepository{
		postgres: postgres,
	}
}

// Create inserts provided picture to the pictures table.
func (r PictureRepository) Create(ctx context.Context, picture model.Picture) error {
	query := `INSERT INTO pictures (id, copyright, date, explanation, hdurl, media_type, service_version, title, url)
	VALUES (:id, :copyright, :date, :explanation, :hdurl, :media_type, :service_version, :title, :url)`

	_, err := r.postgres.NamedExecContext(ctx, query, picture)
	if err != nil {
		return errors.Wrap(err, "could not insert picture")
	}

	return nil
}

// GetByDate selects picture from the pictures table by provided date.
func (r PictureRepository) GetByDate(ctx context.Context, dateTime time.Time) (model.Picture, error) {
	query := `SELECT * FROM pictures WHERE date = $1`

	var picture model.Picture

	err := r.postgres.GetContext(ctx, &picture, query, dateTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Picture{}, ErrNoPicture
		}

		return model.Picture{}, errors.Wrap(err, "could not select picture by date")
	}

	return picture, nil
}

// List selects all pictures from the pictures table.
func (r PictureRepository) List(ctx context.Context) ([]model.Picture, error) {
	query := `SELECT * FROM pictures`

	var pictures []model.Picture

	err := r.postgres.SelectContext(ctx, &pictures, query)
	if err != nil {
		return nil, errors.Wrap(err, "could not select all pictures")
	}

	return pictures, nil
}
