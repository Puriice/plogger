package logger

import (
	"net/http"

	"github.com/puriice/golibs/pkg/middleware"
	"github.com/puriice/golibs/pkg/server"
	"github.com/puriice/plogger/internal/handler/logger"
	"github.com/puriice/plogger/internal/repository"
)

func RegisterRoute(s *server.Server) {
	router := http.NewServeMux()

	loggerRepo := repository.NewPostgresLoggerRepository(s.Database)
	loggerHandler := logger.NewHandler(loggerRepo)
	loggerHandler.RegisterRoute(router)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	s.Handler = middleware.Logger(mux)
}
