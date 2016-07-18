package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type CommentController struct {
	beego.Controller
}

//just delete comment
func (this *CommentController) Get() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	id := this.Input().Get("id")
	tid := this.Input().Get("tid")
	err := models.DelComment(id, tid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/"+tid+"#comment", 302)
	return

}

func (this *CommentController) Post() {

	name := this.Input().Get("name")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")

	err := models.AddComment(tid, name, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/"+tid+"#comment", 302)
}
