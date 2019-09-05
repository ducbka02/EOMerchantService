package models

// Merchant represent the area model
type Merchant struct {
	ID           int64   `json:"mb_merchant_id"`
	MbCategoryID int64   `json:"mb_category_id"`
	AreaID       int64   `json:"area_id"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Phone        string  `json:"phone"`
	Image        string  `json:"image"`
	Description  string  `json:"description"`
	Delivery     int64   `json:"delivery"`
	TimeStart    string  `json:"time_start"`
	TimeEnd      string  `json:"time_end"`
}
