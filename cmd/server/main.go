package main

import (
	"log/slog"
	"net/http"
	"os"

	"pipelines/internal/handler"
	"pipelines/internal/service"
)

var logger = slog.New(
	slog.NewJSONHandler(os.Stdout, nil),
)

func main() {

	//metrics.Init()

	// Create service layer
	userService := service.NewUserService()

	// Create HTTP handler layer
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/pipelines", userHandler.Users)

	logger.Info(
		"user service started",
		"port",
		8080,
	)

	//handler := middleware.Metrics(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error(
			"server stopped",
			"error",
			err,
		)
	}
}
