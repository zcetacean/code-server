package main

import (
	"zcxt-go/models"
	_ "zcxt-go/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "prod" {
		beego.BeeLogger.DelLogger("console")
		beego.SetLevel(beego.LevelWarning)
	} else if beego.BConfig.RunMode == "dev" {
		beego.SetLevel(beego.LevelDebug)
	}

	orm.RegisterDataBase("default", "mysql", "root:Zpj075485329759@tcp(8.129.44.170:3306)?charset=utf8")
	orm.RegisterModel(new(models.Organize), new(models.Item), new(models.Verify), new(models.Admin), new(models.Record), new(models.Student))
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
