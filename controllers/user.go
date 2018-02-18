package controllers

import (
	"github.com/astaxie/beego"
	"web-api/models/user"
	"web-api/service"
	"web-api/models/exception"
)

type UserController struct {
	beego.Controller
}

func (uc *UserController) PostUser() {
	name := uc.GetString("name")
	insertId := UserService.Save(&user.User{0, name})
	uc.Data["json"] = insertId
	uc.ServeJSON()
	return
}

func (uc *UserController) GetUser() {
	id, error := uc.GetInt64("id")
	if error != nil {
		uc.Data["json"] = exception.PARAM_ILLEGL
		uc.ServeJSON()
		return
	}

	uc.Data["json"] = UserService.GetUserById(id)
	uc.ServeJSON()
}
