package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type OrderCart struct {
	Id       int    `json:id`
	User_id  int    `json:user_id`
	Ebook_id string `json:ebook_id`
}

func GetOrderCartById(id string) string {
	v, errv := strconv.Atoi(id)
	//获取orm对象
	o := orm.NewOrm()
	o.Using("default")
	//获取查询对象
	var odercart OrderCart
	// 查询
	odercart.User_id = v
	err := o.Read(&odercart, "User_id")

	if err != nil {
		return "error1"
	}
	if errv != nil {
		return "error2"
	}
	return odercart.Ebook_id
}
func OrderCartUpdate(ordercart *OrderCart) string {
	o := orm.NewOrm()
	o.Using("default")
	//获取查询对象
	var order OrderCart
	// 查询
	order.User_id = ordercart.User_id
	err := o.Read(&order, "User_id")

	if err != nil {
		return "error2"
	} else {
		order.Ebook_id = ordercart.Ebook_id
		// fmt.Println(order, "zzzz")
		_, errv := o.Update(&order, "Ebook_id")
		if errv != nil {
			return "error1"
		}
		return "success"
	}
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(OrderCart))
}
