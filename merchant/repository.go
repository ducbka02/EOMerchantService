package merchant

import (
	"context"

	models "merchant-service/models"
)

// Repository represent the merchant's repository contract
type Repository interface {
	Fetch(ctx context.Context) (res []*models.Merchant, err error)
	GetByID(ctx context.Context, id int64) (*models.Merchant, error)
	GetImagesByID(ctx context.Context, id int64) ([]*models.Image, error)
	FilterByMulti(ctx context.Context, clause string) ([]*models.Merchant, error)
	SearchByKeyword(ctx context.Context, title string) ([]*models.Merchant, error)
	Update(ctx context.Context, ar *models.Merchant) error
	Store(ctx context.Context, a *models.Merchant) error
	Delete(ctx context.Context, id int64) error
}
