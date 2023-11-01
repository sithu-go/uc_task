package models

type VehicleTypeEnum string

const (
	VehicleTypeP  VehicleTypeEnum = "P"
	VehicleTypeM  VehicleTypeEnum = "M"
	VehicleTypeL  VehicleTypeEnum = "L"
	VehicleTypeH  VehicleTypeEnum = "H"
	VehicleTypeC  VehicleTypeEnum = "C"
	VehicleTypeT  VehicleTypeEnum = "T"
	VehicleTypeB  VehicleTypeEnum = "B"
	VehicleTypeO  VehicleTypeEnum = "O" // didn't describe in pdf, but get in api
	VehicleTypeN  VehicleTypeEnum = "N" // didn't describe in pdf, but get in api
	VehicleTypePD VehicleTypeEnum = "P_D"
	VehicleTypeMD VehicleTypeEnum = "M_D"
	VehicleTypeLD VehicleTypeEnum = "L_D"
	VehicleTypeHD VehicleTypeEnum = "H_D"
	VehicleTypeCD VehicleTypeEnum = "C_D"
	VehicleTypeTD VehicleTypeEnum = "T_D"
	VehicleTypeBD VehicleTypeEnum = "B_D"
)

// related to vacancy_all_pretty.json
type VehicleType struct {
	ID                uint               `gorm:"primaryKey"`
	CarParkID         string             `gorm:"type:varchar(20);not null"`
	Type              VehicleTypeEnum    `gorm:"type:char(3);not null;type:ENUM('P', 'M', 'L', 'H', 'C', 'T', 'B', 'O', 'N', 'P_D', 'M_D', 'L_D', 'H_D', 'C_D', 'T_D', 'B_D')"`
	ServiceCategories []*ServiceCategory `gorm:"foreignKey:VehicleTypeID"`
}
