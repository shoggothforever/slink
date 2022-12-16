package dao

import "gorm.io/gorm"

func Getdb() *gorm.DB {
	return Db
}
