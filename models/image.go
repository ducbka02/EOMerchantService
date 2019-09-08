package models

// Image represent the merchant model
type Image struct {
	ID         int64  `json:"id"`
	MerchantID int64  `json:"mb_merchant_id"`
	Image      string `json:"image"`
}
