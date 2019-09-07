package merchant

import (
	"context"

	models "merchant-service/models"
)

// Usecase represent the merchant's repository contract
type Usecase interface {
	Fetch(ctx context.Context) ([]*models.Merchant, error)
	GetByID(ctx context.Context, id int64) (*models.Merchant, error)
	FilterByMulti(ctx context.Context, clause string) ([]*models.Merchant, error)
	SearchByKeyword(ctx context.Context, title string) ([]*models.Merchant, error)
	Update(ctx context.Context, ar *models.Merchant) error
	Store(ctx context.Context, a *models.Merchant) error
	Delete(ctx context.Context, id int64) error
}
