
package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"zcxt-go/models"
)

type VerifyController struct {
	beego.Controller
}

type bindVerifyUpdate struct {
	models.BaseType
  	models.BaseId
	models.BaseRemark
	models.BaseRealScore
}
func (this *VerifyController) Update() {
	id := models.GetCookie(this.Ctx.GetCookie("token"),"id")
	if id == "" {
		this.Data["json"] = models.Status(135)
		this.ServeJSON()
		return
	}
	var req bindVerifyUpdate
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Data["json"] = models.Status(132)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.VerifyUpdate(req.Type,req.Id,id,req.Remark,req.RealScore)
	this.ServeJSON()
	return
}

func (this *VerifyController) ReadWait() {
	id := models.GetCookie(this.Ctx.GetCookie("token"),"id")
	if id == "" {
		this.Data["json"] = models.Status(135)
		this.ServeJSON()
		return
	}

	var req models.BaseName
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Data["json"] = models.Status(132)
		this.ServeJSON()
		return
	}

	this.Data["json"] = models.VerifyReadWait(req.Name)
	this.ServeJSON()
	return
}

func (this *VerifyController) Read() {
	id := models.GetCookie(this.Ctx.GetCookie("token"),"id")
	if id == "" {
		this.Data["json"] = models.Status(135)
		this.ServeJSON()
		return
	}

	var req models.BaseState
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Data["json"] = models.Status(132)
		this.ServeJSON()
		return
	}

	this.Data["json"] = models.VerifyRead(req.State,id)
	this.ServeJSON()
	return
}

func (this *VerifyController) Create() {
	id := models.GetCookie(this.Ctx.GetCookie("token"),"id")
	if id == "" {
		this.Data["json"] = models.Status(135)
		this.ServeJSON()
		return
	}
	reqCurrentTreeId := this.GetString("initTreeId")
	reqId := this.GetString("id")
	reqType := this.GetString("type")
	reqApplyScore,_ := this.GetInt("applyScore")
	f, h, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = models.Status(134)
		this.ServeJSON()
		return
	}
	uploadRes := models.UploadImgToMinio(h.Filename,f)
	if uploadRes == "" {
		this.Data["json"] = models.Status(133)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.VerifyCreate(uploadRes,reqCurrentTreeId,id,reqId,reqType,reqApplyScore)
	this.ServeJSON()
	return
}
