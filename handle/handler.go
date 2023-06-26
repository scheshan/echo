package handle

import "github.com/gin-gonic/gin"

type Handler interface {
	Path() string
	Handle(ctx *gin.Context)
}

var Handlers []Handler

func init() {
	Handlers = make([]Handler, 0)

	Handlers = append(Handlers, &home{})
	Handlers = append(Handlers, &instanceList{})
}
