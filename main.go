package main

import "github.com/basic-echo-golang/routes"


func main() {
	e := routes.Init()
	
	e.Logger.Fatal(e.Start(":1234"))
}