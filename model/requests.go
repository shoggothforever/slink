package model

//use for login
type LoginInfo struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}
