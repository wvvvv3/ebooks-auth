package controllers

import (
	"ebooks-auth-go/models"
	"fmt"

	"github.com/astaxie/beego"
)

type EbookMsgController struct {
	beego.Controller
}

func (u *EbookMsgController) GetAllEbookMsgC() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllEbookMsg()
	u.Data["json"] = ss
	u.ServeJSON()

}
func (u *EbookMsgController) GetEbookMsgByIdC() {
	// fmt.Println("???")
	// id, _ := u.GetString(":id")
	// id, _ := u.Ctx.Input.Param(":id")
	id := u.GetString("id")
	// u.Ctx.WriteString(id)
	fmt.Println(id, "xxx")
	s := models.GetEbookMsgById(id)
	u.Data["json"] = s
	u.ServeJSON()
}

// func (u *EbookMsgController) GetById() {
// 	fmt.Println("???")
// 	id, _ := u.GetInt(":id")
// 	fmt.Println(id)
// 	s := models.GetEbookMsgById(id)
// 	u.Data["json"] = s
// 	u.ServeJSON()
// }
