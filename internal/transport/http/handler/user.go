package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"onelab/internal/model"
	"onelab/internal/service"
	"onelab/internal/transport/http/middleware"
)

type UserHandler struct {
	Service *service.Service
	jwt     *middleware.JWTAuth
}

func (h *UserHandler) GetOrders(c echo.Context) error {
	userID := c.Request().Context().Value(model.ContextData{}).(string)
	orders, err := h.Service.User.GetOrders(c.Request().Context(), userID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var input model.User

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)
	id, err := h.Service.User.SignUp(c.Request().Context(), input)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.User.GetAllUsers(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByUsername(c echo.Context) error {
	user, err := h.Service.User.GetByUsername(c.Request().Context(), c.Param("username"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Login(c echo.Context) error {
	var input model.LogIn
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)

	user, err := h.Service.User.Login(c.Request().Context(), input)
	if err != nil {
		fmt.Println(err)
		return err
	}
	token, err := h.jwt.GenerateJWT(user.UserID)
	c.Set("User", user)
	return c.JSON(http.StatusOK, token)
}

type IUserHandler interface {
	GetOrders(c echo.Context) error
	SignUp(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetByUsername(c echo.Context) error
	Login(c echo.Context) error
}

func NewUserHandler(s *service.Service, jwt *middleware.JWTAuth) *UserHandler {
	return &UserHandler{
		Service: s,
		jwt:     jwt,
	}
}
