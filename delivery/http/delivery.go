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
	// g.GET("/post", handler.GetAll)
	// g.GET("/post/:id", handler.GetByID)
	// g.POST("/post/create", handler.Create)
	// // g.POST("/post/:id:", handler.Update) // panic: only one wildcard per path segment is allowed, has: ':id:' in path '/post/:id:'
	// g.DELETE("/post/:id", handler.Delete)

	post := g.Group("/post")
	{
		post.GET("/", handler.GetAll)
		post.GET("/:id", handler.GetByID)
		post.POST("/create", handler.Create)
		// post.POST("/:id:", handler.Update)
		post.DELETE("/:id", handler.Delete)
	}

}

func (p *PostHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "create"})
}

func (p *PostHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "all"})
}

func (p *PostHandler) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "id"})
}

func (p *PostHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "update"})
}

func (p *PostHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "delete"})
}
