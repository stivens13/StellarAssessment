package main

import (
	"github.com/gin-gonic/gin"
)

const baseUrl = "http://localhost"

func CreateRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/snippets/", InsertSnippet)
	router.GET("/snippets/:snippet", GetSnippet)

	return router
}
