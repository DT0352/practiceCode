package model

import (
	"chat/util"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (user *UserBasic) TableName() string {
	return "userBasic"
}
func GetUserList() []*UserBasic {
	var user []*UserBasic
	util.DB.Find(&user)
	return user
}
