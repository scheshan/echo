package handle

import (
	"github.com/gin-gonic/gin"
)

type home struct {
}

func (t *home) Path() string {
	return "/"
}

func (t *home) Handle(ctx *gin.Context) {
	ctx.HTML(200, "index", nil)
}
