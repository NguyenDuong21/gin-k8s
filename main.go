package main

import (
	"gin-k8s/Routes"
)

var err error

func main() {

	r := Routes.SetupRouter()
	// running
	r.Run()
}
