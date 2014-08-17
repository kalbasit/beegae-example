package controllers

import (
	"github.com/astaxie/beegae"
)

type MainController struct {
	beegae.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beegae.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
