package model

//common response
type Response struct {
	Code int    `form:"code" json:"code"`
	Msg  string `form:"msg" json:"msg"`
	//Data interface{} `form:"data" json:"data"` //根据要求传入具体的响应
}

//content the list of urls,use for get function
type GetResponse struct {
	Response
	Urls []UrlInfo `form:"urls" json:"urls"`
}

//return user's id
type CreateResponse struct {
	Response
	Id string `form:"id"`
}
type QueryResponse struct {
	Response
	url UrlInfo `form:"url" json:"url"`
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
