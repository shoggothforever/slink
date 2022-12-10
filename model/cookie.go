package model

import "time"

//use for login
type Cookie struct {
	UserId    int    `gorm:"type:int;primarykey" form:"user_id" json:"user_id"`
	Jwt       string `gotm:"not null" form:"jwt" json:"jwt"`
	CreatedAt time.Time
}
