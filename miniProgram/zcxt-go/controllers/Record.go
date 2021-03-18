
package controllers

import (
	"github.com/astaxie/beego"
	"zcxt-go/models"
)

type RecordController struct {
	beego.Controller
}

func (this *RecordController) Read() {
	id := models.GetCookie(this.Ctx.GetCookie("token"),"id")
	if id == "" {
		this.Data["json"] = models.Status(135)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.RecordRead(id)
	this.ServeJSON()
	return
}
