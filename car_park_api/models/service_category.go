package models

import (
	"time"

	"gorm.io/gorm"
)

type CategoryEnum string

const (
	CategoryHourly  CategoryEnum = "HOURLY"
	CategoryDaily   CategoryEnum = "DAILY"
	CategoryMonthly CategoryEnum = "MONTHLY"
)

type VacancyTypeEnum string

const (
	VacancyTypeA VacancyTypeEnum = "A"
	VacancyTypeB VacancyTypeEnum = "B"
	VacancyTypeC VacancyTypeEnum = "C"
)

// related to vacancy_all_pretty.json
type ServiceCategory struct {
	ID             uint            `gorm:"primaryKey"`
	VehicleTypeID  uint            `gorm:"not null"`
	Category       CategoryEnum    `gorm:"type:ENUM('HOURLY', 'DAILY', 'MONTHLY');not null"`
	VacancyType    VacancyTypeEnum `gorm:"type:ENUM('A', 'B', 'C');not null;comment:A- with actual number, B - without acutal number, C -Closed"`
	CurrentVacancy int             `gorm:"not null;comment:For A: (0 - full, -1 - no data), For B: (0 - full, 1 - available, -1 - no data), For C: (always 0)"`
	CreatedAt      time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `json:"-"`
}
