package db

type Instance struct {
	ID          int
	Name        string
	Description string
	Condition   string
	CreateUser  int
	CreateTime  int
	UpdateTime  int
	DeleteTime  int
}

var Instances = &InstanceRepo{}

type InstanceRepo struct {
}

func (t *InstanceRepo) QueryPageByUser(userId int, pageIndex int, pageSize int) (entities []*Instance, total int64) {
	tx := db.Where("user_id = ?", userId)
	tx.Count(&total)
	tx.Find(&entities)
	return
}
