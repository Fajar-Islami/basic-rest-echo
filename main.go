package main

import (
	"github.com/basic-echo-golang/db"
	"github.com/basic-echo-golang/routes"
)


func main() {
	db.Init()
	e := routes.Init()
	
	e.Logger.Fatal(e.Start(":1234"))
}