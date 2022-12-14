package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

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
	Id      int `gorm:"type:uint;primaryKey autoincrement " form:"id" json:"id"`
	UserId  int `gorm:"type:varchar(10);column:user_id" form:"user_id" json:"user_id"`
	LoginAt time.Time
}

//UserTable
type User struct {
	Id        int    `gorm:"type:uint;primaryKey autoincrement" form:"id" json:"id"`
	Name      string `gorm:"type:varchar(40)" form:"name"`
	Email     string `gorm:"type:varchar(40)" form:"email"`
	Pwd       string `gorm:"type:varchar(40)" form:"pwd"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Url       []UrlInfo `gorm:"foreignKey:UserId "`
}

//用于存储暂停的url

type PauseUrl struct {
	UserId int    `gorm:"type:int" form:"user_id" json:"user_id"`
	UrlId  int    `gorm:"type:int" form:"url_id" json:"url_id"`
	Short  string `gorm:"type:varchar(50)" form:"short" json:"short"`
}

//用于存储验证信息的表
type Claims struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var AuthClaims *Claims
var AuthJwt string

/*
记录当前登录用户信息
*/
var DefaultUser User
var NOTLOGIN int = -1

func (u User) TableName() string {
	return "users"
}
func (u User) GetId() int {
	return u.Id
}
func (u UrlInfo) TableName() string {
	return "url_infos"
}
