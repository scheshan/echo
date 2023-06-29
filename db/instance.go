package db

type Instance struct {
	ID          int
	Name        string
	Description string
	CreateUser  int
	CreateTime  int64
	UpdateTime  int64
	DeleteTime  int64
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

func (t *InstanceRepo) Add(instance *Instance) int64 {
	tx := db.Save(instance)
	return tx.RowsAffected
}
