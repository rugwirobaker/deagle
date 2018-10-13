package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rugwirobaker/deagle/api"
)

var (
	shutdownTimeout = flag.Duration("shutdown-timeout", 10*time.Second,
		"shutdown timeout (5s,5m,5h) before connections are cancelled")
)

//App ....
type App struct {
	Service string
	PORT    string
	Version string
	API     *api.API
	Server  *http.Server
}

//NewApp ...
func NewApp(port, src, vrs string) *App {
	app := &App{
		Service: src,
		Version: vrs,
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  time.Second * 60,
		},
	}
	return app
}

//HTTPServe ...
func (a *App) HTTPServe(h http.Handler) {
	a.Server.Handler = h
	//log.Fatal(a.Server.ListenAndServe())

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("%s listening on 0.0.0.0:%s with %v timeout",
			a.Service, a.PORT, *shutdownTimeout,
		)
		if err := a.Server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stop

	log.Printf("%s shutting down ...\n", a.Service)

	ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s down\n", a.Service)

}
