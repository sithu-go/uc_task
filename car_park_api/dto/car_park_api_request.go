package dto

import (
	"uc_task/car_park_api/models"
)

type CarParkReq struct {
	ParkID  *string  `form:"park_id" json:"park_id"`
	Name    *string  `form:"name" json:"name"`
	Address *string  `form:"address" json:"address"`
	Lat     *float64 `form:"lat" json:"lat"`
	Lng     *float64 `form:"lng" json:"lng"`
	Radius  *float64 `form:"radius" json:"radius"`
	// OrderBy *string  `form:"order" json:"order_by" binding:"oneof='ASC' 'DESC'"`
	OrderBy *string `form:"order_by" json:"order_by"` // to see new car park
	PaginationRequest
}

type VacancyReq struct {
	ParkID         *string                 `form:"park_id" json:"park_id"`
	StartDate      *string                 `form:"start_date" json:"start_date"`
	EndDate        *string                 `form:"end_date" json:"end_date"`
	VehicleType    *models.VehicleTypeEnum `form:"vehicle_type" json:"vehicle_type"`
	VacancyType    *string                 `form:"vacancy_type" json:"vacancy_type" binding:"omitempty,oneof=A B C"`
	CurrentVacancy *int                    `form:"current_vacancy" json:"current_vacancy"`
	PaginationRequest
}
