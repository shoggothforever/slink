package model

import "time"

type Store interface {
	//读入长url并获取短url的外部接口,返回短url字符串
	Set(lurl string) string
	//读入长url，获取短url的内部实现方法
	set(lurl, surl string) bool
	//生成短url的函数
	genshort(lurl string) string
}

//use for create and update and relative operations
type UrlInfo struct {
	Id         int       `gorm:"type:uint;primaryKey autoincrement" form:"id" json:"id"`
	UserId     int       `gorm:"type:varchar(10) column:user_id" form:"user_id" json:"user_id"`
	Origin     string    `gorm:"type:varchar(200)" form:"origin" json:"origin"`
	Short      string    `gorm:"type:varchar(40)" form:"short" json:"short"`
	Comment    string    `gorm:"type:varchar(100)" form:"comment" json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime"`
	ExpireTime time.Time `gorm:"type:datetime"`
}
type LoginInfo struct {
	Id      int `gorm:"type:uint;primaryKey autoincrement" form:"id" json:"id"`
	UserId  int `gorm:"type:varchar(10);column:user_id" form:"user_id" json:"user_id"`
	LoginAt time.Time
}

//UserTable
type User struct {
	Id        int    `gorm:"type:uint;primaryKey autoincrement" form:"id" json:"id"`
	Name      string `gorm:"type:varchar(40) " form:"name"`
	Email     string `gorm:"type:varchar(40) " form:"email"`
	Pwd       string `gorm:"type:varchar(40) " form:"pwd"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Url       UrlInfo `gorm:"foreignKey:UserId "`
}

var CurrentUser User

func (u User) TableName() string {
	return "users"
}
func (u User) GetId() int {
	return u.Id
}
