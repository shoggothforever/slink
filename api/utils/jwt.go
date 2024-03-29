package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"shortlink/model"
	"time"
)

var JwtSecret string
var Salt string
var ExpireTime time.Duration

/*
传入用户的id，姓名作为生成jwt的参数
*/

func GenerateJwt(user_id int, name string) (string, error) {
	nowTime := time.Now().In(time.Local)
	expireTime := nowTime.Add(ExpireTime)
	claims := model.Claims{
		user_id,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    model.DefaultUser.Name,
			Subject:   "login jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(JwtSecret))
	return token, err
}

func ParseToken(token string) (*model.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
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

/*
将jwt对应数据存储到本地数据库中
*/

func Messagedigest5(s string) string {
	data := md5.Sum([]byte(s + Salt))
	return fmt.Sprintf("%x", data)
}
