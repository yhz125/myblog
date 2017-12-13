package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*博客分类表*/
type Category struct {
	Id        int
	CateName  string
	LastModif int64
}

/*博客*/
type Blog struct {
	Id        int
	Title     string `orm:size(100)`
	Content   string //博客内容
	Created   int64  //创建时间
	LastModif int64  //最后修改时间
}

type BlogComment struct {
	Id       int
	BlogId   int
	Comments string `orm:size(500)` //评论内容
	Created  int64  //评论时间
}

func init() {
	//注册数据库并创建db连接
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/goDb?charset=utf8", 30)

	//关联表对象
	orm.RegisterModel(new(Category), new(Blog), new(BlogComment))

	//创建表
	orm.RunSyncdb("default", false, true)
}
