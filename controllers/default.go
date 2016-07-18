package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "index.tpl"

	c.Data["IsLogin"] = checkAccount(c.Ctx)

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	category, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Category"] = category
	c.Data["Topics"] = topics
}
