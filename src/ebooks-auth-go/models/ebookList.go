package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type EbookList struct {
	Id        int
	List_name string
	List_pic  string
	Ebook_id  string
	// `orm:"column(uid);pk"`
}

// 获取所有书单信息
func GetAllEbookList() []*EbookList {
	o := orm.NewOrm()
	o.Using("default")
	var ebookList []*EbookList
	q := o.QueryTable("ebook_list")
	q.All(&ebookList)
	return ebookList

}

func GetEbookListById(id string) EbookList {
	// string转换成int
	v, errv := strconv.Atoi(id)
	u := EbookList{Id: v}
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
	orm.RegisterModel(new(EbookList))
}
