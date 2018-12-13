package controllers

import (
	"account"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

type UserUpdateController struct {
	beego.Controller
}

func (c *UserUpdateController) Put() {
	var userUpdate account.UserForUpdate
	var userPasswd account.UserPasswordChange

	//bind data to userPasswd
	err := c.Ctx.Input.Bind(&userPasswd.ID, "userID")
	if err != nil {
		log.Println("change user password error:", err)
	}
	err = c.Ctx.Input.Bind(&userPasswd.Password, "userPassword")
	if err != nil {
		log.Println("change user password error:", err)
	}
	//fmt.Println("userPasswd: ", userPasswd)

	//bind data to userUpdate
	err = c.Ctx.Input.Bind(&userUpdate.ID, "userID")
	if err != nil {
		log.Println("update user error:", err)
	}
	err = c.Ctx.Input.Bind(&userUpdate.UserName, "userName")
	if err != nil {
		log.Println("update user error:", err)
	}
	err = c.Ctx.Input.Bind(&userUpdate.Email, "userEmail")
	if err != nil {
		log.Println("update user error:", err)
	}

	var isActive, isAdmin string
	err = c.Ctx.Input.Bind(&isActive, "isActive")
	if err != nil {
		log.Println(err)
	}
	err = c.Ctx.Input.Bind(&isAdmin, "isGlobalAdmin")
	if err != nil {
		log.Println(err)
	}
	if isActive == "是" {
		userUpdate.IsActive = true
	}
	if isAdmin == "是" {
		userUpdate.IsAdmin = true
	}
	//fmt.Println("userUpdate: ", userUpdate)

	//call rest api to make update work
	updateResult := account.UpdateUser(userUpdate, userUpdate.ID)
	passwdResult := account.ChangeUserPassword(userPasswd.ID, userPasswd)

	if updateResult && passwdResult {
		fmt.Fprint(c.Ctx.ResponseWriter, "modify user success")
		c.Redirect("/user", 301)
	}else {
		fmt.Fprint(c.Ctx.ResponseWriter, "modify user failed")
	}
}