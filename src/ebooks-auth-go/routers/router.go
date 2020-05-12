// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ebooks-auth-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	beego.Router("/ebook/student", &controllers.StudentController{})
	beego.Router("/ebook/editors", &controllers.EbookEditorsController{}, "get:GetAllEbookEditorsC")
	beego.Router("/ebook/list", &controllers.EbookListController{}, "get:GetAllEbookListC")
	beego.Router("/ebook/addlist", &controllers.EbookListController{}, "post:AddEbookListC")
	beego.Router("/ebook/dellist", &controllers.EbookListController{}, "get:DelEbookListC")
	beego.Router("/ebook/updatelist", &controllers.EbookListController{}, "post:UpdateEbookListC")
	beego.Router("/ebook/getlist/?:id", &controllers.EbookListController{}, "get:GetEbookListByIdC")
	beego.Router("/ebook/category", &controllers.EbookCategoryController{}, "get:GetAllEbookCategoryC")
	beego.Router("/ebook/getcategory/?:id", &controllers.EbookCategoryController{}, "get:GetEbookCategoryByIdC")
	beego.Router("/ebook/msg", &controllers.EbookMsgController{}, "get:GetAllEbookMsgC")
	beego.Router("/ebook/getmsg/?:id", &controllers.EbookMsgController{}, "get:GetEbookMsgByIdC")
	beego.Router("/ebook/addbook", &controllers.EbookMsgController{}, "post:AddEbookMsgC")
	beego.Router("/ebook/delbook", &controllers.EbookMsgController{}, "get:DelEbookMsgC")
	beego.Router("/ebook/updatebook", &controllers.EbookMsgController{}, "post:UpdateEbookMsgC")
	beego.Router("/ebook/getcomment/?:id", &controllers.EbookCommentIdController{}, "get:GetEbookCommentByIdC")
	beego.Router("/ebook/login/?:username/?:password", &controllers.UserMsgController{}, "get:UserLoginC")
	beego.Router("/ebook/logon/", &controllers.UserMsgController{}, "post:UserUpdateC")
	beego.Router("/ebook/usermsg/", &controllers.UserMsgController{}, "get:GetAllUserMsgC")
	beego.Router("/ebook/userinfo/", &controllers.UserMsgController{}, "get:GetAllUserInfoC")
	beego.Router("/ebook/deluser/", &controllers.UserMsgController{}, "get:DelUserMsgC")
	beego.Router("/ebook/updateuser/", &controllers.UserMsgController{}, "post:UpdateUserMsgC")
	beego.Router("/ebook/userlevel/", &controllers.UserMsgController{}, "get:GetUserLevelC")
	beego.Router("/ebook/getusermsg/", &controllers.UserMsgController{}, "get:GetUserMsgByIdC")
	beego.Router("/ebook/getuserinfo/", &controllers.UserMsgController{}, "get:GetUserInfoByIdC")
	beego.Router("/ebook/getcart/?:id", &controllers.OrderCartController{}, "get:GetOrderCartByIdC")
	beego.Router("/ebook/delcart/", &controllers.OrderCartController{}, "post:OrderCartUpdateByIdC")
	beego.Router("/ebook/getordermsg/", &controllers.OrderMsgController{}, "get:GetAllOrderMsgC")
	beego.Router("/ebook/updateorder/", &controllers.OrderMsgController{}, "post:OrderUpdateC")
	beego.Router("/ebook/getorderdetails/", &controllers.OrderMsgController{}, "get:GetAllOrderDetailsC")

}
