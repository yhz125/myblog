package models

import (
	"github.com/astaxie/beego/orm"
)

func Insert(category *Category) {
	o := orm.NewOrm()

	o.Insert(category)
}

func Update(category *Category) {
	o := orm.NewOrm()

	o.Update(category)
}

func FetchById(id int) {
	o := orm.NewOrm()
	category := &Category{Id: id}
	o.Read(category)
}

func init() {

}
