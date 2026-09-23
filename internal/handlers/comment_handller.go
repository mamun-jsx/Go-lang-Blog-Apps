package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/services"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/utils"
)

type CommentHandler struct {
	Service *services.CommentService
}

func NewCommentHandler(s *services.CommentService) *CommentHandler {
	return &CommentHandler{Service: s}
}

type AddCommentReq struct {
	Content  string  `json:"content,omitempty"`
	ParentID *string `json:"parent_id,omitempty"`
}

func (h *CommentHandler) Add(c echo.Context) error {
	postID := c.Param("id")
	var req AddCommentReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid payload")

	}
	var userID *string
	if u := c.Get("user_id"); u != nil {
		uid := u.(string)
		userID = &uid
	}
	comment, err := h.Service.Add(postID, userID, req.ParentID, req.Content)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())

	}
	return utils.JSON(c, http.StatusCreated, true, "comment added", comment)
}

func (h *CommentHandler) List(c echo.Context) error {
	postID := c.Param(":id")
	comment, err := h.Service.ListByPost(postID)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())

	}
	return utils.JSON(c, http.StatusOK, true, "comment", comment)

}
