package internal

import "time"

func Matching(consumer *Consumer) map[int]int64 {
	agents := GetOnlineAgent(consumer.Skill)
	//	计算登录时长
	var loginTimeMap = make(map[int]int64)
	now := time.Now().Unix()
	for _, agent := range agents {
		loginTimeMap[agent.UserId] = now - agent.LoginTime
	}
	return loginTimeMap
}

func loginTimeMatching(agents []*Agent) map[int]int64 {
	var loginTimeMap = make(map[int]int64)
	var loginTimeResultMap = make(map[int]int64)
	now := time.Now().Unix()
	sumTime := int64(0)
	for _, agent := range agents {
		loginTimeMap[agent.UserId] = now - agent.LoginTime
		sumTime += loginTimeMap[agent.UserId]
	}
	for userId, loginTime := range loginTimeMap {
		loginTimeResultMap[userId] = loginTime / sumTime
	}
	return loginTimeResultMap
}
func receptionMatching(agents []*Agent) map[int]int {
	var receptionMap = make(map[int]int)
	var receptionResultMap = make(map[int]int)
	sumReception := 0
	for _, agent := range agents {
		receptionMap[agent.UserId] = agent.CurrencyReception
		sumReception += receptionMap[agent.UserId]
	}
	for userId, reception := range receptionMap {
		receptionResultMap[userId] = reception / sumReception
	}
	return receptionResultMap
}
