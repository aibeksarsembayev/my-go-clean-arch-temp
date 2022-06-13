package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cv *Handler) commentLike(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"commentvote": "liked"})

}

func (cv *Handler) commentDislike(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"commentvote": "disliked"})

}
