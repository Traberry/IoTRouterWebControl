package controllers

import "github.com/astaxie/beego"

type UserManagementController struct {
	beego.Controller
}

func (c *UserManagementController) Get() {
	//judge user login first
	if !LogIn {
		c.Redirect("/login", 301)
	}

	c.TplName = "user-mgt.html"
}
