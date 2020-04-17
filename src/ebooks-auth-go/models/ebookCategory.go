package models

import (
	"fmt"
	"strconv"

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
func GetEbookCategoryById(id string) EbookCategory {
	// string转换成int
	v, errv := strconv.Atoi(id)
	u := EbookCategory{Id: v}
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
	orm.RegisterModel(new(EbookCategory))
}
