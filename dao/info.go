package dao

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

var Db *gorm.DB
var JwtSecret string
var Salt string
var Lock sync.Mutex
var ExpireTime time.Duration

func Getdb() *gorm.DB {
	return Db
}

/*
* @brief init the config of viper and database
set default viper config path and read data from config.yaml
to get mysql userInfo,jwt secretKey
open mysql database with gorm and create table with user Struct
*/
