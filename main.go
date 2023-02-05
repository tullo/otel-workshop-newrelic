package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	app, err := newrelic.NewApplication(newrelic.ConfigFromEnvironment())
	if err != nil {
		l.Fatal(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start web server.
	go func() {
		errCh <- ServeNewRelic(context.Background(), app)
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
