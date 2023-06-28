package handle

import (
	"github.com/gin-gonic/gin"
)

type instanceAdd struct {
}

func (t *instanceAdd) Handle(ctx *gin.Context) {
	ctx.HTML(200, "/instance/edit", gin.H{})
}
