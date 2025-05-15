package main

import (
	"sever/core"
	"sever/global"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()

	core.RunServer()
}
