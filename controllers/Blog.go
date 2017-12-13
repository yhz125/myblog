package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myBlog/models"
	"strconv"
	//"strings"
)

/*列表页*/
type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {

	//blog := &models.Blog{}
	blog := new(models.Blog)
	blogs, err := blog.List()
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	for _, v := range blogs {
		this.Ctx.WriteString(v.Title + "\t" + v.Content + "\n")
	}
}

/*添加/修改*/
func (this *BlogController) Post() {
	blog := new(models.Blog)
	blog.Title = this.Input().Get("Title")
	blog.Id, _ = strconv.Atoi(this.Input().Get("Id"))
	blog.Content = this.Input().Get("Content")
	//num, err := int64(0), inteface{}

	if blog.Id <= 0 {
		fmt.Println("添加博客...")
		_, err := blog.Add()
		if err != nil {
			fmt.Println("操作异常:" + err.Error())
		}
	} else {
		fmt.Println("编辑博客...")
		_, err := blog.Update()
		if err != nil {
			fmt.Println("操作异常:" + err.Error())
		}
	}

	this.Ctx.Redirect(302, "/home")
	return
}

/*添加/修改初始页*/
type BlogEditInit struct {
	beego.Controller
}

func (this *BlogEditInit) Get() {
	Id, err := strconv.Atoi(this.Input().Get("id"))
	fmt.Println(Id, err)
	if err == nil && Id > 0 { //编辑
		blog := &models.Blog{Id: Id}
		blog.FetchById()
		this.Data["blog"] = blog
		fmt.Println(blog)
	}
	this.Layout = "Shard/Layout.tpl"
	this.TplName = "detailInit.tpl"
	return

}

/*详情页*/
type BlogDetailController struct {
	beego.Controller
}

func (this *BlogDetailController) Get() {
	id, err := strconv.Atoi(this.Input().Get("id"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	blog := &models.Blog{Id: id}
	blog.FetchById()
	this.Data["detail"] = blog
	this.Layout = "Shard/Layout.tpl"
	this.TplName = "detail.tpl"
	return
}
