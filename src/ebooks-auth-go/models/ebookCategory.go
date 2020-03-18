package models

import (
	"github.com/astaxie/beego/orm"
)

type EbookCategory struct {
	Id            int
	Category_name string
	Ebook_id      string
}

func GetAllEbookCategory() []*EbookCategory {
	o := orm.NewOrm()
	o.Using("default")
	var ebookCategory []*EbookCategory
	q := o.QueryTable("ebook_category")
	q.All(&ebookCategory)
	return ebookCategory

}

func init() {
	orm.RegisterModel(new(EbookCategory))
}
