package usecase

import (
	"context"
	"time"

	"merchant-service/merchant"
	"merchant-service/models"
)

type merchantUsecase struct {
	merchantRepo   merchant.Repository
	contextTimeout time.Duration
}

// NewMerchantUsecase will create new an merchantUsecase object representation of merchant.Usecase interface
func NewMerchantUsecase(a merchant.Repository, timeout time.Duration) merchant.Usecase {
	return &merchantUsecase{
		merchantRepo:   a,
		contextTimeout: timeout,
	}
}

func (a *merchantUsecase) Fetch(c context.Context, page string, offset string) ([]*models.Merchant, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listAr, err := a.merchantRepo.Fetch(ctx, page, offset)
	if err != nil {
		return nil, err
	}

	return listAr, nil
}

func (a *merchantUsecase) FetchArea(c context.Context) ([]*models.Area, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listAr, err := a.merchantRepo.FetchArea(ctx)
	if err != nil {
		return nil, err
	}

	return listAr, nil
}

func (a *merchantUsecase) FetchCategories(c context.Context) ([]*models.MbDiscoveryCategory, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listAr, err := a.merchantRepo.FetchCategories(ctx)
	if err != nil {
		return nil, err
	}

	return listAr, nil
}

func (a *merchantUsecase) GetByID(c context.Context, id int64) (*models.Merchant, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.merchantRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *merchantUsecase) GetImagesByID(c context.Context, id int64) ([]*models.Image, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.merchantRepo.GetImagesByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *merchantUsecase) FilterByMulti(c context.Context, clause string, page string, offset string) ([]*models.Merchant, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.merchantRepo.FilterByMulti(ctx, clause, page, offset)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *merchantUsecase) SearchByKeyword(c context.Context, keyword string) ([]*models.Merchant, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.merchantRepo.SearchByKeyword(ctx, keyword)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *merchantUsecase) Update(ctx context.Context, m *models.Merchant) error {
	return nil
}

func (a *merchantUsecase) Store(ctx context.Context, m *models.Merchant) error {
	return nil
}

func (a *merchantUsecase) Delete(ctx context.Context, id int64) error {
	return nil
}
