package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//通过这个获取到的Engine，不包含默认的Logger、Recovery中间件
	//engine := gin.New()
	/**
	源码：使用Default获取，默认添加了Logger、Recovery中间件
	func Default() *Engine {
		debugPrintWARNINGDefault()
		engine := New()
		engine.Use(Logger(), Recovery())
		return engine
	}
	 */
	//engine := gin.Default()

	engine := gin.New()
	engine.Use(gin.Logger(),gin.Recovery())//添加Logger和Recovery中间件

	engine.GET("/students/:id",MyLogger(), func(context *gin.Context) {
		id := context.Param("id")
		context.String(http.StatusOK,id)
	})

	engine.Run()
}

//自定义中间件
/**
中间件中可以添加日志、权限校验等，类似网关，在请求到达接口前处理问题。
 */
func MyLogger() gin.HandlerFunc{
	return func(context *gin.Context) {
		fmt.Println("这是自定义中间件")
		id := context.Param("id")
		if id == "123" {
			context.Next()
			return
		}else {
			context.Abort()
			return
		}
	}
}
