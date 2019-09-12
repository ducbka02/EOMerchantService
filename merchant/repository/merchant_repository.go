package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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
			&t.MbCategoryID,
			&t.AreaID,
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

func (a *mysqlMerchantRepository) fetchCategories(ctx context.Context, query string, args ...interface{}) ([]*models.MbDiscoveryCategory, error) {
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

	results := make([]*models.MbDiscoveryCategory, 0)

	for rows.Next() {
		t := new(models.MbDiscoveryCategory)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Description,
			&t.Code,
			&t.Image,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		results = append(results, t)
	}

	return results, nil
}

func (a *mysqlMerchantRepository) fetchDetail(ctx context.Context, query string, id int64) ([]*models.Merchant, error) {
	rows, err := a.DB.QueryContext(ctx, query)
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

	images, _ := a.GetImagesByID(ctx, id)

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
			&t.MbCategoryID,
			&t.AreaID,
			&t.Image,
			&t.Delivery,
			&t.TimeStart,
			&t.TimeEnd,
			&t.Facebook,
		)
		t.Images = images
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		results = append(results, t)
	}

	return results, nil
}

func (a *mysqlMerchantRepository) Fetch(ctx context.Context, page string, offset string) ([]*models.Merchant, error) {
	query := `SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, mb_category_id, area_id, image, delivery, time_start, time_end, facebook FROM mb_merchant`
	if len(page) != 0 && len(offset) != 0 {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return nil, errors.New("Page must be a number")
		}
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return nil, errors.New("Offset must be a number")
		}
		if pageInt < 0 || offsetInt < 0 {
			return nil, errors.New("Could not enter a negative number")
		}
		query = fmt.Sprintf("SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, mb_category_id, area_id, image, delivery, time_start, time_end, facebook FROM mb_merchant ORDER BY mb_merchant_id ASC LIMIT %d, %s", (pageInt-1)*offsetInt, offset)
	}

	res, err := a.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *mysqlMerchantRepository) FetchCategories(ctx context.Context) ([]*models.MbDiscoveryCategory, error) {
	query := `SELECT mb_category_id, name, description, code, image FROM mb_merchant_category`
	res, err := a.fetchCategories(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *mysqlMerchantRepository) GetByID(ctx context.Context, id int64) (res *models.Merchant, err error) {
	query := fmt.Sprintf("SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, mb_category_id, area_id, image, delivery, time_start, time_end, facebook FROM mb_merchant WHERE mb_merchant_id = %d", id)

	list, err := a.fetchDetail(ctx, query, id)
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

func (a *mysqlMerchantRepository) GetImagesByID(ctx context.Context, id int64) ([]*models.Image, error) {
	query := fmt.Sprintf("SELECT id, mb_merchant_id, image FROM mb_merchant_image WHERE mb_merchant_id = %d", id)

	rows, err := a.DB.QueryContext(ctx, query)
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

	results := make([]*models.Image, 0)

	for rows.Next() {
		t := new(models.Image)
		err = rows.Scan(
			&t.ID,
			&t.MerchantID,
			&t.Image,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		results = append(results, t)
	}

	return results, nil
}

func (a *mysqlMerchantRepository) FilterByMulti(ctx context.Context, clause string, page string, offset string) ([]*models.Merchant, error) {
	query := fmt.Sprintf("SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, mb_category_id, area_id, image, delivery, time_start, time_end, facebook FROM mb_merchant WHERE %s", clause)
	if len(page) != 0 && len(offset) != 0 {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return nil, errors.New("Page must be a number")
		}
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return nil, errors.New("Offset must be a number")
		}
		if pageInt < 0 || offsetInt < 0 {
			return nil, errors.New("Could not enter a negative number")
		}
		query = fmt.Sprintf("SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, mb_category_id, area_id, image, delivery, time_start, time_end, facebook FROM mb_merchant WHERE %s ORDER BY mb_merchant_id ASC LIMIT %d, %s", clause, (pageInt-1)*offsetInt, offset)
	}
	list, err := a.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *mysqlMerchantRepository) SearchByKeyword(ctx context.Context, keyword string) ([]*models.Merchant, error) {
	query := `SELECT mb_merchant_id, name, address, latitude, longitude, phone, description, mb_category_id, area_id, image, delivery, time_start, time_end, facebook FROM mb_merchant WHERE name like ?`

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
