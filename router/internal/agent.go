package internal

import (
	"fuvidy/internal/util"
	"github.com/jinzhu/gorm"
	"time"
)

type Agent struct {
	UserId        int `gorm:"primary_key"`
	UserName      string
	MaxReception  int
	Area          string
	LoginTime     time.Time
	Skill         string
	OverflowSkill string
	OnLine        int
}

var db *gorm.DB

func init() {
	db = util.GetConn()
}

// 获取坐席信息
func GetAgentInfo(userId int) *Agent {
	var agent = &Agent{}
	db.Where("user_id = ?", userId).First(&agent)
	return agent
}
