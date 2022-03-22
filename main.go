package main

import (
	"gca/config"
	"gca/handler"
	"gca/repository"
	"gca/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Db()

	repositoryUser := repository.NewRepositoryUser(db)
	serviceUser := service.NewServiceUser(repositoryUser)
	handler := handler.NewHandlerUser(serviceUser)

	router := gin.Default()

	user := router.Group("/api/v1/user")
	user.GET("/", handler.GetUser)

	router.Run()
}
