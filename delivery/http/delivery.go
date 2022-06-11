package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

// Handler respoesent the httphandler for post, user
type Handler struct {
	PostUsecase models.PostUsecase
	UserUsecase models.UserUsecase
}

// NewHandler will initialize the resources endpoint
func NewHandler(g *gin.Engine, handler *Handler) {

	post := g.Group("/post")
	{
		post.GET("/", handler.GetAll)
		post.GET("/:id", handler.GetByID)
		post.POST("/create", handler.Create)
		// post.POST("/:id:", handler.Update) // panic: only one wildcard per path segment is allowed, has: ':id:' in path '/post/:id:'
		post.DELETE("/:id", handler.Delete)
	}

	user := g.Group("/user")
	{
		user.POST("/signup", handler.signup)
		user.POST("/signin", handler.signin)
		user.POST("/signout", handler.signout)

	}

}
