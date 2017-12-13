package main

import (
	"github.com/astaxie/beego"
	_ "myBlog/models"
	_ "myBlog/routers"
)

func main() {
	beego.Run()
}
