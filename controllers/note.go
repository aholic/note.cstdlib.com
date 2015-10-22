package controllers

import (
	"github.com/astaxie/beego"
	"note.cstdlib.com/models"
)

type NoteController struct {
	beego.Controller
}

func (c *NoteController) New() {
	uname := c.Ctx.GetCookie("uname")
	if uname == "" {
		c.Ctx.SetCookie("uname", models.GenerateUUID())
	}
	c.TplNames = "note/new.tpl"
}

func makeAjaxResponse(flag bool, data string, msg string) map[string]interface{} {
	return map[string]interface{}{"falg": flag, "data": data, "msg": msg}
}

func (c *NoteController) Submit() {
	uname := c.Ctx.GetCookie("uname")
	if uname == "" {
		uname = models.GenerateUUID()
		c.Ctx.SetCookie("uname", uname)
	}

	noteContent := c.Input().Get("noteContent")
	noteEntry := models.NewNoteEntry(uname, noteContent)
	if noteEntry.Save() {
		c.Data["json"] = makeAjaxResponse(true, noteEntry.GetUrl(), "succ")
	} else {
		c.Data["json"] = makeAjaxResponse(false, "", "something wrong with server")
	}

	c.ServeJson()
}
