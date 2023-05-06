package main

import (
	"ticket/goapi/httpserv"
	"ticket/goapi/infrastructure"
	"ticket/goapi/logs"
)

func init() {
	infrastructure.InitConfig()
	logs.InitLogger()
}

func main() {
	infrastructure.InitTimeZone()
	infrastructure.InitDB()
	httpserv.Run()

}
