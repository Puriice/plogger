package logger

import (
	"net/http"

	"github.com/puriice/httplibs/pkg/json"
	"github.com/puriice/httplibs/pkg/pgutils"
	"github.com/puriice/plogger/internal/constant"
	"github.com/puriice/plogger/internal/model"
	"github.com/puriice/plogger/internal/repository"
)

type Handler struct {
	repo repository.LoggerRepository
}

func NewHandler(repo repository.LoggerRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) GetLogs(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")

	logs, err := h.repo.GetLogByProject(r.Context(), projectID)

	err = pgutils.CheckError(err, w)

	if err != nil {
		return
	}

	if len(logs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.SendJSON(w, http.StatusOK, logs)
}

func (h *Handler) CreateLog(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")

	payload := new(model.Log)

	err := json.ParseJSON(r, payload)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if payload.Type == nil {
		defaultLogType := constant.DefaultLogType
		payload.Type = &defaultLogType
	}

	err = h.repo.CreateLog(r.Context(), projectID, *payload.Type, *payload.Message)

	err = pgutils.CheckError(err, w)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("GET /logs/{projectID}", h.GetLogs)
	router.HandleFunc("POST /logs/{projectID}", h.CreateLog)
}
