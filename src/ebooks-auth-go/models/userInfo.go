package models

// import (
// 	"github.com/astaxie/beego/orm"
// )

// type UserInfo struct {
// 	Id            int    `json:id`
// 	User_id       string `json:userId`
// 	User_realname string `json:userRealName`
// 	User_number   string `json:userNumber`
// 	User_phone    string `json:userPhone`
// 	User_sex      string `json:userSex`
// 	User_img      string `json:userImg`
// }

// func UserInfoUpdate(username, password string) string {
// 	// 插入数据库（数据库表，Users）
// 	//获取orm对象
// 	o := orm.NewOrm()
// 	//   获取插入对象
// 	var user UserMsg
// 	//   插入数值
// 	user.User_name = username
// 	user.User_psw = password

// 	_, err := o.Insert(&user)
// 	if err != nil {
// 		// beego.Info("插入数据失败，用户相同或者其他错误！！！")
// 		// this.TplName = "register.html"
// 		// u.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
// 		return "插入数据失败"
// 	}
// 	return "success"
// }
// func init() {
// 	// 需要在init中注册定义的model
// 	orm.RegisterModel(new(UserInfo))
// }
