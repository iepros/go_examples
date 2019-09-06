package httprequest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.New()

	//http://127.0.0.1:8081/get/ming
	engine.GET("/get/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	engine.POST("/post", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	engine.DELETE("/delete/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	engine.GET("/someJson", func(context *gin.Context) {
		names := []string{"li", "zhang", "ma"}
		context.SecureJSON(http.StatusOK, names)
	})

	engine.GET("/get/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		context.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})

	engine.Run(":8081")
}
