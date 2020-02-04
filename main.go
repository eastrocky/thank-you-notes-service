package main

import (
	"github.com/eastrocky/thank-you-notes-service/handle"
	"github.com/eastrocky/thank-you-notes-service/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	repo := repository.NewMemoryRepository()
	r.GET("/thanks/:to", handle.ThanksGet(repo))
	r.POST("/thanks", handle.ThanksPost(repo))
	return r
}
