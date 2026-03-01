package logger

import (
	"net/http"

	"github.com/puriice/httplibs/pkg/server"
	"github.com/puriice/plogger/internal/handler/logger"
	"github.com/puriice/plogger/internal/repository/postgres"
)

func RegisterRoute(s *server.Server) {
	router := http.NewServeMux()

	loggerRepo := postgres.NewLoggerRepository(s.Database)
	loggerHandler := logger.NewHandler(loggerRepo)
	loggerHandler.RegisterRoute(router)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	s.Handler = mux
}
