package main

import (
	"github.com/gin-gonic/gin"
	"go_examples/gin/httprequest/function"
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
	//创建一个组，并设置组的根路由
	group := router.Group("/v1")
	{
		//将路由的处理函数集中封装成函数
		group.GET("/student/:id", function.GetStudent)
	}

	//由源码可知，后面的函数是中间件函数
	routerGroup := router.Group("/v2", func(context *gin.Context) {
		//在此可以添加中间件，如：日志、权限校验等，类似spring AOP 中的前置事件
	})

	routerGroup.GET("/student/:id", function.GetStudent)
	routerGroup.DELETE("/student/:id", function.DeleteStudent)

	router.Run()
}
