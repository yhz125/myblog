package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myBlog/models"
	//"strconv"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	var blog = &models.Blog{}
	blogs, err := blog.List()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	this.Data["isHome"] = true
	this.Data["blogs"] = blogs
	this.Layout = "Shard/Layout.tpl"
	this.TplName = "home.tpl"
	return
	//this.Ctx.WriteString("ttttt")

}
