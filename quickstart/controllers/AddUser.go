package controllers

import (
	"account"
	"fmt"
	"github.com/astaxie/beego"
)

type UserAddController struct {
	beego.Controller
}

func (c *UserAddController) Post() {
	//fmt.Println(c.Ctx.Request.Form["userName"])
	//fmt.Println(c.Ctx.Input.GetData("userName"))

	var userToCreat account.UserToBeCreated


	err := c.Ctx.Input.Bind(&userToCreat.UserName, "userName")
	if err != nil {
		fmt.Println(err)
	}
	err = c.Ctx.Input.Bind(&userToCreat.Email, "userEmail")
	if err != nil {
		fmt.Println(err)
	}
	err = c.Ctx.Input.Bind(&userToCreat.Password, "userPassword")
	if err != nil {
		fmt.Println(err)
	}


	var orgID, isActive, isAdmin string
	err = c.Ctx.Input.Bind(&orgID, "isOrgAdmin")
	if err != nil {
		fmt.Println(err)
	}
	err = c.Ctx.Input.Bind(&isActive, "isActive")
	if err != nil {
		fmt.Println(err)
	}
	err = c.Ctx.Input.Bind(&isAdmin, "isGlobalAdmin")
	if err != nil {
		fmt.Println(err)
	}

	if isActive == "是" {
		userToCreat.IsActive = true
	}
	if isAdmin == "是" {
		userToCreat.IsAdmin = true
	}

	userToCreat.Organizations = append(userToCreat.Organizations, account.BelongOrg{OrganizationID: orgID, IsAdmin: userToCreat.IsAdmin})


	fmt.Println(userToCreat)

	userID, err := account.CreateUser(userToCreat)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(c.Ctx.ResponseWriter, "error while creating user")
	}else {
		fmt.Fprintf(c.Ctx.ResponseWriter, "success:" + " " + userID)
	}
}