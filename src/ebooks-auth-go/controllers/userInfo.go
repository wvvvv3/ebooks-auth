package controllers

import (
	"github.com/astaxie/beego"
)

type UserInfoController struct {
	beego.Controller
}

func (u *UserMsgController) UserRegisterC() {
	// username := u.GetString("username")
	// password := u.GetString("password")
	// if models.UserRegister(username, password) {
	// 	u.Data["json"] = "login success"
	// } else {
	// 	u.Data["json"] = "login fail"
	// }
	u.ServeJSON()
}
