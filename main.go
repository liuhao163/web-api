package main

import (
	"github.com/astaxie/beego"
	_ "web-api/routers"
	"web-api/utils"
)

func main() {
	DbUtils.Init()
	beego.Run()
}
