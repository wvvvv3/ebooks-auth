package controllers

import (
	"ebooks-auth-go/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type OrderMsgController struct {
	beego.Controller
}

// 获取所有订单
func (u *OrderMsgController) GetAllOrderMsgC() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllOrderMsg()
	u.Data["json"] = ss
	u.ServeJSON()
}

// 添加订单
func (u *OrderMsgController) OrderUpdateC() {

	var k models.OrderMsg
	var f models.OrderDetails
	json.Unmarshal(u.Ctx.Input.RequestBody, &k)
	json.Unmarshal(u.Ctx.Input.RequestBody, &f)

	t := models.OrderUpdate(&k, &f)
	u.Data["json"] = t
	u.ServeJSON()
}

// 获取所有详细订单
func (u *OrderMsgController) GetAllOrderDetailsC() {
	// fmt.Println("models is inited!co")
	ss := models.GetAllOrderDetails()
	u.Data["json"] = ss
	u.ServeJSON()
}

// func (u *EbookMsgController) GetEbookMsgByIdC() {
// 	// fmt.Println("???")
// 	// id, _ := u.GetString(":id")
// 	// id, _ := u.Ctx.Input.Param(":id")
// 	id := u.GetString("id")
// 	// u.Ctx.WriteString(id)
// 	fmt.Println(id, "xxx")
// 	s := models.GetEbookMsgById(id)
// 	u.Data["json"] = s
// 	u.ServeJSON()
// }

// func (u *EbookMsgController) GetById() {
// 	fmt.Println("???")
// 	id, _ := u.GetInt(":id")
// 	fmt.Println(id)
// 	s := models.GetEbookMsgById(id)
// 	u.Data["json"] = s
// 	u.ServeJSON()
// }
