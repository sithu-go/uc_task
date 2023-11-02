package dto

import "uc_task/car_park_api/models"

type CarParkReq struct {
	ParkID  *string  `form:"park_id" json:"park_id"`
	Name    *string  `form:"name" json:"name"`
	Address *string  `form:"address" json:"address"`
	Lat     *float64 `form:"lat" json:"lat"`
	Lng     *float64 `form:"lng" json:"lng"`
	Radius  *float64 `form:"radius" json:"radius"`
	PaginationRequest
}

type VacacncyReq struct {
	ParkID      *string                 `form:"park_id" json:"park_id"`
	StartDate   *string                 `form:"start_date" json:"start_date"`
	EndDate     *string                 `form:"end_date" json:"end_date"`
	VehicleType *models.VehicleTypeEnum `form:"vehicle_type" json:"vehicle_type"`
	PaginationRequest
}
