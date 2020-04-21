package controllers

import (
	"encoding/json"
	"house-auth/models"

	"github.com/astaxie/beego"
)

type UserMsgController struct {
	beego.Controller
}

func (u *UserMsgController) UserLoginC() {
	username := u.GetString("username")
	password := u.GetString("password")
	k := models.UserLogin(username, password)
	u.Data["json"] = k
	// if models.UserLogin(username, password) {
	// 	u.Data["json"] = "login success"
	// } else {
	// 	u.Data["json"] = "login fail"
	// }
	u.ServeJSON()
}

// 注册用户名及密码
func (u *UserMsgController) UserUpdateC() {

	var k models.VUserMsg
	var f models.VUserInfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &k)
	json.Unmarshal(u.Ctx.Input.RequestBody, &f)

	t := models.UserMsgUpdate(&k, &f)
	u.Data["json"] = t
	u.ServeJSON()

	// 数据处理
	// if username == "" || password == "" {
	// 	beego.Info("用户名或密码不能为空")
	// 	u.Data["errmsg"] = "用户名或密码不能为空 ！"
	// 	return
	// }
	// s := models.UserUpdate(username, password)
	// u.Data["json"] = s
	// u.ServeJSON()

}
