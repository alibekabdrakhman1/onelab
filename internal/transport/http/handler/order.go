package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"onelab/internal/model"
	"onelab/internal/service"
)

type OrderHandler struct {
	Service *service.Service
}

func (h *OrderHandler) Create(c echo.Context) error {
	var input model.Order

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)
	id, err := h.Service.Order.Create(c.Request().Context(), input)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *OrderHandler) ReturnBook(c echo.Context) error {
	return h.Service.Order.ReturnBook(c.Request().Context(), c.Param("orderID"))
}

func (h *OrderHandler) GetAllOrders(c echo.Context) error {
	orders, err := h.Service.Order.GetAllOrders(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetNotReturned(c echo.Context) error {
	orders, err := h.Service.Order.GetNotReturned(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetLastMonthOrders(c echo.Context) error {
	orders, err := h.Service.Order.GetLastMonthOrders(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}

type IOrderHandler interface {
	GetAllOrders(c echo.Context) error
	GetNotReturned(c echo.Context) error
	GetLastMonthOrders(c echo.Context) error
	Create(c echo.Context) error
	ReturnBook(c echo.Context) error
}

func NewOrderHandler(s *service.Service) *OrderHandler {
	return &OrderHandler{
		Service: s,
	}
}
