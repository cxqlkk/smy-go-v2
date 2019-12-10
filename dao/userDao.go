package dao

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/iris-contrib/go.uuid"
	"smyappTwo/pojo"
	"strings"
)

type UserDao interface {
	Create(user *pojo.User) error
	GetById(id string) *pojo.User
	List(user *pojo.User, pageNo, pageSize int) []pojo.User
	GetByPhoneAndPassword(phoneNum,password string)*pojo.User
	Count(user *pojo.User) int64
	Delete(id string) error
	Update(usr *pojo.User) error
}

type userDao struct {
	engin *xorm.Engine
}

func NewUserDao(e *xorm.Engine) UserDao {
	fmt.Println("userDao ...........new")
	return &userDao{e}

}

func (u *userDao) Create(user *pojo.User) error {
	if ""==user.Id{
		uid,_:=uuid.NewV4()
		user.Id=strings.ReplaceAll(uid.String(), "-", "")
	}
	user.Status=1
	_, e := u.engin.InsertOne(user)
	return e
}

//关于 0 值和字符串的问题
func (u *userDao) GetById(id string) *pojo.User {
	usr := &pojo.User{Id: id}
	ok, e := u.engin.Where("status=1").Get(usr) // id where 加上 usr 里的 非空 非零 字段
	//get 条件  默认也是忽略 空字符和 0 字段 用 where 来处理这些特殊的字段
	// get 里用 cols 是代表选择要查询的字段
	fmt.Println(ok, e)
	if ok && e == nil {
		return usr
	}
	return nil
}

func (u *userDao) Delete(id string) error {
	// update 默认是 忽略non-empty and non-zero fields  ,cols 用来 打破这一规则
	n, e := u.engin.ID(id).Cols("status").Update(&pojo.User{Status: 0})
	if n != 0 && e == nil {
		return nil
	}
	return errors.New("删除失败")
}

func (u *userDao) List(user *pojo.User, pageNo, pageSize int) ([]pojo.User) {
	usrs := []pojo.User{}
	err := u.engin.Where("status =1 and userName like ?", "%"+user.UserName+"%").Limit(pageSize, (pageNo-1)*pageSize).Find(&usrs) //注意 limit 用法
	// limit 5,6 从第6条开始，取6条数据
	// limit 6  offset 5 从第6 条开始，取 6 条数据
	if err != nil {
		return nil
	}
	return usrs
}
func (u *userDao) Count(user *pojo.User) int64 {
	n, e := u.engin.Table( /*"tb_user"*/ &pojo.User{}).Where("status =1 and userName like ?", "%"+user.UserName+"%").Count()
	if e != nil {
		return 0
	}
	return n
}

func (u *userDao) Update(usr *pojo.User) error {
	n, e := u.engin.ID(usr.Id).Update(usr) // 默认只更新 usr 里不为空 或者0 的字段， 如果 前台真的修改为0 或者空的时候就会出现修改不掉的情况 必须使用Allcols() 或者cols
	if n == 0 || e != nil {
		return errors.New("更新失败")
	}
	return nil
}
//登陆使用
func (u *userDao)GetByPhoneAndPassword(phoneNum,password string)*pojo.User{
	usr:=&pojo.User{}
	has,e:=u.engin.Omit("password").Where("phoneNum=? and password=? and status=1",phoneNum,password).Get(usr)
	if has&&e==nil{
		return usr

	}
	return nil
}
