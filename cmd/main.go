package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var port = os.Getenv("PORT")

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World !"))
}
func main() {
	server := http.Server{
		Addr:         port,
		Handler:      http.HandlerFunc(handler),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(server.ListenAndServe())
}
