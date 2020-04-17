package main

import (
	_ "ebooks-auth-go/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 连接数据库
	orm.RegisterDataBase("default", "mysql", "root:wei934934@tcp(cdb-cw3mglxe.cd.tencentcdb.com:10127)/ebooks_vv?charset=utf8")
	// 对数据库中user_msg表进行查询并显示
	// fmt.Println("models is inited!")

	// o := orm.NewOrm()

	// var maps []orm.Params
	// num, _ := o.Raw("SELECT * FROM user_info").Values(&maps)

	// for _, term := range maps {
	// 	fmt.Println(term["id"], ":", term["name"], num)
	// }
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
