package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Handler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "create"})
}

func (p *Handler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "all"})
}

func (p *Handler) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "id"})
}

func (p *Handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "update"})
}

func (p *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "delete"})
}
