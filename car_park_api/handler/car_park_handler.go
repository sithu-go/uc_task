package handler

import (
	"fmt"
	"log"
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
	group.GET("/vacancy", ctr.getVacancyInfo)
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

		c.JSON(res.HttpStatusCode, res)
		return
	}

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

// localhost:9000/api/carParks/vacancy?page=1&page_size=2&start_date=2023-11-01 10:00:00&end_date=2023-11-03 10:00:00&vehicle_type=C
func (ctr *carParkHandler) getVacancyInfo(c *gin.Context) {
	start := time.Now()

	req := dto.VacacncyReq{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		metric.Metrics.ErrorCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprint(res.HttpStatusCode)).Inc()
		c.JSON(res.HttpStatusCode, res)
		return
	}

	carParks, err := ctr.repo.CarPark.FindVacancyData(&req)

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
