package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id   int
	Name string
}

func GetAllStudents() []*Student {
	fmt.Println("models is inited!mo")
	o := orm.NewOrm()
	o.Using("default")
	var students []*Student
	q := o.QueryTable("student")
	q.All(&students)
	return students

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Student))
}
