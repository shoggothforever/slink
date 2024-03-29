package dal

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shortlink/api/utils"
	"sync"
)

var Db *gorm.DB
var Lock sync.Mutex

func Getdb() *gorm.DB {
	return Db
}

func init() {
	//Configure the viper
	Config := viper.New()
	Config.SetConfigFile("/www/wwwroot/slink.shoggothy.xyz/config/config.yaml")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./config")
	err := Config.ReadInConfig() // 查找并读取配置文件
	if err != nil {              // 处理读取配置文件的错误
		logrus.Error("Read config file failed: %s \n", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("no error in config file")
		} else {
			logrus.Error("found error in config file\n", ok)
		}
	}
	utils.Salt = Config.GetString("Salt")
	jwtInfo := Config.GetStringMapString("Jwt")
	utils.JwtSecret = jwtInfo["secret"]
	loginInfo := Config.GetStringMapString("mysql")
	Dsn := loginInfo["predsn"] + loginInfo["database"] + loginInfo["mode"]

	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       Dsn, //"link:link@tcp(localhost:3306)/slink?charset=utf8&parseTime=true&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		SkipDefaultTransaction: true})
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("gorm OPENS MySQL failed")
	}
	//err = Getdb().AutoMigrate(&model.User{}, &model.UrlInfo{}, &model.LoginInfo{}, &model.Cookie{}, &model.PauseUrl{})
	//if err != nil {
	//	logrus.Error("build tables corrupt!\n", err)
	//}
	//model.DefaultUser.Id = model.NOTLOGIN
	//model.DefaultUser.Name = "origin"

}

/*
* @brief init the config of viper and database
set default viper config path and read data from config.yaml
to get mysql userInfo,jwt secretKey
open mysql database with gorm and create table with user Struct
*/
