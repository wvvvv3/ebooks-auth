package controllers

import (
	"ebooks-auth-go/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type OrderCartController struct {
	beego.Controller
}

func (u *OrderCartController) GetOrderCartByIdC() {
	id := u.GetString("id")
	s := models.GetOrderCartById(id)
	u.Data["json"] = s
	u.ServeJSON()
}

func (u *OrderCartController) OrderCartUpdateByIdC() {
	var k models.OrderCart

	json.Unmarshal(u.Ctx.Input.RequestBody, &k)
	// fmt.Println(k, "zzzz")
	t := models.OrderCartUpdate(&k)
	u.Data["json"] = t
	u.ServeJSON()
}
