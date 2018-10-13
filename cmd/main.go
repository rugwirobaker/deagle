package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/deagle/api"
	"github.com/rugwirobaker/deagle/app"
)

var (
	port    = os.Getenv("PORT")
	service = os.Getenv("SERVICE_NAME")
	version = os.Getenv("SERVICE_VERSION")
)

func main() {
	r := mux.NewRouter()
	a := api.Init(r)
	app := app.NewApp(port, service, version)

	app.HTTPServe(a.BaseRoutes.Root)
}
