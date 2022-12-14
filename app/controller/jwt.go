package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"shortlink/dao"
	"shortlink/model"
	"time"
)

/*
传入用户的id，密码作为生成jwt的参数
*/
func GenerateJwt(user_id int, name string) (string, error) {
	nowTime := time.Now().In(time.Local)
	expireTime := nowTime.Add(dao.ExpireTime)
	claims := model.Claims{
		user_id,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    model.CurrentUser.Name,
			Subject:   "login jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(name))
	return token, err
}

func ParseToken(token string) (*model.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(model.CurrentUser.Name), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*model.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
func SaveJwt(id int, name string) {
	dao.ExpireTime = 4 * time.Hour
	model.AuthJwt, _ = GenerateJwt(id, name)
	model.AuthClaims, _ = ParseToken(model.AuthJwt)
	var cookie model.Cookie
	cookie.UserId = id
	cookie.Jwt = model.AuthJwt
	cookie.CreatedAt = time.Now().In(time.Local)
	var del model.Cookie
	dao.Db.Model(&cookie).Where("user_id=?", id).Delete(&del)
	dao.Db.Model(&cookie).Create(&cookie)
}
func messagedigest5(s string) string {
	data := md5.Sum([]byte(s + dao.Salt))
	return fmt.Sprintf("%x", data)
}
