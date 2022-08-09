package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/config"
	delivery "github.com/quazar2000/my-go-clean-arch-temp/delivery/http"
	"github.com/quazar2000/my-go-clean-arch-temp/repository"
	postgresql "github.com/quazar2000/my-go-clean-arch-temp/repository/postgres"
	"github.com/quazar2000/my-go-clean-arch-temp/usecase"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load config ..
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conf)
	}

	// create new logrus logger for postgres logging adapter
	logger := logrus.New()
	// Create pool of connection for DB
	db, err := postgresql.NewConnectPostgresDB(&conf, logrusadapter.NewLogger(logger), pgx.LogLevelInfo)
	if err != nil {
		log.Fatal("Unable to connect ot database:", err)
	}
	defer db.Close()

	// Initialize server
	route := gin.Default()

	// repositories ...
	postRepository := repository.NewPostgresPostRepository(db)
	userRepository := repository.NewPostgresUserRepository(db)
	categoryRepository := repository.NewPostgresCategoryRepository(db)
	commentRepository := repository.NewPostgresCommentRepository(db)
	postvoteRepository := repository.NewPostgresPostVoteRepository(db)
	commentvoteRepository := repository.NewPostgresCommentVoteRepository(db)

	// context timeout setup
	timeoutContext := time.Duration(conf.Context.Timeout) * time.Second

	// usecases ...
	postUsecase := usecase.NewPostUsecase(postRepository, timeoutContext)
	userUsecase := usecase.NewUserUsecase(userRepository)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)
	postvoteUsecase := usecase.NewPostVoteUsecase(postvoteRepository)
	commentvoteUsecase := usecase.NewCommentVoteUsecase(commentvoteRepository)

	// delivery ...
	handler := &delivery.Handler{
		PostUsecase:        postUsecase,
		UserUsecase:        userUsecase,
		CategoryUsecase:    categoryUsecase,
		CommentUsecase:     commentUsecase,
		PostVoteUsecase:    postvoteUsecase,
		CommentVoteUsecase: commentvoteUsecase,
	}

	delivery.NewHandler(route, handler)

	err = route.Run(conf.Server.Address) // listen and serve on 0.0.0.0:9090
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
