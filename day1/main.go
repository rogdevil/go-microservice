package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rogdevil/handler"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers
	nph := handler.NewProducts(l)

	// custom server mux
	sm := http.NewServeMux()
	sm.Handle("/", nph)

	s := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("starting server at http://localhost:8080")

		err := s.ListenAndServe()

		if err != nil {
			l.Printf("there is an error %s \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c

	log.Print("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
