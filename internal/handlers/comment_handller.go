package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/services"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/utils"
)

type AuthHandler struct {
	Service *services.UserService
}

func NewAuthHandler(s *services.UserService) *AuthHandler {
	return &AuthHandler{Service: s}
}

type SignUpReq struct {
	UserName    string `json:"user_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var req SignUpReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid payload")
	}
	u, err := h.Service.Register(req.UserName, req.Email, req.Password)

	if err != nil {
		return utils.Err(c, http.StatusBadRequest, err.Error())

	}
	return utils.JSON(c, http.StatusCreated, true, "user_created", u)

}

type login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req login
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid payload")
	}
	user, token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid credentials ")
	}
	return utils.JSON(c, http.StatusOK, true, "login succesfull", map[string]interface{}{"token": token, "user": user})
}
