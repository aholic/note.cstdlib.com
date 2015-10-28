package controllers

import (
	"html"

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

func (c *NoteController) Show() {
	url := c.Ctx.Input.Param("0")
	if noteEntry, err := models.GetNoteEntryByUrl(url); err != nil {
		c.Abort("404")
	} else {
		c.Data["noteContent"] = html.UnescapeString(noteEntry.GetContent())
		c.Data["noteDate"] = noteEntry.GetDate()
	}
	c.TplNames = "note/show.tpl"
}

func makeAjaxResponse(flag bool, data string, msg string) map[string]interface{} {
	return map[string]interface{}{"flag": flag, "data": data, "msg": msg}
}

func (c *NoteController) Submit() {
	uname := c.Ctx.GetCookie("uname")
	if uname == "" {
		uname = models.GenerateUUID()
		c.Ctx.SetCookie("uname", uname)
	}

	noteContent := html.EscapeString(c.Input().Get("noteContent"))
	if noteEntry, err := models.NewNoteEntry(uname, noteContent); err != nil {
		c.Data["json"] = makeAjaxResponse(false, "", "something wrong with server")
	} else {
		beego.Debug(noteEntry.GetUrl())
		c.Data["json"] = makeAjaxResponse(true, noteEntry.GetUrl(), "succ")
	}

	c.ServeJson()
}
