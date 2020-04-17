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
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// 	// beego.NSNamespace("/student",
	// 	// 	beego.NSInclude(
	// 	// 		&controllers.StudentController{},
	// 	// 	),
	// 	// ),
	// )
	// beego.AddNamespace(ns)
	beego.Router("/", &controllers.UserController{})
	beego.Router("/ebook/student", &controllers.StudentController{})
	beego.Router("/ebook/editors", &controllers.EbookEditorsController{}, "get:GetAllEbookEditorsC")
	beego.Router("/ebook/list", &controllers.EbookListController{}, "get:GetAllEbookListC")
	beego.Router("/ebook/getlist/?:id", &controllers.EbookListController{}, "get:GetEbookListByIdC")
	beego.Router("/ebook/category", &controllers.EbookCategoryController{}, "get:GetAllEbookCategoryC")
	beego.Router("/ebook/getcategory/?:id", &controllers.EbookCategoryController{}, "get:GetEbookCategoryByIdC")
	beego.Router("/ebook/msg", &controllers.EbookMsgController{}, "get:GetAllEbookMsgC")
	beego.Router("/ebook/getmsg/?:id", &controllers.EbookMsgController{}, "get:GetEbookMsgByIdC")
	beego.Router("/ebook/getcomment/?:id", &controllers.EbookCommentIdController{}, "get:GetEbookCommentByIdC")
	beego.Router("/ebook/login/?:username/?:password", &controllers.UserMsgController{}, "get:UserLoginC")
	beego.Router("/ebook/logon/", &controllers.UserMsgController{}, "post:UserUpdateC")
	beego.Router("/ebook/getcart/?:id", &controllers.OrderCartController{}, "get:GetOrderCartByIdC")
	beego.Router("/ebook/delcart/", &controllers.OrderCartController{}, "post:OrderCartUpdateByIdC")

}
