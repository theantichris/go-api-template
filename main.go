package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/theantichris/go-api-template/handlers"
	"golang.org/x/net/context"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("SENTRY_DSN")
	if dsn != "" {
		loadSentry(dsn)
		defer sentry.Flush(time.Second * 2)
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.HealthCheck).Methods(http.MethodGet)

	server := &http.Server{
		Addr:         "0.0.0.0:" + os.Getenv("PORT"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go serve(server)

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	server.Shutdown(ctx)
	log.Infoln("shutting down")
	os.Exit(0)
}

func loadSentry(dsn string) {
	if err := sentry.Init(sentry.ClientOptions{Dsn: dsn}); err != nil {
		log.Errorln("failed to connect to Sentry: ", err)
	} else {
		log.Infoln("connected to Sentry")
	}
}

func serve(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		sentry.CaptureMessage("err")
		log.Errorln("failed to start the server: ", err)
	}
}
