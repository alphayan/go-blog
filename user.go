package main

import "github.com/sirupsen/logrus"

type User struct {
	ID       int    `json:"id" gorm:"id;AUTO_INCREMENT" form:"-"`                     //主键
	Username string `json:"username" gorm:"username;unique;not null" form:"username"` //用户名
	Password string `json:"password" gorm:"password" form:"password"`                 //密码
	Status   int    `json:"status" gorm:"status default:0" form:"status"`             //账号状态 0未审核 1已审核 2已注销
	Model
}

func (u *User) post() error {
	return db.Create(u).Error
}
func (u *User) get() error {
	return db.First(u).Error
}
func (u *User) gets() interface{} {
	us := []User{}
	if err := db.Find(&us).Error; err != nil {
		logrus.Error(err)
	}
	return us
}
func (u *User) delete() error {
	return db.Delete(u).Error
}
func (u *User) put() error {
	return db.Updates(u).Error
}

func (u *User) login() error {
	return db.Model(u).Where("username=? and password=?", u.Username, u.Password).Error
}
