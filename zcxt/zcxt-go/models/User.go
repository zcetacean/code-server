
package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func UserLogin(Id string,Password string,State string) map[string]interface{} {
  o := orm.NewOrm()
	result := make(map[string]string)
	result["id"] = Id
	if State == "student" {

		var student Student
		o.QueryTable("student").Filter("id", Id).One(&student)
		if student.Password != Password {
			res := Status(131)
			return res
		}
		result["name"] = student.Name
		result["initTreeName"] = student.InitTreeName
		return Status(200,result)
	} else {

		var admin Admin
		o.QueryTable("admin").Filter("id", Id).One(&admin)
		if admin.Password != Password {
			res := Status(131)
			return res
		}
		result["name"] = admin.Name
		result["permissionTreeName"] = admin.PermissionTreeName
		return Status(200,result)
	}
}

func UserUpdate(Name,InitTreeName,Id string) map[string]interface{} {
	o := orm.NewOrm()
	num, err := o.QueryTable("student").Filter("id", Id).Update(orm.Params{
		"init_tree_name": InitTreeName,
		"name":Name,
	})
	if err != nil {
		res := Status(141)
		return res
	}
	if num != 1 {
		res := Status(142)
		return res
	}
	res := Status(200)
	return res
}

func UserReadItem() map[string]interface{} {
  o := orm.NewOrm()
	var ormRes []orm.Params
	_, err := o.Raw(`SELECT * FROM item`).Values(&ormRes)

	if err != nil {
		res := Status(141)
		return res
	}
	res := Status(200,ormRes)
	return res
}

func UserReadOrganize() map[string]interface{} {
  o := orm.NewOrm()
	var ormRes []orm.Params
	_, err := o.Raw(`SELECT name FROM organize WHERE state = 0`).Values(&ormRes)
	if err != nil {
		res := Status(141)
		return res
	}
	res := Status(200,ormRes)
	return res
}
