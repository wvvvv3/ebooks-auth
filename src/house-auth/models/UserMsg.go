package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type VUserMsg struct {
	Id        int64  `json:id`
	User_name string `json:username`
	User_psw  string `json:userpsw`
}

type VUserInfo struct {
	Id            int64  `json:id`
	User_id       int64  `json:userId`
	User_realname string `json:userRealName`
	User_number   string `json:userNumber`
	User_phone    string `json:userPhone`
	User_sex      string `json:userSex`
	User_img      string `json:userImg`
	Publish_times int64  `json:publishTimes`
	Top_times     int64  `json:topTimes`
}

func UserLogin(username, password string) int64 {
	//获取orm对象
	o := orm.NewOrm()
	o.Using("default")
	//获取查询对象
	var user VUserMsg
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
func UserMsgUpdate(userMsg *VUserMsg, userInfo *VUserInfo) int64 {
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

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(VUserMsg))
	orm.RegisterModel(new(VUserInfo))
}
