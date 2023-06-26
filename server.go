package echo

import "github.com/gin-gonic/gin"

func Run() error {
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"hello": "world",
		})
	})

	return engine.Run(":10023")
}
