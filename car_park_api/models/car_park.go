package models

import (
	"database/sql"
	"encoding/json"
)

// From basic_info_all
// related to vacancy_info_all_pretty.json
type CarPark struct {
	ParkID           string          `gorm:"primaryKey;type:varchar(20)" json:"park_id"`
	NameEN           string          `gorm:"type:varchar(50);not null" json:"name_en"`
	NameTC           string          `gorm:"type:varchar(50);not null" json:"name_tc"`
	NameSC           string          `gorm:"type:varchar(50);not null" json:"name_sc"`
	DisplayAddressEN string          `gorm:"type:varchar(200);not null" json:"display_address_en"`
	DisplayAddressTC string          `gorm:"type:varchar(200);not null" json:"display_address_tc"`
	DisplayAddressSC string          `gorm:"type:varchar(200);not null" json:"display_address_sc"`
	Latitude         float64         `gorm:"type:double(13,10);not null" json:"latitude"`
	Longitude        float64         `gorm:"type:double(13,10);not null" json:"longitude"`
	DistrictEN       sql.NullString  `gorm:"type:varchar(40)" json:"district_en"`
	DistrictTC       sql.NullString  `gorm:"type:varchar(40)" json:"district_tc"`
	DistrictSC       sql.NullString  `gorm:"type:varchar(40)" json:"district_sc"`
	ContactNo        sql.NullString  `gorm:"type:char(50)" json:"contact_no"` // I make 50 char, even though it describes 10 char in pdf. some data have 21 char, see in park_id "tdc65p1" of https://resource.data.one.gov.hk/td/carpark/basic_info_all".json
	OpeningStatus    sql.NullString  `gorm:"type:varchar(5)" json:"opening_status"`
	Height           sql.NullFloat64 `gorm:"type:double(3,1)" json:"height"`
	RemarkEN         sql.NullString  `gorm:"type:varchar(1200)" json:"remark_en"` // here's too
	RemarkTC         sql.NullString  `gorm:"type:varchar(1200)" json:"remark_tc"`
	RemarkSC         sql.NullString  `gorm:"type:varchar(1200)" json:"remark_sc"`
	WebsiteEN        sql.NullString  `gorm:"type:varchar(100)" json:"website_en"`
	WebsiteTC        sql.NullString  `gorm:"type:varchar(100)" json:"website_tc"`
	WebsiteSC        sql.NullString  `gorm:"type:varchar(100)" json:"website_sc"`
	CarparkPhoto     sql.NullString  `gorm:"type:varchar(100)" json:"carpark_photo"`

	VehicleTypes []*VehicleType `gorm:"foreignKey:CarParkID" json:"-"`
}

func (cp *CarPark) MarshalJSON() ([]byte, error) {
	type Alias CarPark

	var districtEN, districtTC, districtSC, contactNo, openingStatus, remarkEN, remarkTC, remarkSC, websiteEN, websiteTC, websiteSC, carparkPhoto string
	var height float64

	if cp.DistrictEN.Valid {
		districtEN = cp.DistrictEN.String
	}

	if cp.DistrictTC.Valid {
		districtTC = cp.DistrictTC.String
	}

	if cp.DistrictSC.Valid {
		districtSC = cp.DistrictSC.String
	}

	if cp.ContactNo.Valid {
		contactNo = cp.ContactNo.String
	}

	if cp.OpeningStatus.Valid {
		openingStatus = cp.OpeningStatus.String
	}

	if cp.Height.Valid {
		height = cp.Height.Float64
	}

	if cp.RemarkEN.Valid {
		remarkEN = cp.RemarkEN.String
	}

	if cp.RemarkTC.Valid {
		remarkTC = cp.RemarkTC.String
	}

	if cp.RemarkSC.Valid {
		remarkSC = cp.RemarkSC.String
	}

	if cp.WebsiteEN.Valid {
		websiteEN = cp.WebsiteEN.String
	}

	if cp.WebsiteTC.Valid {
		websiteTC = cp.WebsiteTC.String
	}

	if cp.WebsiteSC.Valid {
		websiteSC = cp.WebsiteSC.String
	}

	if cp.CarparkPhoto.Valid {
		carparkPhoto = cp.CarparkPhoto.String
	}

	return json.Marshal(&struct {
		*Alias
		DistrictEN    string  `json:"district_en"`
		DistrictTC    string  `json:"district_tc"`
		DistrictSC    string  `json:"district_sc"`
		ContactNo     string  `json:"contact_no"`
		OpeningStatus string  `json:"opening_status"`
		Height        float64 `json:"height"`
		RemarkEN      string  `json:"remark_en"`
		RemarkTC      string  `json:"remark_tc"`
		RemarkSC      string  `json:"remark_sc"`
		WebsiteEN     string  `json:"website_en"`
		WebsiteTC     string  `json:"website_tc"`
		WebsiteSC     string  `json:"website_sc"`
		CarparkPhoto  string  `json:"carpark_photo"`
	}{
		DistrictEN:    districtEN,
		DistrictTC:    districtTC,
		DistrictSC:    districtSC,
		ContactNo:     contactNo,
		OpeningStatus: openingStatus,
		Height:        height,
		RemarkEN:      remarkEN,
		RemarkTC:      remarkTC,
		RemarkSC:      remarkSC,
		WebsiteEN:     websiteEN,
		WebsiteTC:     websiteTC,
		WebsiteSC:     websiteSC,
		CarparkPhoto:  carparkPhoto,
		Alias:         (*Alias)(cp),
	})
}
