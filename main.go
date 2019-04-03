package main

import (
	"test/db"
	_ "test/routers"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db.InitDatabase()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
