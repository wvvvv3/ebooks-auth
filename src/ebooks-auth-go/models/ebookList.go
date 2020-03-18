package models

import (
	"github.com/astaxie/beego/orm"
)

type EbookList struct {
	Id        int
	List_name string
	List_pic  string
	Ebook_id  string
	// `orm:"column(uid);pk"`
}

func GetAllEbookList() []*EbookList {
	// fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var ebookList []*EbookList
	q := o.QueryTable("ebook_list")
	q.All(&ebookList)
	return ebookList

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(EbookList))
}
