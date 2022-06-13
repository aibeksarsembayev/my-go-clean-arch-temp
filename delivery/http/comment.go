package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Handler) commentCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"comment": "created"})
}

