package server

import (
	"github.com/sdkopen/sdkopen-go-webbase/logging"
	"github.com/sdkopen/sdkopen-go-webbase/observer"
)

type restObserver struct {
}

func (o restObserver) Close() {
	logging.Info("waiting to safely close the http server")
	if observer.WaitRunningTimeout() {
		logging.Warn("WaitGroup timed out, forcing close http server")
	}
	logging.Info("closing http server")
	if err := Srv.Shutdown(); err != nil {
		logging.Error("error when closing http server: %v", err)
	}
	Srv = nil
}
