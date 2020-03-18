package controllers

import (
	"ebooks-auth-go/models"

	"github.com/astaxie/beego"
)

type EbookMsgController struct {
	beego.Controller
}

func (u *EbookMsgController) Get() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllEbookMsg()
	u.Data["json"] = ss
	u.ServeJSON()

}
