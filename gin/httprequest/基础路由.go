package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个新的引擎，不使用默认的 Default()
	//由源码可知：默认的引擎比新建的引擎多了两个中间件：Logger、Recovery
	engine := gin.New()

	/**
	Get请求链接上包含参数，可通过`:name`添加
	*/
	//http://127.0.0.1:8081/user/ming		返回数据：{"name":"xiaohei"}
	//如果请求后面带 /  则找不到路径
	engine.GET("/user/:name", func(context *gin.Context) {
		fmt.Println("11111111111")
		//获取链接上的参数
		name := context.Param("name")
		//以JSON的格式返回链接上的参数
		context.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	/**
	Get请求链接上包含参数，弥补了上面的不足，同时满足:
		/user/:name/
		/user/:name/*action
	*/
	engine.GET("/user/:name/*action", func(context *gin.Context) {
		fmt.Println("22222222222222")
		name := context.Param("name")
		action := context.Param("action")
		//获取查询参数，问号后面的参数key-value   http://localhost:8081/user/xiaohei/2?type=1
		query := context.Query("type")
		context.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
			"type":   query,
		})
	})

	engine.POST("/user", func(context *gin.Context) {
		//获取表单(body)参数
		name := context.PostForm("name")
		age := context.PostForm("age")

		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	engine.PUT("/user", func(context *gin.Context) {

	})

	engine.DELETE("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	engine.Run(":8081")
}
