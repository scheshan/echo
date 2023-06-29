package handle

import (
	"echo/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type instanceSave struct {
}

func (t *instanceSave) Handle(ctx *gin.Context) {
	req := &InstanceSaveRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.AbortWithError(200, err)
	}

	entity := &db.Instance{}
	entity.Name = req.Name
	entity.Description = req.Description
	entity.CreateTime = time.Now().UnixMilli()
	entity.CreateUser = 1
	entity.UpdateTime = time.Now().UnixMilli()

	db.Instances.Add(entity)
	ctx.Redirect(302, fmt.Sprintf("/instances/%d", entity.ID))
}

type InstanceSaveRequest struct {
	ID          int
	Name        string
	Description string
}
