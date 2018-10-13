package api

import (
	"github.com/gorilla/mux"
)

//API ...
type API struct {
	BaseRoutes *Routes
}

//Routes defines are REST endpoints
type Routes struct {
	Root   *mux.Router //api/
	Users  *mux.Router //api/users
	Images *mux.Router //api/images
}

//Init initialises application API
func Init(root *mux.Router) *API {
	api := &API{
		BaseRoutes: &Routes{},
	}
	api.BaseRoutes.Root = root.PathPrefix("/api").Subrouter()
	api.BaseRoutes.Users = api.BaseRoutes.Root.PathPrefix("/users").Subrouter()
	api.BaseRoutes.Images = api.BaseRoutes.Root.PathPrefix("/images").Subrouter()

	api.InitUsers()
	api.InitImages()

	return api
}
