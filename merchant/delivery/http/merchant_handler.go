package http

import (
	"context"
	"merchant-service/merchant"
	"merchant-service/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// MerchantHandler  represent the httphandler for merchant
type MerchantHandler struct {
	MUsecase merchant.Usecase
}

// NewMerchantHandler will initialize the merchants/ resources endpoint
func NewMerchantHandler(e *echo.Echo, us merchant.Usecase) {
	handler := &MerchantHandler{
		MUsecase: us,
	}
	e.GET("/merchant/merchants", handler.FetchMerchant)
	e.GET("/merchant/area", handler.FetchArea)
	e.GET("/merchant/:id", handler.GetByID)
	e.GET("/merchant/filter", handler.FilterByMulti)
	e.GET("/merchant/search", handler.SearchByKeyword)
	e.GET("/merchant/categories", handler.FetchCategories)
}

// FetchArea will fetch the merchant based on given params
func (a *MerchantHandler) FetchArea(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, err := a.MUsecase.FetchArea(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": 1,
		"data":   listAr,
	})
}

// FetchCategories will fetch the merchant based on given params
func (a *MerchantHandler) FetchCategories(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, err := a.MUsecase.FetchCategories(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": 1,
		"data":   listAr,
	})
}

// FetchMerchant will fetch the merchant based on given params
func (a *MerchantHandler) FetchMerchant(c echo.Context) error {
	page := c.QueryParam("page")
	offset := c.QueryParam("offset")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, count, err := a.MUsecase.Fetch(ctx, page, offset)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	if len(page) != 0 && len(offset) != 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"status": 1,
			"data":   listAr,
			"page":   page,
			"offset": offset,
			"total":  count,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": 1,
		"data":   listAr,
		"total":  count,
	})
}

// Store new merchant to database
func (a *MerchantHandler) Store(c echo.Context) error {
	return c.String(http.StatusOK, "Store")
}

// GetByID an merchant by id
func (a *MerchantHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	mers, err := a.MUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": 1,
		"data":   mers,
	})
}

// FilterByMulti some merchant by clause
func (a *MerchantHandler) FilterByMulti(c echo.Context) error {
	queryParams := c.QueryParams()
	page := c.QueryParam("page")
	offset := c.QueryParam("offset")
	queryParamsSlice := []string{}
	for key, value := range queryParams {
		if key != "page" && key != "offset" {
			if key == "keyword" {
				if value[0] != "null" {
					queryParamsSlice = append(queryParamsSlice, "name like \"%"+value[0]+"%\"")
					//queryParamsSearch := fmt.Sprintf("name like CONCAT('%s', CONVERT('%s', BINARY), '%s')", "%", value[0], "%")
					//fmt.Println(queryParamsSearch)
					//queryParamsSlice = append(queryParamsSlice, queryParamsSlice)
				}
			} else if value[0] != "null" {
				if strings.Contains(value[0], ",") {
					queryParamsSlice = append(queryParamsSlice, key+" in ("+value[0]+")")
				} else {
					queryParamsSlice = append(queryParamsSlice, key+"="+value[0])
				}
			}
		}
	}
	queryParamsString := strings.Join(queryParamsSlice[:], " AND ")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, count, err := a.MUsecase.FilterByMulti(ctx, queryParamsString, page, offset)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	if len(page) != 0 && len(offset) != 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"status": 1,
			"data":   listAr,
			"page":   page,
			"offset": offset,
			"total":  count,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": 1,
		"data":   listAr,
		"total":  count,
	})
}

// SearchByKeyword some merchant by clause
func (a *MerchantHandler) SearchByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if len(keyword) == 0 {
		return c.JSON(http.StatusOK, "Chưa nhập từ khóa tìm kiếm")
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, count, err := a.MUsecase.SearchByKeyword(ctx, keyword)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": 1,
		"data":   listAr,
		"total":  count,
	})
}

// Delete an merchant by id
func (a *MerchantHandler) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
