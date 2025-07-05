package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/yuki-maruyama/actions-with-cache/models"
)

type Handler struct {
	logger      *logrus.Logger
	httpRequests *prometheus.CounterVec
}

func NewHandler(logger *logrus.Logger, httpRequests *prometheus.CounterVec) *Handler {
	return &Handler{
		logger:      logger,
		httpRequests: httpRequests,
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	h.httpRequests.WithLabelValues(r.Method, "/health").Inc()
	h.logger.WithFields(logrus.Fields{
		"method":   r.Method,
		"endpoint": "/health",
	}).Info("Health check requested")
	
	w.Header().Set("Content-Type", "application/json")
	response := models.Response{
		Message:   "API is healthy",
		Status:    200,
		Timestamp: time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	h.httpRequests.WithLabelValues(r.Method, "/hello").Inc()
	h.logger.WithFields(logrus.Fields{
		"method":   r.Method,
		"endpoint": "/hello",
	}).Info("Hello endpoint requested")
	
	w.Header().Set("Content-Type", "application/json")
	response := models.Response{
		Message:   "Hello from Go API!",
		Status:    200,
		Timestamp: time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}