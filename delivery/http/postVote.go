package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pv *Handler) like(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"postvote": "liked"})
}

func (pv *Handler) dislike(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"postvote": "disliked"})
}
