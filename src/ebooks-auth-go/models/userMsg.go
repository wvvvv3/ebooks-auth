package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type UserMsg struct {
	Id        int64  `json:id`
	User_name string `json:username`
	User_psw  string `json:userpsw`
}

type UserInfo struct {
	Id            int64  `json:id`
	User_id       int64  `json:userId`
	User_realname string `json:userRealName`
	User_number   string `json:userNumber`
	User_phone    string `json:userPhone`
	User_sex      string `json:userSex`
	User_img      string `json:userImg`
}

func UserLogin(username, password string) int64 {
	//获取orm对象
	o := orm.NewOrm()
	o.Using("default")
	//获取查询对象
	var user UserMsg
	// 查询
	user.User_name = username
	err := o.Read(&user, "User_name")

	if err != nil {
		fmt.Println("用户名登录失败！！")
		return -1
	}
	// 判断密码是否一致
	if user.User_psw != password {
		fmt.Println("密码登录失败！！！")
		return -1
	}

	return user.Id
}
func UserMsgUpdate(userMsg *UserMsg, userInfo *UserInfo) int64 {
	o := orm.NewOrm()
	o.Using("default")

	k, err := o.Insert(userMsg)

	userInfo.User_id = k

	_, errs := o.Insert(userInfo)
	if err != nil {
		fmt.Println("失败！！！")
	}
	if errs != nil {
		fmt.Println("失败！！！")
	}
	return userMsg.Id
}
func UserMsgGetAll(userId string) (UserMsg, UserInfo) {
	//获取orm对象
	o := orm.NewOrm()
	o.Using("default")

	v, errv := strconv.ParseInt(userId, 10, 64)

	//获取查询对象
	var usermsg UserMsg
	var userinfo UserInfo
	// 查询
	usermsg.Id = v
	userinfo.Id = v

	err := o.Read(&usermsg, "Id")
	err2 := o.Read(&userinfo, "Id")

	if err != nil || err2 != nil || errv != nil {
		fmt.Println("error")
	}

	return usermsg, userinfo
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(UserMsg))
	orm.RegisterModel(new(UserInfo))
}
