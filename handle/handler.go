package handle

import "github.com/gin-gonic/gin"

type Handler interface {
	Handle(ctx *gin.Context)
}

var engine *gin.Engine

func Init(eng *gin.Engine) {
	engine = eng

	addHandler("/", &home{})

	addHandler("/instances", &instanceList{})
	addHandler("/instances/add", &instanceAdd{})
}

func addHandler(path string, handler Handler) {
	engine.GET(path, handler.Handle)
}
