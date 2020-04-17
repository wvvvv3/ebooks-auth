package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type EbookComment struct {
	Id              int    `json:"id"`
	Ebook_id        int    `json:"ebook_id"`
	User_id         int    `json:"user_id"`
	Comment_content string `json:"comment_content"`
	Comment_time    string `json:"comment_time"`
	Comment_grade   int    `json:"comment_grade"`
	User_name       string `json:"user_name"`
	User_img        string `json:"user_img"`
}

func GetEbookCommentById(id string) []*EbookComment {
	// string转换成int
	v, errv := strconv.Atoi(id)

	o := orm.NewOrm()
	o.Using("default")
	var ebookComment []*EbookComment
	q := o.QueryTable("ebook_comment").Filter("Ebook_id", v).RelatedSel()
	q.All(&ebookComment)
	if errv != nil {
		fmt.Println("id获取错误")
	}
	// q.All(&ebookList)
	return ebookComment

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(EbookComment))
}
