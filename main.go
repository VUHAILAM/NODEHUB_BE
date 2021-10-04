package main

import (
	"context"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/shutdown"

	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/cmd/server"
)

func main() {
	defer shutdown.SigtermHandler().Wait()
	s := server.InitServer()
	shutdown.SigtermHandler().RegisterErrorFuncContext(context.Background(), s.HttpServer.Shutdown)
	if err := s.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}

}
