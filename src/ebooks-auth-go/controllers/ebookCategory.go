package controllers

import (
	"ebooks-auth-go/models"

	"github.com/astaxie/beego"
)

type EbookCategoryController struct {
	beego.Controller
}

func (u *EbookCategoryController) Get() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllEbookCategory()
	u.Data["json"] = ss
	u.ServeJSON()

}
