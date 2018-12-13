package controllers

import (
	"account"
	"fmt"
	"github.com/astaxie/beego"
)

type UserJudgeController struct {
	beego.Controller
}

func (c *UserJudgeController) Get() {
	res := account.IsGlobalAdmin(UserName)
	if res {
		fmt.Fprint(c.Ctx.ResponseWriter, "true")
	}else {
		fmt.Fprint(c.Ctx.ResponseWriter, "false")
	}
}
