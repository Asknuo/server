package main

import (
	"sever/core"
	"sever/flag"
	"sever/global"
	"sever/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	//global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectEs()

	defer global.Redis.Close()
	flag.InitFlag()
	initialize.InitCron()

	core.RunServer()
}
