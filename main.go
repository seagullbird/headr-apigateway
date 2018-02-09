package main

import (
	"github.com/go-kit/kit/log"
	"os"
	sitemgrtransport "github.com/seagullbird/headr-sitemgr/transport"
	sitemgrendpoint "github.com/seagullbird/headr-sitemgr/endpoint"
	"google.golang.org/grpc"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os/signal"
	"syscall"
	"github.com/seagullbird/headr-apigateway/config"
)

func main() {
	// Logging
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Transport
	r := mux.NewRouter()

	// Sitemgr
	conn, err := grpc.Dial("sitemgr:2018", grpc.WithInsecure())
	if err != nil {
		logger.Log("Error connecting sitemgr", err.Error())
	}
	service := sitemgrtransport.NewGRPCClient(conn, logger)
	newsiteEndpoint := sitemgrendpoint.MakeNewSiteEndpoint(service)

	endpoints := sitemgrendpoint.Set{
		NewSiteEndpoint: newsiteEndpoint,
	}

	// Routes
	// for test
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	r.PathPrefix("/sitemgr").Handler(http.StripPrefix("/sitemgr", sitemgrtransport.NewHTTPHandler(endpoints, logger)))

	// Interrupt handler.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport.
	go func() {
		logger.Log("transport", "HTTP", "port", config.PORT)
		errc <- http.ListenAndServe(config.PORT, r)
	}()

	logger.Log("exit", <-errc)
}
