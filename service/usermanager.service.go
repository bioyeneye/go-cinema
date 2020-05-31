package service

type IUserManager interface {
}

type UserManager struct {
}

func NewUserManager() IUserManager {
	return &UserManager{}
}

func (manager *UserManager) SaveUser() string {
	return ""
}


func (manager *UserManager) FindUser() string {
	return ""
}
