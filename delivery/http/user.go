package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Handler) signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": "signup"})
}

func (u *Handler) signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": "signin"})
}

func (u *Handler) signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": "signout"})
}
