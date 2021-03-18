
package routers

import (
  "zcxt-go/controllers"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/plugins/cors"
)

func init() {
  beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
    AllowOrigins:     []string{},
    AllowMethods:     []string{"GET", "POST", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
    AllowCredentials: true,
  }))

  ns := beego.NewNamespace("/api",
    beego.NSAutoRouter(&controllers.VerifyController{}),
    beego.NSAutoRouter(&controllers.RecordController{}),
    beego.NSAutoRouter(&controllers.UserController{}),
  )

  beego.AddNamespace(ns)
}
