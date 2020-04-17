package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type EbookMsg struct {
	Id           int
	Isbn         int
	Name         string
	Price        int
	Publish_date string
	Publisher    string
	Writer       string
	Page         int
	Synopsis     string
	Buy_times    int
	Pic          string
	Grade        int
	Category     string
}

func GetAllEbookMsg() []*EbookMsg {
	// fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var ebookMsg []*EbookMsg
	q := o.QueryTable("ebook_msg")
	q.All(&ebookMsg)
	return ebookMsg

}

func GetEbookMsgById(id string) EbookMsg {
	// string转换成int
	v, errv := strconv.Atoi(id)
	u := EbookMsg{Id: v}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	if errv != nil {
		fmt.Println("id获取错误")
	}
	return u
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(EbookMsg))
}
