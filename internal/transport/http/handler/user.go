package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"onelab/internal/model"
	"onelab/internal/service"
	"onelab/internal/transport/http/middleware"
	"strconv"
)

type UserHandler struct {
	Service *service.Service
	jwt     *middleware.JWTAuth
}

func (h *UserHandler) GetBooks(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// ShowSpentMoney godoc
// @Summary      Show
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *UserHandler) GetSpentMoney(c echo.Context) error {
	res, _ := h.Service.User.GetSpentMoney(c.Request().Context())
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) RentBook(c echo.Context) error {
	tr, _ := h.Service.User.RentBook(c.Request().Context(), c.Param("username"), c.Param("bookId"))
	return c.JSON(http.StatusOK, tr)
}

func (h *UserHandler) ReturnBook(c echo.Context) error {
	return h.Service.User.ReturnBook(c.Request().Context(), c.Param("orderId"))
}

func (h *UserHandler) ReplenishBalance(c echo.Context) error {
	value, _ := strconv.ParseFloat(c.Param("balance"), 32)
	b := float32(value)
	t, _ := h.Service.User.ReplenishBalance(c.Request().Context(), c.Param("username"), b)
	return c.JSON(http.StatusOK, t)
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

func (h *UserHandler) GetOrders(c echo.Context) error {
	userID := c.Request().Context().Value(model.ContextData{}).(string)
	orders, err := h.Service.User.GetOrders(c.Request().Context(), userID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

type IUserHandler interface {
	GetOrders(c echo.Context) error
	SignUp(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetByUsername(c echo.Context) error
	Login(c echo.Context) error
	GetBooks(c echo.Context) error
	GetSpentMoney(c echo.Context) error
	RentBook(c echo.Context) error
	ReturnBook(c echo.Context) error
	ReplenishBalance(c echo.Context) error
}

func NewUserHandler(s *service.Service, jwt *middleware.JWTAuth) *UserHandler {
	return &UserHandler{
		Service: s,
		jwt:     jwt,
	}
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
