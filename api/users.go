package api

import (
	"net/http"

	"github.com/rugwirobaker/deagle/utils"
)

//UserHandler ...
type UserHandler struct{}

//InitUsers ...
func (api *API) InitUsers() {
	h := UserHandler{}
	api.BaseRoutes.Users.Handle("/create", h.Create())
	api.BaseRoutes.Users.Handle("/update", h.Update())
	api.BaseRoutes.Users.Handle("/retrieve", h.Retrieve())
	api.BaseRoutes.Users.Handle("/delete", h.Delete())
}

//Create ...
func (h *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is a creation!\n",
		)
	}
}

//Update ...
func (h *UserHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is an update!\n",
		)
	}
}

//Retrieve ...
func (h *UserHandler) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is a retrieval!\n",
		)
	}
}

//Delete ...
func (h *UserHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccessResult(
			w,
			"Hello World, this is annilation!\n",
		)
	}
}
