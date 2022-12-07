package model

import "time"

//use for login
type LoginInfo struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}

//use for record get
type LoginRecord struct {
	Time   time.Time
	Ip     string
	Status string
}

//use for create and update
type UrlInfo struct {
	Origin     string
	Short      string
	Comment    string
	StartTime  time.Time
	ExpireTimr time.Time
}
