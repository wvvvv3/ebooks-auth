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

// 通过书单id获取书单信息
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

// 添加书单
func AddEbookList(ebookList *EbookList) int {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("xixi", ebookList)
	_, err := o.Insert(ebookList)

	if err != nil {
		fmt.Println("失败！！！")
	}

	return ebookList.Id
}

// 删除书单
func DelEbookList(id string) string {
	v, errv := strconv.Atoi(id)
	u := EbookList{Id: v}
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("daada", u)
	_, err := o.Delete(&u)
	if err != nil {
		fmt.Println("删除失败", err)
		return "error"
	}
	if errv != nil {
		fmt.Println("id获取错误")
	}
	return "seccess"
}

// 更新书单
func UpdateEbookList(ebookList *EbookList) string {
	//1.有ORM对象
	o := orm.NewOrm()
	//2.查询要更新的结构体对象
	list := EbookList{}
	//3.查询需要更新的数据
	list.Id = ebookList.Id
	err := o.Read(&list)
	//4.给数据重新赋值
	if err == nil {
		list.Ebook_id = ebookList.Ebook_id
		_, err := o.Update(&list)
		if err != nil {
			return "error"
		}
		return "seccess"
	}
	return "error"
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(EbookList))
}
