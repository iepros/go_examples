package helloworld

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取Gin引擎
	engine := gin.Default()
	//通过引擎注册一个get请求接口
	engine.GET("/hello", func(context *gin.Context) {
		//设置返回数据，返回json格式的模板
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//发布并设置端口为8080，看源码可得出，默认端口为8080，不能设置多个端口，设置多个端口，抛出异常
	engine.Run(":8080")
}
