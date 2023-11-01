package dto

import (
	"strings"
	"time"
)

type APIServiceCategory struct {
	Category    string     `json:"category"`
	VacancyType string     `json:"vacancy_type"`
	Vacancy     int        `json:"vacancy"`
	LastUpdate  CustomTime `json:"lastupdate"`
}

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	input := string(b)
	if input == "null" {
		return nil
	}

	// Remove double quotes around the time value
	input = strings.Trim(input, "\"")

	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, input)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

type APIVehicleType struct {
	Type            string                `json:"type"`
	ServiceCategory []*APIServiceCategory `json:"service_category"`
}

type APICarPark struct {
	ParkID      string            `json:"park_id"`
	VehicleType []*APIVehicleType `json:"vehicle_type"`
}

type APICarParkVacancyData struct {
	CarPark []*APICarPark `json:"car_park"`
}
