package routers

import (
	"github.com/astaxie/beego"
	"note.cstdlib.com/controllers"
)

func init() {
	beego.AutoRouter(&controllers.NoteController{})
	beego.Router("/", &controllers.MainController{})
}
