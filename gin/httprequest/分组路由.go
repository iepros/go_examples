package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
分组路由，顾名思义就是把路由（请求链接）进行分组
如：/v1/user/1
	/v1/user

	/v2/user/1
	/v2/user
*/
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
