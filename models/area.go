package models

// Area represent the area model
type Area struct {
	ID       int64  `json:"area_id"`
	RegionID int64  `json:"region_id"`
	Name     string `json:"name"`
}
