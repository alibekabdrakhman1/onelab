package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"onelab/internal/model"
	"onelab/internal/service"
	transactions "onelab/proto"
)

type TransactionHandler struct {
	service *service.Service
	grpc    transactions.TransactionServiceClient
}

func (h *TransactionHandler) Create(c echo.Context) error {
	var m model.Transaction
	if err := c.Bind(&m); err != nil {
		return err
	}
	req := transactions.CreateTransRequest{
		Transaction: &transactions.Transaction{
			Username:    m.Username,
			Type:        m.TypeOfTransaction,
			Amount:      int32(m.Amount),
			Description: m.Description,
		}}
	transaction, err := h.grpc.Create(c.Request().Context(), &req)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transaction.Id)
}

func (h *TransactionHandler) Get(c echo.Context) error {
	id := c.Param("id")
	res, err := h.grpc.Get(c.Request().Context(), &transactions.GetTransRequest{
		TransactionID: id,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *TransactionHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	_, err := h.grpc.Delete(c.Request().Context(), &transactions.DeleteTransRequest{
		TransactionID: id,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "deleted")
}

func NewTransactionHandler(service *service.Service, grpc transactions.TransactionServiceClient) *TransactionHandler {
	return &TransactionHandler{
		service: service,
		grpc:    grpc,
	}
}

type ITransactionHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Delete(c echo.Context) error
}
