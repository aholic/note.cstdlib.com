package main

import (
	"github.com/astaxie/beego"
	_ "note.cstdlib.com/routers"
)

func main() {
	beego.SessionOn = true
	beego.SessionCookieLifeTime = 3600 * 24 * 30 // a month
	beego.Run()
}
