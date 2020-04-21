package routers

import (
	"house-auth/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/house/login/?:username/?:password", &controllers.UserMsgController{}, "get:UserLoginC")
	beego.Router("/house/logon/", &controllers.UserMsgController{}, "post:UserUpdateC")
}
