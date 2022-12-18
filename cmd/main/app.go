package main

import (
	"diplom/internal/config"
	"diplom/internal/user"
	"diplom/pkg/logging"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := chi.NewRouter()

	logger.Info("get config")
	cfg := config.GetConfig()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router chi.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	logger.Info("listen tcp")
	address := fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("server is listening port %s", address)
	logger.Fatal(server.Serve(listener))
}
