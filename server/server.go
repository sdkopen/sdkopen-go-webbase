package server

import (
	"log"

	"github.com/sdkopen/sdkopen-go-webbase/logging"
	"github.com/sdkopen/sdkopen-go-webbase/observer"
)

var (
	SrvRoutes []Route
	Srv       Server
)

type Server interface {
	Initialize()
	Shutdown() error
	InjectMiddlewares()
	InjectRoutes()
	ListenAndServe() error
}

func AddRoutes(routes []Route) {
	SrvRoutes = append(SrvRoutes, routes...)
}

func ListenAndServe(server func() Server) {
	Srv = server()
	Srv.Initialize()
	Srv.InjectMiddlewares()
	Srv.InjectRoutes()

	observer.Attach(restObserver{})
	logging.Info("Service '%s' running in %d port", "REST-SERVER", 8080)
	log.Fatal(Srv.ListenAndServe())
}
