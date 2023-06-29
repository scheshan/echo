package handle

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Handle(ctx *gin.Context)
}

var engine *gin.Engine

func Init(eng *gin.Engine) {
	engine = eng

	addHandler("/", &home{})

	addHandler("/instances", &instanceList{})
	addHandler("/instances/add", &instanceAdd{})
	addPostHandler("/instances/save", &instanceSave{})
}

func addHandler(path string, handler Handler) {
	engine.GET(path, handler.Handle)
}

func addPostHandler(path string, handler Handler) {
	engine.POST(path, handler.Handle)
}
