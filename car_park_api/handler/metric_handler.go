package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricsHandler struct {
	R *gin.Engine
}

func newMetricsHandler(h *Handler) *metricsHandler {
	return &metricsHandler{
		R: h.R,
	}
}

func (ctr *metricsHandler) register() {
	group := ctr.R.Group("/api/metrics")

	// see in
	// http_requests_total
	// http_request_errors_totals
	// http_request_duration_seconds_sum
	group.GET("", gin.WrapH(promhttp.Handler()))
}
