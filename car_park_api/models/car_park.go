package models

import "database/sql"

// From basic_info_all
// related to vacancy_info_all_pretty.json
type CarPark struct {
	ParkID           string         `gorm:"primaryKey;type:varchar(20)"`
	NameEN           string         `gorm:"type:varchar(50);not null"`
	NameTC           string         `gorm:"type:varchar(50);not null"`
	NameSC           string         `gorm:"type:varchar(50);not null"`
	DisplayAddressEN string         `gorm:"type:varchar(200);not null"`
	DisplayAddressTC string         `gorm:"type:varchar(200);not null"`
	DisplayAddressSC string         `gorm:"type:varchar(200);not null"`
	Latitude         float64        `gorm:"type:double(13,10);not null"`
	Longitude        float64        `gorm:"type:double(13,10);not null"`
	DistrictEN       sql.NullString `gorm:"type:varchar(40)"`
	DistrictTC       sql.NullString `gorm:"type:varchar(40)"`
	DistrictSC       sql.NullString `gorm:"type:varchar(40)"`
	ContactNo        sql.NullString `gorm:"type:char(50)"` // I make 50 char, even though it describes 10 char in pdf. some data have 21 char, see in park_id "tdc65p1" of https://resource.data.one.gov.hk/td/carpark/basic_info_all.json
	OpeningStatus    sql.NullString `gorm:"type:varchar(5)"`
	Height           float64        `gorm:"type:double(3,1)"`
	RemarkEN         sql.NullString `gorm:"type:varchar(1200)"` // here's too
	RemarkTC         sql.NullString `gorm:"type:varchar(1200)"`
	RemarkSC         sql.NullString `gorm:"type:varchar(1200)"`
	WebsiteEN        sql.NullString `gorm:"type:varchar(100)"`
	WebsiteTC        sql.NullString `gorm:"type:varchar(100)"`
	WebsiteSC        sql.NullString `gorm:"type:varchar(100)"`
	CarparkPhoto     sql.NullString `gorm:"type:varchar(100)"`

	VehicleTypes []*VehicleType `gorm:"foreignKey:CarParkID"`
}
