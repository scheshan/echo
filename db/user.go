package db

type User struct {
	ID         int
	UserName   string
	Password   string
	CreateTime int
	UpdateTime int
	DeleteTime int
}

var Users = &UserRepo{}

type UserRepo struct {
}

func (t *UserRepo) GetByAuth(userName string, password string) *User {
	return nil
}
