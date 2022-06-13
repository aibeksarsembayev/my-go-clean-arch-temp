package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Handler) postCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "create"})
}

func (p *Handler) postGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "all"})
}

func (p *Handler) postGetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "id"})
}

func (p *Handler) postUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "update"})
}

func (p *Handler) postDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "delete"})
}
