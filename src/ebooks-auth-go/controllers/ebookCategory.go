package controllers

import (
	"ebooks-auth-go/models"

	"github.com/astaxie/beego"
)

type EbookCategoryController struct {
	beego.Controller
}

func (u *EbookCategoryController) GetAllEbookCategoryC() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllEbookCategory()
	u.Data["json"] = ss
	u.ServeJSON()

}

func (u *EbookCategoryController) GetEbookCategoryByIdC() {
	id := u.GetString("id")
	s := models.GetEbookCategoryById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
