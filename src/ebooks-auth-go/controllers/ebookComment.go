package controllers

import (
	"ebooks-auth-go/models"
	"fmt"

	"github.com/astaxie/beego"
)

type EbookCommentIdController struct {
	beego.Controller
}

func (u *EbookCommentIdController) GetEbookCommentByIdC() {
	// fmt.Println("???")
	// id, _ := u.GetString(":id")
	// id, _ := u.Ctx.Input.Param(":id")
	id := u.GetString("id")
	// u.Ctx.WriteString(id)
	fmt.Println(id, "xxx")
	s := models.GetEbookCommentById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
