package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/tullo/otel-workshop/web/fib"
	"go.opentelemetry.io/otel"
)

const name = "otel-workshop"

func ServeNewRelic(ctx context.Context, app *newrelic.Application) error {
	_, span := otel.Tracer(name).Start(ctx, "ServeNewRelic")
	defer span.End()

	mux := http.NewServeMux()
	mux.Handle(newrelic.WrapHandle(app, "/", http.HandlerFunc(fib.RootHandler)))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle(newrelic.WrapHandle(app, "/fib", http.HandlerFunc(fib.FibHandler)))
	mux.Handle(newrelic.WrapHandle(app, "/fibinternal", http.HandlerFunc(fib.FibHandler)))

	fmt.Println("Your server is live!\nTry to navigate to: http://127.0.0.1:3000/fib?n=6")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return fmt.Errorf("could not start web server: %w", err)
	}

	return nil
}
