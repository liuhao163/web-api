package routers

import (
	"github.com/astaxie/beego"
	"web-api/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	//beego.Router("/api/user", &controllers.UserController{}, "get:GetUser")
	beego.Router("/api/user", &controllers.UserController{}, "post:PostUser")
}
