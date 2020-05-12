package controllers

import (
	"ebooks-auth-go/models"
	"encoding/json"
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

// 通过id查询书单信息
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

// 添加书单
func (u *EbookListController) AddEbookListC() {
	// fmt.Println("models is inited!co")
	var k models.EbookList

	json.Unmarshal(u.Ctx.Input.RequestBody, &k)

	t := models.AddEbookList(&k)

	u.Data["json"] = t
	u.ServeJSON()

}

// 删除书单
func (u *EbookListController) DelEbookListC() {
	id := u.GetString("id")

	s := models.DelEbookList(id)
	u.Data["json"] = s
	u.ServeJSON()

}

// 更新书单

func (u *EbookListController) UpdateEbookListC() {
	var k models.EbookList

	json.Unmarshal(u.Ctx.Input.RequestBody, &k)

	s := models.UpdateEbookList(&k)
	u.Data["json"] = s
	u.ServeJSON()

}
