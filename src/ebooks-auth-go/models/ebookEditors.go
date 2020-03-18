package models

import (
	"github.com/astaxie/beego/orm"
)

type EbookEditors struct {
	Id       int
	Ebook_id int
	Name     string
	Pic      string
}

func GetAllEbookEditors() []*EbookEditors {
	// fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var ebookEditors []*EbookEditors
	q := o.QueryTable("ebook_editors")
	q.All(&ebookEditors)
	return ebookEditors

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(EbookEditors))
}
