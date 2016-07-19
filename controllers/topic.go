package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"os"
	"path"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	id := this.Ctx.Input.Param(":id")
	if len(id) != 0 {

		this.Data["IsTopic"] = true
		this.TplName = "topic_view.tpl"
		this.Data["IsLogin"] = checkAccount(this.Ctx)

		topic, err := models.ViewTopic(id)
		if err != nil {
			beego.Error(err)
		}

		comment, err := models.GetAllComment(id)
		if err != nil {
			beego.Error(err)
		}

		this.Data["Topic"] = topic
		this.Data["Comment"] = comment
		return
	}
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
	category, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["Category"] = category
}

func (this *TopicController) Post() {

	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	var err error
	title := this.Input().Get("title")
	category := this.Input().Get("category")
	content := this.Input().Get("content")
	id := this.Input().Get("id")

	var attachment string
	_, fh, _ := this.GetFile("attachment")

	if fh != nil {
		attachment = fh.Filename
		beego.Info(attachment)
	}

	if len(id) == 0 {
		tid, err := models.AddTopic(title, category, content, attachment)
		if err != nil {
			beego.Error(err)
		}

		os.Mkdir(path.Join("attachment", string(tid)), os.ModePerm)
		err = this.SaveToFile("attachment", path.Join("attachment", string(tid), attachment))
		if err != nil {
			beego.Error(err)
		}
	} else {
		oldAttach := this.Input().Get("oldAttach")
		if oldAttach != attachment {
			os.Remove(path.Join("attachment", id, oldAttach))
		}
		if fh == nil {
			err = models.ModifyTopic(id, title, category, content, oldAttach)
			if err != nil {
				beego.Error(err)
			}
		} else {
			os.Mkdir(path.Join("attachment", id), os.ModePerm)
			err = this.SaveToFile("attachment", path.Join("attachment", id, attachment))
			if err != nil {
				beego.Error(err)
			}
			err = models.ModifyTopic(id, title, category, content, attachment)
			if err != nil {
				beego.Error(err)
			}
		}
	}

	this.Redirect("/topic/"+id, 302)
}
func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	id := this.Ctx.Input.Param("0")
	category := this.Ctx.Input.Param("1")

	err := models.DelTopic(id, category)
	if err != nil {
		beego.Error(err)
	}
	os.RemoveAll(path.Join("attachment", id))
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

	category, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topic"] = topic
	this.Data["Id"] = id
	this.Data["Category"] = category
}
