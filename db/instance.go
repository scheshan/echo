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
