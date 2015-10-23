package routers

import (
	"github.com/astaxie/beego"
	"note.cstdlib.com/controllers"
)

func init() {
	beego.Router("/*", &controllers.NoteController{}, "*:Show")
	beego.Router("/note/submit", &controllers.NoteController{}, "*:Submit")
	beego.Router("/", &controllers.NoteController{}, "*:New")
}
