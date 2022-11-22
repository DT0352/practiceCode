package main

import "fuvidy/internal"

func main() {
	agentInfo := internal.GetAgentInfo(5)
	println(agentInfo)
}
