package handle

import (
	"github.com/gin-gonic/gin"
)

type instanceList struct {
}

func (t *instanceList) Path() string {
	return "/instances"
}

func (t *instanceList) Handle(ctx *gin.Context) {
	//entities, total := db.Instances.QueryPageByUser(1, 1, 10)
	ctx.HTML(200, "instance/list", gin.H{})
}
