package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/category", &controllers.CategoryController{})

	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/topic", &controllers.TopicController{})
}
