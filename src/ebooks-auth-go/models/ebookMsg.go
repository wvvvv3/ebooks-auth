package models

import (
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

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(EbookMsg))
}
