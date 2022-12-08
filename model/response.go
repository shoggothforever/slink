package model

import "time"

//common response
type Response struct {
	Code int    `form:"code" json:"code"`
	Msg  string `form:"msg" json:"msg"`
	//Data interface{} `form:"data" json:"data"` //根据要求传入具体的响应
}
type RegisterResponse struct {
	Response
	User User
}
type LoginResponse struct {
	Response
	Id int
}

//content the list of urls,use for get function
type InfoResponse struct {
	Response
	Id    int
	Name  string
	Email string
}

//return id of user
type CreateResponse struct {
	Response
	url UrlInfo `form:"url" json:"url"'`
}

//
type QueryResponse struct {
	Response
	Url []UrlInfo `form:"url" json:"url"`
}

//use for record get
type LoginRecord struct {
	Id      int
	LoginAt time.Time
}
type LoginInfoResponse struct {
	Response
	Records []LoginRecord
}

//只需返回基本信息，可以适当添加额外信息
type UpdateResponse struct {
	Response
}
type DeleteResponse struct {
	Response
}
type PauseResponse struct {
	Response
}
