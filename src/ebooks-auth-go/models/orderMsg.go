package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego/orm"
)

type OrderMsg struct {
	Id           int64  `json:id`
	Order_number string `json:order_number`
	User_id      int64  `json:user_id`
	Order_name   string `json:order_name`
	Province     string `json:province`
	City         string `json:city`
	District     string `json:district`
	Address      string `json:address`
	Order_price  int64  `json:order_price`
	Order_time   string `json:order_time`
	Order_phone  string `json:order_phone`
}
type OrderDetails struct {
	Id       int64  `json:id`
	Order_id int64  `json:order_id`
	Ebook_id string `json:ebook_id`
}

func GetAllOrderMsg() []*OrderMsg {
	// fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var orderMsg []*OrderMsg
	q := o.QueryTable("order_msg")
	q.All(&orderMsg)
	return orderMsg

}
func OrderUpdate(orderMsg *OrderMsg, orderDetails *OrderDetails) int64 {
	o := orm.NewOrm()
	o.Using("default")
	t := CreateCaptcha()
	orderMsg.Order_number = t
	// fmt.Println(k)
	// strconv.Itoa(k)
	k, err := o.Insert(orderMsg)

	orderDetails.Order_id = k
	fmt.Println(orderDetails)
	_, errs := o.Insert(orderDetails)
	if err != nil {
		fmt.Println("失败！！！")
	}
	if errs != nil {
		fmt.Println("失败！！！")
	}
	return orderMsg.Id
}
func CreateCaptcha() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}
func GetAllOrderDetails() []*OrderDetails {
	// fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var orderDetails []*OrderDetails
	q := o.QueryTable("order_details")
	q.All(&orderDetails)
	return orderDetails

}

// func GetEbookMsgById(id string) EbookMsg {
// 	// string转换成int
// 	v, errv := strconv.Atoi(id)
// 	u := EbookMsg{Id: v}
// 	o := orm.NewOrm()
// 	o.Using("default")
// 	err := o.Read(&u)
// 	if err == orm.ErrNoRows {
// 		fmt.Println("查询不到")
// 	} else if err == orm.ErrMissPK {
// 		fmt.Println("找不到主键")
// 	}
// 	if errv != nil {
// 		fmt.Println("id获取错误")
// 	}
// 	return u
// }

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(OrderMsg))
	orm.RegisterModel(new(OrderDetails))
}
