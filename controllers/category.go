package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	category := this.Ctx.Input.Param(":category")
	if len(category) != 0 {

		var err error
		this.Data["Topic"], err = models.GetCateTopics(category)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Category"] = category

		this.Data["Categories"], err = models.GetAllCategories()

		if err != nil {
			beego.Error(err)
		}

		this.Data["IsLogin"] = checkAccount(this.Ctx)
		this.Data["IsCategory"] = true
		this.TplName = "category_view.tpl"
		return
	}
	op := this.Input().Get("op")

	switch op {
	case "add":
		if !checkAccount(this.Ctx) {
			this.Redirect("/login", 302)
			return
		}
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	case "del":
		if !checkAccount(this.Ctx) {
			this.Redirect("/login", 302)
			return
		}
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	}
	this.Data["IsCategory"] = true
	this.TplName = "category.tpl"

	var err error
	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}

	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
