package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/category", &controllers.CategoryController{})

	beego.Router("/category/?:category", &controllers.CategoryController{})

	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/topic", &controllers.TopicController{})

	beego.Router("/topic/?:id", &controllers.TopicController{})

	beego.Router("/comment", &controllers.CommentController{})
}
