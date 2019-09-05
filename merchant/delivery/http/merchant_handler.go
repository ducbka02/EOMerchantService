package http

import (
	"context"
	"fmt"
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
	e.GET("/merchant/:id", handler.GetByID)
	e.GET("/merchant/filter", handler.FilterByMulti)
}

// FetchMerchant will fetch the merchant based on given params
func (a *MerchantHandler) FetchMerchant(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, err := a.MUsecase.Fetch(ctx)
	if err != nil {
		return c.String(http.StatusOK, "Fetch fail")
	}
	return c.JSON(http.StatusOK, listAr)
}

// Store new merchant to database
func (a *MerchantHandler) Store(c echo.Context) error {
	return c.String(http.StatusOK, "Store")
}

// GetByID an merchant by id
func (a *MerchantHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	fmt.Println(idP)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.MUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, art)
}

// FilterByMulti some merchant by clause
func (a *MerchantHandler) FilterByMulti(c echo.Context) error {
	queryParams := c.QueryParams()
	queryParamsSlice := []string{}
	for key, value := range queryParams {
		queryParamsSlice = append(queryParamsSlice, key+"="+value[0])
	}
	queryParamsString := strings.Join(queryParamsSlice[:], "&")
	fmt.Println(queryParamsString)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, err := a.MUsecase.FilterByMulti(ctx, queryParamsString)
	if err != nil {
		return c.String(http.StatusOK, "Fetch fail")
	}
	return c.JSON(http.StatusOK, listAr)

	//return c.String(http.StatusOK, "FilterByMulti")
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
