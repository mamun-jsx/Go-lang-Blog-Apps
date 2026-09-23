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
	Content  string  `json:"content"`
	ParentID *string `json:"parent_id,omitempty"`
}

type UpdateCommentReq struct {
	Content string `json:"content"`
}

// Add creates a new comment on a post
func (h *CommentHandler) Add(c *echo.Context) error {
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

// List returns all comments for a post
func (h *CommentHandler) List(c *echo.Context) error {
	postID := c.Param("id")
	comments, err := h.Service.ListByPost(postID)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusOK, true, "comments", comments)
}

// Update updates an existing comment
func (h *CommentHandler) Update(c *echo.Context) error {
	id := c.Param("id")
	var req UpdateCommentReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid payload")
	}
	comment, err := h.Service.Update(id, req.Content)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusOK, true, "comment updated", comment)
}

// Delete deletes a comment by ID
func (h *CommentHandler) Delete(c *echo.Context) error {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusOK, true, "comment deleted", nil)
}
