package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/deagle/api"
)

var port = os.Getenv("PORT")

func main() {
	r := mux.NewRouter()
	a := api.Init(r)
	Server := &http.Server{
		Addr:         port,
		Handler:      a.BaseRoutes.Root,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	log.Fatal(Server.ListenAndServe())
}
