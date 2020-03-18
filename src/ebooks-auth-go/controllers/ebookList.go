package controllers

import (
	"ebooks-auth-go/models"

	"github.com/astaxie/beego"
)

type EbookListController struct {
	beego.Controller
}

// @Title 获得所有学生
// @Description 返回所有的学生数据
// @Success 200 {object} models.Student
// @router / [get]
func (u *EbookListController) Get() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllEbookList()
	u.Data["json"] = ss
	u.ServeJSON()

}
