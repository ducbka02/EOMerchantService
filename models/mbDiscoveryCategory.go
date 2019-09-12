package models

// MbDiscoveryCategory represent the category model
type MbDiscoveryCategory struct {
	ID          int64  `json:"mb_discovery_category_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
