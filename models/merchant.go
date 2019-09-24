package models

import (
	"gopkg.in/guregu/null.v3"
)

// Merchant represent the area model
type Merchant struct {
	ID           int64       `json:"mb_merchant_id"`
	MbCategoryID null.Int    `json:"mb_category_id"`
	AreaID       null.Int    `json:"area_id"`
	Name         null.String `json:"name"`
	Address      null.String `json:"address"`
	Latitude     null.Float  `json:"latitude"`
	Longitude    null.Float  `json:"longitude"`
	Phone        null.String `json:"phone"`
	Image        null.String `json:"image"`
	Description  null.String `json:"description"`
	Delivery     null.Int    `json:"delivery"`
	TimeStart    null.String `json:"time_start"`
	TimeEnd      null.String `json:"time_end"`
	Facebook     null.String `json:"facebook"`
	Images       []*Image    `json:"images"`
}
