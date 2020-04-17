package controllers

import (
	"ebooks-auth-go/models"
	"fmt"

	"github.com/astaxie/beego"
)

type EbookListController struct {
	beego.Controller
}

// 获取所有书单信息
func (u *EbookListController) GetAllEbookListC() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllEbookList()
	u.Data["json"] = ss
	u.ServeJSON()

}
func (u *EbookListController) GetEbookListByIdC() {
	// fmt.Println("???")
	// id, _ := u.GetString(":id")
	// id, _ := u.Ctx.Input.Param(":id")
	id := u.GetString("id")
	// u.Ctx.WriteString(id)
	fmt.Println(id, "xxx")
	s := models.GetEbookListById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
