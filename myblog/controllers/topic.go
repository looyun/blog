package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplName = "topic.tpl"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}

func (this *TopicController) Add() {
	this.Data["IsTopic"] = true
	this.TplName = "topic_add.tpl"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	id := this.Input().Get(("id"))
	var err error
	if len(id) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(id, title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}
func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	id := this.Ctx.Input.Param("0")

	err := models.DelTopic(id)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic_modify.tpl"

	id := this.Ctx.Input.Param("0")

	topic, err := models.ViewTopic(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topic"] = topic
	this.Data["Id"] = id
}

func (this *TopicController) View() {
	this.Data["IsTopic"] = true
	this.TplName = "topic_view.tpl"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	id := this.Ctx.Input.Param("0")

	topic, err := models.ViewTopic(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topic"] = topic
}
