package controllers

import (
	"account"
	"github.com/astaxie/beego"
)

var WrongUserPass = false
var LogIn = false
var UserName string
var Password string

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
	UserName = ""
	Password = ""

	UserName = c.Input().Get("username")
	Password = c.Input().Get("password")

	result := account.AccountValidate(UserName, Password)

	if result {
		LogIn = true
		c.Redirect("/", 301)
	} else {
		WrongUserPass = true
		c.Redirect("/login", 301)
	}
}


