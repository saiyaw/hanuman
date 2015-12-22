package routers

import (
	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.ListController{})

	beego.Router("/new", &controllers.NewController{})

	beego.Router("/view", &controllers.ViewController{})
}
