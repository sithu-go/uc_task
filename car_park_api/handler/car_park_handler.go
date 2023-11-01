package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"uc_task/car_park_api/dto"
	"uc_task/car_park_api/metric"
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

// localhost:9000/api/carParks?page=1&page_size=20&lat=22.362&lng=114.102&radius=30&name=发停&address=Kon
func (ctr *carParkHandler) getCarParks(c *gin.Context) {
	start := time.Now()

	req := dto.CarParkReq{}
	if err := c.ShouldBind(&req); err != nil {
		// increment The number of errors in processing HTTP requests to API endpoints
		log.Println(err.Error())
		res := utils.GenerateValidationErrorResponse(err)
		metric.Metrics.ErrorCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprint(res.HttpStatusCode)).Inc()

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	// not complete, have to do some work
	carParks, err := ctr.repo.CarPark.FindAll(&req)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)

		// increment The number of errors in processing HTTP requests to API endpoints
		metric.Metrics.ErrorCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprint(res.HttpStatusCode)).Inc()

		c.JSON(res.HttpStatusCode, res)
		return
	}

	// Processing time of HTTP requests to the API endpoint
	elapsed := time.Since(start).Seconds()
	metric.Metrics.RequestDuration.WithLabelValues(c.Request.Method, c.Request.URL.Path).Observe(elapsed)

	// increment Total number of processed requests to API endpoints;
	metric.Metrics.RequestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path, "200").Inc()

	res := utils.GenerateSuccessResponse(carParks)
	c.JSON(res.HttpStatusCode, res)
}
