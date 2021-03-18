package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func VerifyReadWait(Name string) map[string]interface{} {
	o := orm.NewOrm()

	var allRes []Organize
	o.QueryTable("organize").Filter("name", Name).All(&allRes)

	var outRes []Verify
	var inRes []Verify

	for _, item := range allRes {
		o.QueryTable("verify").Filter("state", "审核中").Filter("current_tree_id", item.Id).RelatedSel().All(&inRes)
		outRes = append(outRes, inRes...)
	}
	for index, item := range outRes {
		outRes[index].Img = GetImgUrlInMinio(item.Img)
	}
	return Status(200, outRes)
}

func VerifyUpdate(Type, Id, AdminId, Remark string, Score int) map[string]interface{} {
	o := orm.NewOrm()
	var remarkStr string
	var verify Verify
	o.QueryTable("verify").Filter("id", Id).RelatedSel().One(&verify)
	if Type == "yes" {
		var organize Organize
		o.QueryTable("organize").Filter("id", verify.CurrentTreeId).One(&organize)
		state := "审核中"
		treeId := organize.ParentTreeId
		if organize.ParentTreeId == "" {
			treeId = verify.CurrentTreeId
			state = "已通过"
		}
		_, err := o.QueryTable("verify").Filter("id", Id).Update(orm.Params{
			"state":           state,
			"remark":          Remark,
			"current_tree_id": treeId,
			"real_score":      Score,
		})
		if err != nil {
			res := Status(141)
			return res
		}
		remarkStr = "通过" + verify.Student.Id + "的综测申请/备注：" + verify.Remark
	} else {
		_, err := o.QueryTable("verify").Filter("id", Id).Update(orm.Params{
			"state":  "未通过",
			"remark": Remark,
		})
		if err != nil {
			res := Status(141)
			return res
		}
		remarkStr = "拒绝" + verify.Student.Id + "的综测申请/备注：" + verify.Remark
	}
	o.Raw("INSERT INTO record VALUES (?, ?,?,?)", RandomCreater(), CurTime(), remarkStr, AdminId).Exec()
	res := Status(200)
	return res
}

func VerifyRead(State, Id string) map[string]interface{} {
	o := orm.NewOrm()
	var res []Verify
	if State == "finish" {
		o.QueryTable("verify").Exclude("state", "审核中").Filter("Student", Id).RelatedSel().All(&res)
	} else {
		o.QueryTable("verify").Filter("state", "审核中").Filter("Student", Id).RelatedSel().All(&res)
	}
	for index, item := range res {
		res[index].Img = GetImgUrlInMinio(item.Img)
	}
	return Status(200, res)
}

func VerifyCreate(img, currentTreeId, studentId, itemId, types string, applyScore int) map[string]interface{} {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO verify VALUES (?, ?,?,?,?,?,?,?,?,?,?)", RandomCreater(), "审核中", img, 0, applyScore, "", types, CurTime(), currentTreeId, studentId, itemId).Exec()
	if err != nil {
		//DelImgInMinio(img)
		res := Status(141)
		return res
	}
	return Status(200)
}
