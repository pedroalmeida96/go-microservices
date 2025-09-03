package server

import (
	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorId := ctx.QueryParam("vendorId")

	vendors, err := s.DB.GetAllProducts(ctx.Request().Context(), vendorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	//parse json
	product := new(models.Product)
	err := ctx.Bind(product)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	//save to db
	savedProduct, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, savedProduct)
}
