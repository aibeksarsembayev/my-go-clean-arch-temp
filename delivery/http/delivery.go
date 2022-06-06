package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

// PostHandler respoesent the httphandler for post
type PostHandler struct {
	PUsecase models.PostUsecase
}

// NewPostHandler will initialize the post resources endpoint
func NewPostHandler(g *gin.Engine, us models.PostUsecase) {
	handler := &PostHandler{
		PUsecase: us,
	}
	g.GET("/post", handler.GetAll)
	g.GET("/post/:id", handler.GetByID)
	g.POST("/post", handler.Create)
	g.POST("/post/:id:", handler.Update)
	g.DELETE("/post/:id", handler.Delete)

}

func (p *PostHandler) Create(c *gin.Context) {

}

func (p *PostHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "all"})
}

func (p *PostHandler) GetByID(c *gin.Context) {
}

func (p *PostHandler) Update(c *gin.Context) {
}

func (p *PostHandler) Delete(c *gin.Context) {
}
