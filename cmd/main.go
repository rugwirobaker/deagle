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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World !"))
}
func main() {
	router := mux.NewRouter()
	a := api.Init(router)
	server := http.Server{
		Addr:         port,
		Handler:      a.,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(server.ListenAndServe())
}
