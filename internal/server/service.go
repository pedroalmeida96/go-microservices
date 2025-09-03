package server

import (
	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllServices(ctx echo.Context) error {
	vendors, err := s.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

func (s *EchoServer) AddService(ctx echo.Context) error {
	//parse json
	service := new(models.Service)
	err := ctx.Bind(service)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	//save to db
	savedService, err := s.DB.AddService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, savedService)
}
