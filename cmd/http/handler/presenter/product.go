package presenter

type GetProductListRequest struct {
	CategoryID int64 `query:"category_id"`
	Page       int16 `query:"page"`
	Limit      int16 `query:"limit"`
}

type GetProductListResponse struct {
	ID                uint64  `json:"id"`
	ProductCategoryID uint64  `json:"category_id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	CategoryName      string  `json:"category_name"`
}
