package server

import "net/http"

type PictureHandler interface {
	GetByDate(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}
