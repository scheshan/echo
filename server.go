package echo

import (
	"echo/handle"
	"github.com/gin-gonic/gin"
)

func Run() error {
	engine := gin.Default()

	for _, handler := range handle.Handlers {
		engine.GET(handler.Path(), handler.Handle)
	}

	return engine.Run(":10023")
}
