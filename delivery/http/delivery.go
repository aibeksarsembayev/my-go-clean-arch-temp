package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

// Handler respoesent the httphandler for post, user
type Handler struct {
	PostUsecase        models.PostUsecase
	UserUsecase        models.UserUsecase
	CategoryUsecase    models.CategoryUsecase
	CommentUsecase     models.CommentUsecase
	PostVoteUsecase    models.PostVoteUsecase
	CommentVoteUsecase models.CommentVoteUsecase
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
	}

	// postvote endpoints
	postVote := g.Group("/post/vote")
	{
		postVote.POST("/like", handler.postLike)
		postVote.POST("/dislike", handler.postDislike)
	}

	// comment and its vote endpoints
	postComment := g.Group("/post/comment")
	{
		postComment.GET("/create", handler.commentCreate)

		postComment.POST("/create", handler.commentCreate)
		postComment.POST("/like", handler.commentLike)
		postComment.POST("/dislike", handler.commentDislike)
	}

	// user endpoints
	user := g.Group("/user")
	{
		user.POST("/signup", handler.signup)
		user.GET("/signup", handler.signup)

		user.GET("/signin", handler.signin)
		user.POST("/signin", handler.signin)
		user.GET("/signout", handler.signout)
	}

}
