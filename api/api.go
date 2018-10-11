package api

import (
	"github.com/gorilla/mux"
)

//API Routes defines are REST endpoints
type API struct {
	Root   *mux.Router //api/
	Users  *mux.Router //api/users
	Images *mux.Router //api/images
}

//Init initialises application API
func Init(root *mux.Router) {}
