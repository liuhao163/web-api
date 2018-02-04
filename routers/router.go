package routers

import (
	"web-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	beego.Router("/api/user", &controllers.UserController{}, "get:GetUser")
}
