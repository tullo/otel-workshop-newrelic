module github.com/tullo/otel-workshop-newrelic

go 1.22

// replace github.com/tullo/otel-workshop/web/fib => ../workshop/web/fib

require (
	github.com/newrelic/go-agent/v3 v3.34.0
	github.com/tullo/otel-workshop/web/fib v1.0.4
	go.opentelemetry.io/otel v1.30.0
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240711142825-46eb208f015d // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
