package main

import (
	"diplom/internal/user"
	"diplom/pkg/logging"
	"github.com/go-chi/chi/v5"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := chi.NewRouter()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, &logger)
}

func start(router chi.Router, logger *logging.Logger) {
	logger.Println("start application")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
