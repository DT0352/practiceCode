package internal

import (
	"math/rand"
	"strconv"
)

type Consumer struct {
	Id    string
	Skill string
}

func GetConsumer(skill string) *Consumer {
	id := strconv.Itoa(rand.Intn(9999) + 10000)[1:]
	return &Consumer{
		Id:    id,
		Skill: skill,
	}
}
