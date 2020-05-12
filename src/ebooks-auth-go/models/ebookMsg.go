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

// 添加书籍
func AddEbookMsg(ebookMsg *EbookMsg) int {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("xixi", ebookMsg)
	_, err := o.Insert(ebookMsg)

	if err != nil {
		fmt.Println("失败！！！", err)
	}

	return ebookMsg.Id

}

// 删除书籍
func DelEbookMsg(id string) string {
	v, errv := strconv.Atoi(id)
	u := EbookMsg{Id: v}
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

// 更新书籍
func UpdateEbookMsg(ebookMsg *EbookMsg) string {
	//1.有ORM对象
	o := orm.NewOrm()
	//2.查询要更新的结构体对象
	list := EbookMsg{}
	//3.查询需要更新的数据
	list.Id = ebookMsg.Id
	err := o.Read(&list)
	//4.给数据重新赋值
	if err == nil {
		list.Isbn = ebookMsg.Isbn
		list.Name = ebookMsg.Name
		list.Price = ebookMsg.Price
		list.Publish_date = ebookMsg.Publish_date
		list.Publisher = ebookMsg.Publisher
		list.Writer = ebookMsg.Writer
		list.Page = ebookMsg.Page
		list.Buy_times = ebookMsg.Buy_times
		// list.Grade = ebookList.Grade
		list.Category = ebookMsg.Category
		fmt.Println(list)
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
	orm.RegisterModel(new(EbookMsg))
}
