package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chiefcake/apod/internal/service"
	"github.com/chiefcake/apod/internal/status"
)

type Picture struct {
	service PictureService
}

func NewPicture(service PictureService) *Picture {
	return &Picture{
		service: service,
	}
}

// GetByDate example
// @Summary Get an astronomy picture of the specific date
// @ID getByDate
// @Produce  json
// @Success 200 {object} model.Picture "OK"
// @Failure 400 {object} status.ErrorResponse "Bad request error"
// @Failure 500 {object} status.ErrorResponse "Internal error"
// @Router /api/v1/pictures/{date} [get]
// GetByDate retrieves picture by a specific date and returns it or error response.
func (a Picture) GetByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	encoder := json.NewEncoder(w)

	date, ok := vars["date"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		response := status.NewErrorResponse(http.StatusBadRequest, "date cannot be blank")

		err := encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}

	picture, err := a.service.GetByDate(r.Context(), date)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, service.ErrInvalidDate) {
			statusCode = http.StatusBadRequest
		}

		w.WriteHeader(statusCode)
		response := status.NewErrorResponse(statusCode, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}

	err = encoder.Encode(&picture)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := status.NewErrorResponse(http.StatusInternalServerError, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}
}

// List example
// @Summary Get all stored astronomy pictures
// @ID list
// @Produce  json
// @Success 200 {array} model.Picture "OK"
// @Failure 400 {object} status.ErrorResponse "Bad request error"
// @Failure 500 {object} status.ErrorResponse "Internal error"
// @Router /api/v1/pictures [get]
// List retrieves all stored pictures returns them or error response.
func (a Picture) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	pictures, err := a.service.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := status.NewErrorResponse(http.StatusInternalServerError, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}

	err = encoder.Encode(&pictures)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := status.NewErrorResponse(http.StatusInternalServerError, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}
}
