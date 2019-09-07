package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"merchant-service/merchant"
	"merchant-service/models"
)

type mysqlMerchantRepository struct {
	DB *sql.DB
}

// NewMysqlMerchantRepository will create an object that represent the merchant.Repository interface
func NewMysqlMerchantRepository(db *sql.DB) merchant.Repository {
	return &mysqlMerchantRepository{
		DB: db,
	}
}

func (a *mysqlMerchantRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Merchant, error) {
	rows, err := a.DB.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	results := make([]*models.Merchant, 0)

	for rows.Next() {
		t := new(models.Merchant)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Address,
			&t.Latitude,
			&t.Longitude,
			&t.Phone,
			&t.Description,
			&t.Image,
			&t.Delivery,
			&t.TimeStart,
			&t.TimeEnd,
			&t.Facebook,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		results = append(results, t)
	}

	return results, nil
}

func (a *mysqlMerchantRepository) Fetch(ctx context.Context) ([]*models.Merchant, error) {
	query := `SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, image, delivery, time_start, time_end, facebook FROM mb_merchant`
	res, err := a.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *mysqlMerchantRepository) GetByID(ctx context.Context, id int64) (res *models.Merchant, err error) {
	query := `SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, image, delivery, time_start, time_end, facebook FROM mb_merchant WHERE mb_merchant_id = ?`

	list, err := a.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (a *mysqlMerchantRepository) FilterByMulti(ctx context.Context, clause string) ([]*models.Merchant, error) {
	query := fmt.Sprintf("SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, image, delivery, time_start, time_end, facebook FROM mb_merchant WHERE %s", clause)
	fmt.Println(query)
	list, err := a.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *mysqlMerchantRepository) SearchByKeyword(ctx context.Context, keyword string) ([]*models.Merchant, error) {
	query := `SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, image, delivery, time_start, time_end, facebook FROM mb_merchant WHERE name like ?`

	fmt.Println(query)
	list, err := a.fetch(ctx, query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *mysqlMerchantRepository) Update(ctx context.Context, m *models.Merchant) error {
	return nil
}

func (a *mysqlMerchantRepository) Store(ctx context.Context, m *models.Merchant) error {
	return nil
}

func (a *mysqlMerchantRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
