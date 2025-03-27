package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/1ef7yy/effective_mobile_test/internal/routes"
	"github.com/1ef7yy/effective_mobile_test/internal/view"
	"github.com/1ef7yy/effective_mobile_test/pkg/logger"
)

func main() {

	logger := logger.NewLogger()

	logger.Info("starting service...")

	view, err := view.NewView(logger)
	if err != nil {
		logger.Fatal(fmt.Sprintf("could not initialize view layer: %s", err.Error()))
		return
	}

	mux := routes.InitRouter(view)

	logger.Info("initialized router...")

	serverAddr, ok := os.LookupEnv("SERVER_ADDRESS")

	if !ok {
		serverAddr = "localhost:3000"
		logger.Warnf("could not resolve SERVER_ADDRESS from environment, reverting to default: %s", serverAddr)
	}

	logger.Infof("starting server on %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		logger.Errorf("Error starting server: %s", err.Error())
	}

}
