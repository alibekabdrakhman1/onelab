package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"onelab/internal/model"
	"onelab/internal/service"
)

type BookHandler struct {
	Service *service.Service
}

func (h *BookHandler) Update(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *BookHandler) Create(c echo.Context) error {
	var input model.Book

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)
	id, err := h.Service.Book.Create(c.Request().Context(), input)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *BookHandler) GetAvailable(c echo.Context) error {
	resp, err := h.Service.Book.GetAvailable(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	resp, err := h.Service.Book.GetAllBooks(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *BookHandler) GetByID(c echo.Context) error {
	resp, err := h.Service.Book.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

type IBookHandler interface {
	GetAvailable(c echo.Context) error
	GetAllBooks(c echo.Context) error
	GetByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
}

func NewBookHandler(s *service.Service) *BookHandler {
	return &BookHandler{
		Service: s,
	}
}
