package model

import "time"

//use for login
type LoginInfo struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}

//use for create and update and relative operations
type UrlInfo struct {
	Origin     string
	Short      string
	Comment    string
	StartTime  time.Time
	ExpireTimr time.Time
}
