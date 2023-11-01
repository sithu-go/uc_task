package dto

type CarParkReq struct {
	Name    *string  `form:"name" json:"name"`
	Address *string  `form:"address" json:"address"`
	Lat     *float64 `form:"lat" json:"lat"`
	Lng     *float64 `form:"lng" json:"lng"`
	Radius  *float64 `form:"radius" json:"radius"`
	PaginationRequest
}
