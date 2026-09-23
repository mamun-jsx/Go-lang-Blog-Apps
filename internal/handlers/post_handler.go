package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/services"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/utils"
)

type PostHandler struct {
	Service *services.PostService
}

func NewPostHandler(s *services.PostService) *PostHandler {
	return &PostHandler{Service: s}
}

type CreatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreatePost creates a new post
func (h *PostHandler) CreatePost(c *echo.Context) error {
	var req CreatePostReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid request")
	}

	userIDStr, ok := c.Get("user_id").(string)
	if !ok {
		return utils.Err(c, http.StatusUnauthorized, "Unauthorized")
	}

	authorID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid user id")
	}

	post, err := h.Service.Create(authorID, req.Title, req.Content)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}

	return utils.JSON(c, http.StatusCreated, true, "post created", post)
}

// ListPosts returns all posts
func (h *PostHandler) ListPosts(c *echo.Context) error {
	posts, err := h.Service.GetAll()
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusOK, true, "posts", posts)
}

// GetPost returns a single post by ID
func (h *PostHandler) GetPost(c *echo.Context) error {
	id := c.Param("id")
	post, err := h.Service.GetByID(id)
	if err != nil {
		return utils.Err(c, http.StatusNotFound, "post not found")
	}
	return utils.JSON(c, http.StatusOK, true, "post", post)
}

// UpdatePost updates an existing post
func (h *PostHandler) UpdatePost(c *echo.Context) error {
	id := c.Param("id")
	var req UpdatePostReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid request")
	}

	post, err := h.Service.Update(id, req.Title, req.Content)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}

	return utils.JSON(c, http.StatusOK, true, "post updated", post)
}

// DeletePost deletes a post by ID
func (h *PostHandler) DeletePost(c *echo.Context) error {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusOK, true, "post deleted", nil)
}
