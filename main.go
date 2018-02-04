package main

import (
	_ "web-api/routers"
	"github.com/astaxie/beego"
	"web-api/utils"
)

func main() {
	DbUtils.Init()
	beego.Run()
}
