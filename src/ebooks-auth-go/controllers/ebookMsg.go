package controllers

import (
	"ebooks-auth-go/models"
	"encoding/json"
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

// 添加书籍
func (u *EbookMsgController) AddEbookMsgC() {
	// fmt.Println("models is inited!co")
	var k models.EbookMsg

	json.Unmarshal(u.Ctx.Input.RequestBody, &k)
	fmt.Println(k, "xxx")
	t := models.AddEbookMsg(&k)

	u.Data["json"] = t
	u.ServeJSON()

}

// 删除书籍
func (u *EbookMsgController) DelEbookMsgC() {
	id := u.GetString("id")

	s := models.DelEbookMsg(id)
	u.Data["json"] = s
	u.ServeJSON()

}

// 更新书籍

func (u *EbookMsgController) UpdateEbookMsgC() {
	var k models.EbookMsg

	json.Unmarshal(u.Ctx.Input.RequestBody, &k)

	s := models.UpdateEbookMsg(&k)
	u.Data["json"] = s
	u.ServeJSON()

}
