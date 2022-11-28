package internal

import (
	"fuvidy/internal/util"
	"gorm.io/gorm"
)

type Agent struct {
	UserId            int `gorm:"primary_key"`
	UserName          string
	MaxReception      int
	CurrencyReception int
	Area              string
	LoginTime         int64
	Skill             string
	OverflowSkill     string
	OnLine            int
}

var db *gorm.DB

func init() {
	db = util.GetConn()
}
func (Agent) TableName() string {
	return "agents"
}

// 获取坐席信息
func GetAgentInfo(userId []int) []*Agent {
	var agent []*Agent
	if len(userId) == 0 {
		db.Find(&agent)
	} else if len(userId) == 1 {
		db.Where("user_id = ?", userId[0]).First(&agent)
	} else {
		db.Find(&agent, userId)
	}
	return agent
}

//获取在线坐席信息
func GetOnlineAgent(skill string) []*Agent {
	var agent []*Agent
	db.Where("online = 1").Where("skill = ?", skill).Find(&agent)
	return agent
}

//坐席上线
//db.Model(&User{}).Where("active = ?", true).Update("name", "hello")

func AgentOnline(userId int) bool {
	result := db.Model(&Agent{}).Where("user_id = ?", userId).Update("online", "1")
	return result.RowsAffected == 1
}

//坐席下线
func AgentOffline(userId int) bool {
	result := db.Model(&Agent{}).Where("user_id = ?", userId).Update("online", "0")
	return result.RowsAffected == 1
}
