package main

import (
	"chat/router"
)

func main() {
	r := router.Router()
	r.Run(":8080")
}
