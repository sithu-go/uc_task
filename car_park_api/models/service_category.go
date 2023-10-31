package models

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
	VehicleTypeID  uint            `gorm:"not null;unique"`
	Category       CategoryEnum    `gorm:"type:ENUM('HOURLY', 'DAILY', 'MONTHLY');not null"`
	VacancyType    VacancyTypeEnum `gorm:"type:ENUM('A', 'B', 'C');not null"`
	CurrentVacancy int             `gorm:"not null"`
}
