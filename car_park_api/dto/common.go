package dto

type PaginationRequest struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"page_size" form:"page_size" binding:"required"`
}

type Response struct {
	ErrCode        uint64 `json:"err_code"`
	ErrMsg         string `json:"err_msg"`
	Data           any    `json:"data,omitempty"`
	HttpStatusCode int    `json:"-"`
}
