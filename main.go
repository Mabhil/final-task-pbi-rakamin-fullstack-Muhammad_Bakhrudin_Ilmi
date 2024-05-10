package main

import (
	"rakamin-submission/router"
)

func main() {
	r := router.StartRoute()
	r.Run(":8080")
}
