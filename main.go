package main

import (
	"mini_project/config"
	"mini_project/routes"
)

func main() {
	router := routes.Init()
	config.InitDB()
	router.Logger.Fatal(router.Start(":1122"))
}
