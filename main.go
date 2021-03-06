package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/exercise/db"
	"github.com/exercise/router"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
)

func main() {
	mongoDB := mongoDBFactory()

	r := router.New(mongoDB)

	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(r)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Info("running service at port 8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(c, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c

	var wait time.Duration
	wait = time.Second * 60
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func mongoDBFactory() *db.MongoDB {
	var dbName string
	var uri string

	if os.Getenv("MONGO_DATABASE") != "" {
		dbName = os.Getenv("MONGO_DATABASE")
	} else {
		dbName = "test"
	}

	if os.Getenv("MONGO_URI") != "" {
		uri = os.Getenv("MONGO_URI")
	} else {
		uri = "mongodb://localhost:27017"
	}

	log.Infof("connected %s %s", dbName, uri)
	return db.NewMongoDB(uri, dbName)
}
