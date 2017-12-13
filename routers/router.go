package routers

import (
	"github.com/astaxie/beego"
	"myBlog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/blog/detail", &controllers.BlogDetailController{})
	beego.Router("/blog/edit", &controllers.BlogEditInit{})
}
