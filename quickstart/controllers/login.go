package controllers

import (
	"account"
	"github.com/astaxie/beego"
)

var WrongUserPass = false

type LogController struct {
	beego.Controller
}

func (c *LogController) Get() {
	if WrongUserPass {
		c.Data["failtip"] = "wrong username or password"
		c.Data["p"] = "animated shake"
		WrongUserPass = false
	}
	c.TplName = "login.html"
}

func (c *LogController) Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")

	result := account.AccountValidate(username, password)

	if result {
		c.Redirect("/", 301)
	} else {
		WrongUserPass = true
		c.Redirect("/login", 301)
	}
}


