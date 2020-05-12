package controllers

import (
	"ebooks-auth-go/models"
	"encoding/json"
	"fmt"

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

func (u *UserMsgController) GetUserLevelC() {
	id := u.GetString("id")

	k := models.GetUserLevel(id)
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

// 查询所有用户基本信息
func (u *UserMsgController) GetAllUserMsgC() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllUserMsg()
	u.Data["json"] = ss
	u.ServeJSON()

}

// 查询所有用户详细信息
func (u *UserMsgController) GetAllUserInfoC() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllUserInfo()
	u.Data["json"] = ss
	u.ServeJSON()

}

// 删除用户信息
func (u *UserMsgController) DelUserMsgC() {
	id := u.GetString("id")

	s := models.DelUserMsg(id)
	u.Data["json"] = s
	u.ServeJSON()

}

// 更新用户信息
func (u *UserMsgController) UpdateUserMsgC() {

	var k models.UserMsg
	var f models.UserInfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &k)
	json.Unmarshal(u.Ctx.Input.RequestBody, &f)
	fmt.Println(f, "sss")
	t := models.UpdateUserMsg(&k, &f)
	u.Data["json"] = t
	u.ServeJSON()
}

// 根据id查询用户信息
func (u *UserMsgController) GetUserMsgByIdC() {
	id := u.GetString("id")
	s := models.GetUserMsgById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
func (u *UserMsgController) GetUserInfoByIdC() {
	id := u.GetString("id")
	s := models.GetUserInfoById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
