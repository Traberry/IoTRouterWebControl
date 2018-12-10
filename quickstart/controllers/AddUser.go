package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserAddController struct {
	beego.Controller
}

func (c *UserAddController) Post() {
	//var userToCreat account.UserToBeCreated

	fmt.Println(c.Ctx.Request.Form["userName"])
	//fmt.Println(c.Ctx.Input.GetData("userName"))

	var u, e string
	err := c.Ctx.Input.Bind(&u, "userName")
	if err != nil {
		fmt.Println(err)
	}
	err = c.Ctx.Input.Bind(&e, "userEmail")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("userName: ", u)
	fmt.Println("userEmail: ", e)

	fmt.Fprintf(c.Ctx.ResponseWriter, "success")
}