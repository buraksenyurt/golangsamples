package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Wellcome() {

	c.Data["Id"] = c.Ctx.Input.Param(":id")
	c.Data["Website"] = "News Room"
	c.Data["Email"] = "selim@buraksenyurt.com"
	c.Data["EmailName"] = "Burak Selim Senyurt"
	c.TplName = "default/hello-newsroom.tpl"
}
