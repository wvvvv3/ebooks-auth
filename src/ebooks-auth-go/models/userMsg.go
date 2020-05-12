package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type UserMsg struct {
	Id         int64  `json:id`
	User_name  string `json:username`
	User_psw   string `json:userpsw`
	User_level int64  `json:userlevel`
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

func GetUserLevel(id string) int64 {
	v, errv := strconv.ParseInt(id, 10, 64)
	//获取orm对象
	o := orm.NewOrm()
	o.Using("default")
	//获取查询对象
	var user UserMsg
	// 查询
	user.Id = v
	err := o.Read(&user)

	if err != nil {
		fmt.Println("用户名登录失败！！")
		return -1
	}
	if errv != nil {
		fmt.Println("id获取错误")
	}
	return user.User_level
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

// func UserMsgGetAll(userId string) (UserMsg, UserInfo) {
// 	//获取orm对象
// 	o := orm.NewOrm()
// 	o.Using("default")

// 	v, errv := strconv.ParseInt(userId, 10, 64)

// 	//获取查询对象
// 	var usermsg UserMsg
// 	var userinfo UserInfo
// 	// 查询
// 	usermsg.Id = v
// 	userinfo.Id = v

// 	err := o.Read(&usermsg, "Id")
// 	err2 := o.Read(&userinfo, "Id")

// 	if err != nil || err2 != nil || errv != nil {
// 		fmt.Println("error")
// 	}

// 	return usermsg, userinfo
// }
// 查询所有用户信息
func GetAllUserMsg() []*UserMsg {
	// fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var userMsg []*UserMsg
	q := o.QueryTable("user_msg")
	q.All(&userMsg)
	return userMsg

}

// 查询所有用户信息
func GetAllUserInfo() []*UserInfo {
	// fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var userInfo []*UserInfo
	q := o.QueryTable("user_info")
	q.All(&userInfo)
	return userInfo

}

// 删除用户基本信息
func DelUserMsg(id string) string {
	v, errv := strconv.ParseInt(id, 10, 64)

	u := UserMsg{Id: v}
	k := UserInfo{User_id: v}
	o := orm.NewOrm()
	o.Using("default")
	// fmt.Println("daada", id)
	_, err := o.Delete(&u)
	_, err1 := o.Delete(&k, "User_id")
	if err != nil {
		fmt.Println("删除失败", err)
		return "error"
	}
	if err1 != nil {
		fmt.Println("删除失败", err)
		return "error"
	}
	if errv != nil {
		fmt.Println("id获取错误")
	}
	return "seccess"
}

// 更新用户信息
func UpdateUserMsg(userMsg *UserMsg, userInfo *UserInfo) string {
	o := orm.NewOrm()
	o.Using("default")

	msg := UserMsg{}

	msg.Id = userMsg.Id
	err := o.Read(&msg)
	if err == nil {
		msg = *userMsg
		_, err := o.Update(&msg)
		if err != nil {
			return "error"
		}
		// return "seccess"
	}

	info := UserInfo{}
	info.User_id = userInfo.User_id

	err1 := o.Read(&info, "User_id")

	if err1 == nil {
		info1 := *userInfo
		_, err1 := o.Update(&info1)
		// fmt.Println(a, "dadad", info1)
		if err1 != nil {
			return "error"
		}
		// return "seccess"
	}
	return "seccess"

}

// 根据用户id查询用户信息
func GetUserMsgById(id string) UserMsg {
	// string转换成int
	v, errv := strconv.ParseInt(id, 10, 64)
	u := UserMsg{Id: v}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	if errv != nil {
		fmt.Println("id获取错误")
	}
	return u
}
func GetUserInfoById(id string) UserInfo {
	// string转换成int
	v, errv := strconv.ParseInt(id, 10, 64)
	u := UserInfo{User_id: v}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u, "User_id")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	if errv != nil {
		fmt.Println("id获取错误")
	}
	return u
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(UserMsg))
	orm.RegisterModel(new(UserInfo))
}
