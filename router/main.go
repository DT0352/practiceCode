package main

import "fuvidy/internal"

func main() {
	for _, agent := range internal.GetOnlineAgent() {
		println(agent.UserName)
	}
	internal.GetAgentInfo([]int{1})
}
