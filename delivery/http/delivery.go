package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

// Handler respoesent the httphandler for post, user
type Handler struct {
	PostUsecase     models.PostUsecase
	UserUsecase     models.UserUsecase
	CategoryUsecase models.CategoryUsecase
	CommentUsecase  models.CommentUsecase
	PostVoteUsecase models.PostVoteUsecase
}

// NewHandler will initialize the resources endpoint
func NewHandler(g *gin.Engine, handler *Handler) {

	// post endpoints
	post := g.Group("/post")
	{
		post.GET("/", handler.postGetAll)
		post.GET("/:id", handler.postGetByID)
		post.POST("/create", handler.postCreate)
		post.DELETE("/:id", handler.postDelete)
		// comment
		post.POST("/comment/create", handler.commentCreate)
	}

	// postvote endpoints
	postVote := g.Group("/post/vote")
	{
		postVote.POST("/like", handler.like)
		postVote.POST("/dislike", handler.dislike)
	}

	// user endpoints
	user := g.Group("/user")
	{
		user.POST("/signup", handler.signup)
		user.POST("/signin", handler.signin)
		user.POST("/signout", handler.signout)
	}

}
