package routers

import (
	"NewsRoom/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wellcome/:id([0-9]+)", &controllers.MainController{}, "get:Wellcome")
}
