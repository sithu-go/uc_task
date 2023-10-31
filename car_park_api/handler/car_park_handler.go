package handler

import (
	"net/http"
	"uc_task/car_park_api/dto"
	"uc_task/car_park_api/repo"
	"uc_task/car_park_api/utils"

	"github.com/gin-gonic/gin"
)

type carParkHandler struct {
	R    *gin.Engine
	repo *repo.Repository
}

func newCarParkHandler(h *Handler) *carParkHandler {
	return &carParkHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *carParkHandler) register() {
	group := ctr.R.Group("/api/carParks")

	group.GET("", ctr.getCarParks)
}

func (ctr *carParkHandler) getCarParks(c *gin.Context) {
	req := dto.PaginationRequest{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}

	carParks, err := ctr.repo.CarPark.FindAll(&req)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}

	c.JSON(http.StatusOK, carParks)
}
