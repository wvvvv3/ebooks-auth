package main

import (
	_ "house-auth/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 连接数据库
	orm.RegisterDataBase("default", "mysql", "root:wei934934@tcp(cdb-cw3mglxe.cd.tencentcdb.com:10127)/ebooks_vv?charset=utf8")
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
