package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	group := router.Group("/v1")
	{
		group.GET("/get/:name", loginPage)
	}

	router.Run()
}

func loginPage(context *gin.Context) {
	name := context.Param("name")
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
