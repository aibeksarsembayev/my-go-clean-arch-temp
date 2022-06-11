package main

import (
	"fmt"
	"log"

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

// func init() {
// 	config.LoadConfig()
// }

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

	route.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})

	// repositories ...
	postRepository := repository.NewPostgresPostRepository(db)
	userRepository := repository.NewPostgresUserRepository(db)

	// usecases ...
	postUsecase := usecase.NewPostUsecase(postRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)

	// delivery ...
	handler := &delivery.Handler{
		PostUsecase: postUsecase,
		UserUsecase: userUsecase,
	}

	delivery.NewHandler(route, handler)

	route.Run() // listen and serve on 0.0.0.0:8080
}
