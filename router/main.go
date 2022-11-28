package main

import (
	"fuvidy/internal"
	"time"
)

func main() {
	consumer := internal.GetConsumer("咨询")
	print("id:", consumer.Id, "\tskill:", consumer.Skill)
	print(time.Now().Unix())
	loginMap := internal.Matching(consumer)
	for k, v := range loginMap {
		println("坐席号:", k, "\t登录时长:", v)
	}

}
