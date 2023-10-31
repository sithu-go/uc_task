package handler

import (
	"uc_task/car_park_api/ds"
	"uc_task/car_park_api/middleware"
	"uc_task/car_park_api/repo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	R    *gin.Engine
	repo *repo.Repository
}

type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {
	repo := repo.NewRepository(c.DS)
	return &Handler{
		R:    c.R,
		repo: repo,
	}
}

func (h *Handler) Register() {
	h.R.Use(middleware.Cors())

	// car park handler
	carParkHandler := newCarParkHandler(h)
	carParkHandler.register()

}
