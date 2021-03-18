
package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"zcxt-go/models"
)

type UserController struct {
	beego.Controller
}

type bindUserLogin struct {
	models.BaseId
  models.BasePassword
  models.BaseState
}
func (this *UserController) Login() {
	var req bindUserLogin
	if err := this.ParseForm(&req); err != nil {
	 	this.Data["json"] = models.Status(132)
		this.ServeJSON()
	 	return
	}
	res := models.UserLogin(req.Id,req.Password,req.State)
	if res["status"] == 200 {
		this.Ctx.SetCookie("token", models.SetCookie(map[string]string{"id": res["data"].(map[string]string)["id"]}), 604800)
	}
	this.Data["json"] = res
	this.ServeJSON()
	return
}

type bindUserUpdate struct {
	models.BaseName
  models.BaseInitTreeName
}
func (this *UserController) Update() {
	id := models.GetCookie(this.Ctx.GetCookie("token"),"id")
	if id == "" {
		this.Data["json"] = models.Status(135)
		this.ServeJSON()
		return
	}
	var req bindUserUpdate
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Data["json"] = models.Status(132)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.UserUpdate(req.Name,req.InitTreeName,id)
	this.ServeJSON()
	return
}

func (this *UserController) ReadItem() {
	this.Data["json"] = models.UserReadItem()
	this.ServeJSON()
	return
}

func (this *UserController) ReadOrganize() {
	this.Data["json"] = models.UserReadOrganize()
	this.ServeJSON()
	return
}
