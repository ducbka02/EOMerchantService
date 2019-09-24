package models

import (
	"gopkg.in/guregu/null.v3"
)

// Area represent the area model
type Area struct {
	ID          int64       `json:"area_id"`
	RegionID    null.Int    `json:"region_id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	Image       null.String `json:"image"`
}
