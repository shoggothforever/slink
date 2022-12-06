package model

type Store interface {
	//读入长url并获取短url的外部接口,返回短url字符串
	Set(lurl string) string
	//读入长url，获取短url的内部实现方法
	set(lurl, surl string) bool
	//生成短url的函数
	genshort(lurl string) string
}
