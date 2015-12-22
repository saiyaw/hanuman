package routers

import (
	"github.com/saiyawang/candidate/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
