package controllers

import "github.com/astaxie/beego"

type UserManagementController struct {
	beego.Controller
}

func (c *UserManagementController) Get() {
	c.TplName = "user-mgt.html"
}
