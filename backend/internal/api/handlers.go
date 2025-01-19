package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"photo-classifier/internal/database"
)

type Handler struct {
	db *database.Manager
}

func NewHandler(db *database.Manager) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetPhotos(c *gin.Context) {
	filter := database.PhotoFilter{
		Category: c.Query("category"),
		Album:    c.Query("album"),
		Sort:     c.Query("sort"),
	}

	photos, err := h.db.GetPhotos(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func (h *Handler) GetAlbums(c *gin.Context) {
	albums, err := h.db.GetAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func (h *Handler) CreateAlbum(c *gin.Context) {
	var album database.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAlbum, err := h.db.CreateAlbum(album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdAlbum})
}

func (h *Handler) UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album database.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album.ID = id
	updatedAlbum, err := h.db.UpdateAlbum(album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedAlbum})
}

func (h *Handler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.DeleteAlbum(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
