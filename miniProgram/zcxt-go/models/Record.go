
package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//o.Raw("INSERT INTO record VALUES (?, ?,?,?)", RandomCreater(), CurTime(),Remark,adminId).Exec()

func RecordRead(Id string) map[string]interface{} {
  o := orm.NewOrm()
	var ormRes []orm.Params
	_, err := o.Raw(`SELECT * FROM record WHERE admin_id = ?`,Id).Values(&ormRes)
	if err != nil {
		res := Status(141)
		return res
	}
	res := Status(200,ormRes)
	return res
}
