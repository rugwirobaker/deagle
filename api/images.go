package api

import "net/http"

//ImageHander ...
type ImageHander struct{}

//InitImages ...
func (api *API) InitImages() {}

//Create ...
func (h *ImageHander) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}

//Update ...
func (h *ImageHander) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}

//Retrieve ...
func (h *ImageHander) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}

//Delete ...
func (h *ImageHander) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}
