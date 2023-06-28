package handle

import (
	"echo/db"
	"github.com/gin-gonic/gin"
)

type instanceList struct {
}

func (t *instanceList) Handle(ctx *gin.Context) {
	//entities, total := db.Instances.QueryPageByUser(1, 1, 10)
	entities := make([]*db.Instance, 0)
	entities = append(entities, &db.Instance{
		Name: "测试的实例",
	})
	ctx.HTML(200, "instance/list", gin.H{
		"entities": entities,
	})
}
