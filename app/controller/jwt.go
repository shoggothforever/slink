package controller

import (
	"github.com/dgrijalva/jwt-go"
	"shortlink/dao"
	"shortlink/model"
	"time"
)

/*
传入用户的id，密码作为生成jwt的参数
*/
func GenerateJwt(id, pwd string) (string, error) {
	nowTime := time.Now().In(time.Local)
	expireTime := nowTime.Add(dao.ExpireTime)
	claims := model.Claims{
		id,
		pwd,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    model.CurrentUser.Name,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(model.CurrentUser.Name))
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
func SaveJwt(id, pwd string) {
	dao.ExpireTime = 4 * time.Hour
	model.AuthToken, _ = GenerateJwt(id, pwd)
	model.AuthClaims, _ = ParseToken(model.AuthToken)
}
