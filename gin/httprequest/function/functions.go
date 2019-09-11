package function

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudent(context *gin.Context) {
	name := context.Param("name")
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func DeleteStudent(ctx *gin.Context) {
	//处理删除逻辑
}

func MiddleWare(ctx *gin.Context) {
	//中间件
}
