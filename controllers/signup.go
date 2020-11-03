package controllers

import (
	"github.com/astaxie/beego"
)

type EntryController struct {
	beego.Controller
}

func (this *EntryController) EntryPage() {
	this.TplName = "entryPage.html"
}

