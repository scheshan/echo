package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type home struct {
}

func (t *home) Path() string {
	return "/"
}

func (t *home) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
