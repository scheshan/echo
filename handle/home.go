package handle

import (
	"echo/model"
	"github.com/gin-gonic/gin"
)

type home struct {
}

func (t *home) Path() string {
	return "/"
}

func (t *home) Handle(ctx *gin.Context) {
	ctx.JSON(200, model.Success[*home]())
}
