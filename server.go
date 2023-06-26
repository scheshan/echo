package echo

import (
	"echo/handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() error {
	engine := gin.Default()

	for _, handler := range handle.Handlers {
		engine.GET(handler.Path(), handler.Handle)
	}

	//TODO 处理静态文件多了前缀的问题
	engine.StaticFS("static", http.FS(static))

	return engine.Run(":10023")
}
