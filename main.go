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

func initGRPCConnection(svcname string, logger log.Logger) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%s:2018", svcname), grpc.WithInsecure())
	if err != nil {
		logger.Log(fmt.Sprintf("Error connecting %s", svcname), err.Error())
	}
	return conn
}

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
	{
		sitemgrconn := initGRPCConnection("sitemgr", logger)
		service := sitemgrtransport.NewGRPCClient(sitemgrconn, logger)
		newsiteEndpoint := sitemgrendpoint.MakeNewSiteEndpoint(service)
		endpoints := sitemgrendpoint.Set{
			NewSiteEndpoint: newsiteEndpoint,
		}
		r.PathPrefix("/sitemgr").Handler(http.StripPrefix("/sitemgr", sitemgrtransport.NewHTTPHandler(endpoints, logger)))
	}

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
