package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pv *Handler) postLike(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"postvote": "liked"})
}

func (pv *Handler) postDislike(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"postvote": "disliked"})
}
