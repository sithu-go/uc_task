package dto

type APICarParkSingle struct {
	ParkID           string  `json:"park_id"`
	NameEN           string  `json:"name_en"`
	NameTC           string  `json:"name_tc"`
	NameSC           string  `json:"name_sc"`
	DisplayAddressEN string  `json:"displayAddress_en"`
	DisplayAddressTC string  `json:"displayAddress_tc"`
	DisplayAddressSC string  `json:"displayAddress_sc"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	DistrictEN       string  `json:"district_en"`
	DistrictTC       string  `json:"district_tc"`
	DistrictSC       string  `json:"district_sc"`
	ContactNo        string  `json:"contactNo"`
	OpeningStatus    string  `json:"opening_status"`
	Height           float64 `json:"height"`
	RemarkEN         string  `json:"remark_en"`
	RemarkTC         string  `json:"remark_tc"`
	RemarkSC         string  `json:"remark_sc"`
	WebsiteEN        string  `json:"website_en"`
	WebsiteTC        string  `json:"website_tc"`
	WebsiteSC        string  `json:"website_sc"`
	CarparkPhoto     string  `json:"carpark_photo"`
}

type APICarParks struct {
	CarPark []APICarParkSingle `json:"car_park"`
}
