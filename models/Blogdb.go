package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
)

func (blog *Blog) FetchById() {
	o := orm.NewOrm()
	o.Read(blog)
}

func (blog *Blog) Add() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(blog)
	return id, err
}

func (blog *Blog) Update() (int64, error) {
	o := orm.NewOrm()
	return o.Update(blog)
}

/*返回列表*/
func (blog *Blog) List() ([]Blog, error) {
	o := orm.NewOrm()
	blogs := make([]Blog, 0)
	//var blogs []Blog
	//blogs := new([]Blog)

	_, err := o.QueryTable(new(Blog)).OrderBy("-Id").All(&blogs)
	//fmt.Println(num)
	return blogs, err

}
