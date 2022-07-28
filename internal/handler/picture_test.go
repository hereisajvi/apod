package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/chiefcake/apod/internal/mock"
	"github.com/chiefcake/apod/internal/model"
)

func TestPictureGetByDate(t *testing.T) {
	type request struct {
		date string
	}

	type response struct {
		body   interface{}
		status int
		err    error
	}

	type behavior struct {
		fn func(service *mock.MockPictureService, request request, response response)
	}

	tests := map[string]struct {
		request  request
		behavior behavior
		response response
	}{
		"OK": {
			request: request{
				date: "2022-07-28",
			},
			behavior: behavior{
				fn: func(service *mock.MockPictureService, request request, response response) {
					service.EXPECT().GetByDate(gomock.Any(), request.date).Return(response.body, response.err)
				},
			},
			response: response{
				body: model.Picture{
					ID:             uuid.MustParse("ed2953ec-b716-4a8b-9a23-0e325cfe4376"),
					Copyright:      "Jeff Dai",
					Date:           time.Now(),
					Explanation:    "...",
					HDURL:          "https://apod.nasa.gov/apod/image/2207/AncientTreeNCP_Dai.jpg",
					MediaType:      "image",
					ServiceVersion: "v1",
					Title:          "North Celestial Tree",
					URL:            "https://apod.nasa.gov/apod/image/2207/AncientTreeNCP_Dai1024.jpg",
				},
				status: http.StatusOK,
				err:    nil,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockPictureService(ctrl)
	handler := NewPicture(service)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").
		Subrouter()
	v1Router := apiRouter.PathPrefix("/v1").
		Subrouter()

	v1Router.HandleFunc("/pictures/{date}", handler.GetByDate).
		Methods(http.MethodGet)

	server := httptest.NewServer(router)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.behavior.fn(service, test.request, test.response)

			expect.GET("/api/v1/pictures/{date}", test.request.date).
				Expect().
				Status(test.response.status).
				JSON().
				Equal(test.response.body)
		})
	}
}

func TestPictureList(t *testing.T) {
	type response struct {
		body   interface{}
		status int
		err    error
	}

	type behavior struct {
		fn func(service *mock.MockPictureService, response response)
	}

	tests := map[string]struct {
		response response
		behavior behavior
	}{
		"OK": {
			behavior: behavior{
				fn: func(service *mock.MockPictureService, response response) {
					service.EXPECT().List(gomock.Any()).Return(response.body, response.err)
				},
			},
			response: response{
				body: []model.Picture{
					{
						ID:             uuid.MustParse("ed2953ec-b716-4a8b-9a23-0e325cfe4376"),
						Copyright:      "Jeff Dai",
						Date:           time.Now(),
						Explanation:    "...",
						HDURL:          "https://apod.nasa.gov/apod/image/2207/AncientTreeNCP_Dai.jpg",
						MediaType:      "image",
						ServiceVersion: "v1",
						Title:          "North Celestial Tree",
						URL:            "https://apod.nasa.gov/apod/image/2207/AncientTreeNCP_Dai1024.jpg",
					},
				},
				status: http.StatusOK,
				err:    nil,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockPictureService(ctrl)
	handler := NewPicture(service)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").
		Subrouter()
	v1Router := apiRouter.PathPrefix("/v1").
		Subrouter()

	v1Router.HandleFunc("/pictures", handler.List).
		Methods(http.MethodGet)

	server := httptest.NewServer(router)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.behavior.fn(service, test.response)

			expect.GET("/api/v1/pictures").
				Expect().
				Status(test.response.status).
				JSON().
				Equal(test.response.body)
		})
	}
}
