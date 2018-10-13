package api

import (
	"net/http"

	"github.com/rugwirobaker/deagle/utils"
)

//ImageHandler ...
type ImageHandler struct{}

//InitImages ...
func (api *API) InitImages() {
	h := ImageHandler{}
	api.BaseRoutes.Images.Handle("/create", h.Create())
	api.BaseRoutes.Images.Handle("/update", h.Update())
	api.BaseRoutes.Images.Handle("/retrieve", h.Retrieve())
	api.BaseRoutes.Images.Handle("/delete", h.Delete())
}

//Create ...
func (h *ImageHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is a creation!\n",
		)
	}
}

//Update ...
func (h *ImageHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is an update!\n",
		)
	}
}

//Retrieve ...
func (h *ImageHandler) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is a retrieval!\n",
		)
	}
}

//Delete ...
func (h *ImageHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is annilation!\n",
		)
	}
}
