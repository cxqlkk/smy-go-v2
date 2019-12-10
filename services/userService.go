package services

import (
	"smyappTwo/dao"
	"smyappTwo/pojo"
)

type UserService interface {
	Create(user *pojo.User) error
	GetById(id string) *pojo.User
	List(user *pojo.User, pageNo, pageSize int) interface{}
	Delete(id string) error
	Update(usr *pojo.User) error
	GetByPhoneAndPassword(phoneNum,password string)*pojo.User
}

type userService struct {
	dao dao.UserDao
}

func NewUserService(d dao.UserDao) UserService {
	return &userService{d}
}

func (u *userService) Create(user *pojo.User) error {
	return u.dao.Create(user)
}
func (u *userService) GetById(id string) *pojo.User {
	return u.dao.GetById(id)
}
func (u *userService) List(user *pojo.User, pageNo, pageSize int) interface{} {
	list := u.dao.List(user, pageNo, pageSize)
	count := u.dao.Count(user)
	return map[string]interface{}{"data": list, "total": count}
}
func (u *userService) Delete(id string) error {

	return u.dao.Delete(id)
}
func (u *userService) Update(usr *pojo.User) error {
	return u.dao.Update(usr)
}
func( u *userService)GetByPhoneAndPassword(phoneNum,password string)*pojo.User{
	return u.dao.GetByPhoneAndPassword(phoneNum,password)
}