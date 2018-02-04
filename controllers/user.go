package controllers

import (
	"github.com/astaxie/beego"
	"web-api/models/user"
	"web-api/models/exception"
	"fmt"
	"web-api/utils"
)

type UserController struct {
	beego.Controller
}

func (uc *UserController) Prepare() {
	fmt.Println("start process")
}

func (uc *UserController) GetUser() {
	id, error := uc.GetInt64("id")
	if error != nil {
		uc.Data["json"] = exception.PARAM_ILLEGL
		uc.ServeJSON()
		return
	}

	stmt, err := DbUtils.DB.Prepare("SELECT * FROM user where  id=?")
	if err != nil {
		fmt.Print(err)
		panic(err)
		return
	}
	rows, err := stmt.Query(id)
	defer rows.Close()
	defer stmt.Close()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	users := make([]user.User, 0)

	for rows.Next() {
		var rowsUser=user.User{}
		rows.Scan(&rowsUser.Id, &rowsUser.Name);
		users = append(users, rowsUser)
	}
	uc.Data["json"] = users
	uc.ServeJSON()
}

func (this *UserController) Finish() {
	fmt.Println("finish process")
}
