package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/seagullbird/headr-apigateway/config"
	contentmgrendpoint "github.com/seagullbird/headr-contentmgr/endpoint"
	contentmgrtransport "github.com/seagullbird/headr-contentmgr/transport"
	sitemgrendpoint "github.com/seagullbird/headr-sitemgr/endpoint"
	sitemgrtransport "github.com/seagullbird/headr-sitemgr/transport"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func initGRPCConnection(svcname string, logger log.Logger) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:2018", svcname), grpc.WithInsecure())
	if err != nil {
		logger.Log(fmt.Sprintf("Error connecting %s", svcname), err.Error())
		return nil, err
	}
	return conn, nil
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
	isReady := true
	// Sitemgr
	{
		sitemgrconn, sitemgrerr := initGRPCConnection("sitemgr", logger)
		if sitemgrerr != nil {
			isReady = false
		}
		service := sitemgrtransport.NewGRPCClient(sitemgrconn, logger)
		endpoints := sitemgrendpoint.New(service, logger)
		r.PathPrefix("/sitemgr").Handler(http.StripPrefix("/sitemgr", sitemgrtransport.NewHTTPHandler(endpoints, logger)))
	}

	// Contentmgr
	{
		contentmgrconn, contentmgrerr := initGRPCConnection("contentmgr", logger)
		if contentmgrerr != nil {
			isReady = false
		}
		service := contentmgrtransport.NewGRPCClient(contentmgrconn, logger)
		endpoints := contentmgrendpoint.New(service, logger)
		r.PathPrefix("/contentmgr").Handler(http.StripPrefix("/contentmgr", contentmgrtransport.NewHTTPHandler(endpoints, logger)))
	}

	// liveness probe
	r.Methods("GET").Path("/healthz").Handler(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	// readiness probe
	r.Methods("GET").Path("/readyz").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isReady {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))

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
		errc <- http.ListenAndServe(config.PORT, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}),
			handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Language", "Origin", "Authorization", "Content-Type"}))(r))
	}()

	logger.Log("exit", <-errc)
}
