package routers

import (
	"newProject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:SendInfo")
}
