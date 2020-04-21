package controllers

import (
	"ebooks-auth-go/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type UserMsgController struct {
	beego.Controller
}

// 用户登录
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

	var k models.UserMsg
	var f models.UserInfo
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

// 获取用户全部信息（登录信息+详细信息）
// 注册用户名及密码
// func (u *UserMsgController) UserGetAllC() {

// var k models.UserMsg
// var f models.UserInfo
// json.Unmarshal(u.Ctx.Input.RequestBody, &k)
// json.Unmarshal(u.Ctx.Input.RequestBody, &f)
// userId := u.GetString("id")

// t, a := models.UserMsgGetAll(userId)
// var m map[string]interface{}
// json.Unmarshal(t, &m)
// json.Unmarshal(a, &m)

// res, _ := json.Marshal(m)
// 	u.Data["json"] = t
// 	u.ServeJSON()

// }
