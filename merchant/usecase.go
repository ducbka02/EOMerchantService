package merchant

import (
	"context"

	models "merchant-service/models"
)

// Usecase represent the merchant's repository contract
type Usecase interface {
	Fetch(ctx context.Context, page string, offset string) ([]*models.Merchant, error)
	FetchCategories(ctx context.Context) ([]*models.MbDiscoveryCategory, error)
	GetByID(ctx context.Context, id int64) (*models.Merchant, error)
	GetImagesByID(ctx context.Context, id int64) ([]*models.Image, error)
	FilterByMulti(ctx context.Context, clause string, page string, offset string) ([]*models.Merchant, error)
	SearchByKeyword(ctx context.Context, title string) ([]*models.Merchant, error)
	Update(ctx context.Context, ar *models.Merchant) error
	Store(ctx context.Context, a *models.Merchant) error
	Delete(ctx context.Context, id int64) error
}
