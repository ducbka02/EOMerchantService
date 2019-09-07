package models

import (
	"gopkg.in/guregu/null.v3"
)

// Merchant represent the area model
type Merchant struct {
	ID           int64       `json:"mb_merchant_id"`
	MbCategoryID int64       `json:"mb_category_id"`
	AreaID       int64       `json:"area_id"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Latitude     float64     `json:"latitude"`
	Longitude    float64     `json:"longitude"`
	Phone        string      `json:"phone"`
	Image        string      `json:"image"`
	Description  string      `json:"description"`
	Delivery     null.Int    `json:"delivery"`
	TimeStart    null.String `json:"time_start"`
	TimeEnd      null.String `json:"time_end"`
	Facebook     null.String `json:"facebook"`
}
