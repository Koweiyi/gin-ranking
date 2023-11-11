package main

import (
	"go-ranking/router"
)

func main(){
	r := routes.GetRoute()
	r.Run(":9999")
}