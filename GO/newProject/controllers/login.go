package controllers

import (
	"github.com/astaxie/beego"
)

//LoginController
type LoginController struct {
	beego.Controller
}

//Get
func (c *LoginController) ShowLogin() {
	// c.Ctx.WriteString("cetacean.top")
	c.TplName = "login.html"
}

//Post
func (c *LoginController) SendInfo() {
	uuid := c.GetString("uuid")
	pwd := c.GetString("password")
	// beego.Info("账号:", uuid, "密码:", pwd)
	// c.Ctx.WriteString("hello")
	if uuid != "demoUuid" {
		c.TplName = "login.html"
		c.Data["wrong"] = "账号错误"
		return
	}
	if pwd != "demoPassword" {
		c.TplName = "login.html"
		c.Data["wrong"] = "密码错误"
		return
	}
	c.TplName = "login.html"
	c.Data["wrong"] = "登录成功"
	return

}
