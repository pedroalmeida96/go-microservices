package server

import (
	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendors, err := s.DB.GetAllVendors(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

func (s *EchoServer) AddVendor(ctx echo.Context) error {
	//parse json
	vendor := new(models.Vendor)
	err := ctx.Bind(vendor)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	//save to db
	savedVendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, savedVendor)
}

func (s *EchoServer) GetVendorById(ctx echo.Context) error {
	ID := ctx.Param("id")

	vendor, err := s.DB.GetVendorById(ctx.Request().Context(), ID)

	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusBadRequest, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusOK, vendor)

}

func (s *EchoServer) UpdateVendor(ctx echo.Context) error {
	ID := ctx.Param("id")
	vendor := new(models.Vendor)

	err := ctx.Bind(vendor)
	if err != nil {

	}

	if ID != vendor.VendorID {

	}

	updatedVendor, err := s.DB.UpdateVendor(ctx.Request().Context(), vendor)

	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusBadRequest, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusOK, updatedVendor)

}

func (s *EchoServer) DeleteVendor(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteVendor(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
