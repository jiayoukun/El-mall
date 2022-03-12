package main

import (
	"os"
	"test15/conf"
	"test15/routes"
	"test15/service"
)

func main()  {
	conf.Init()
	go service.Start()
	r:=routes.NewRouter()
	r.Run(os.Getenv("HttpPort"))
}
