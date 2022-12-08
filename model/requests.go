package model

//use for login
type Login struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}
