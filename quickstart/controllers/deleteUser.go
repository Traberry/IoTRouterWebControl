package controllers

import (
	"account"
	"github.com/astaxie/beego"
	"log"
	"strings"
)

type UserDeleteController struct {
	beego.Controller
}

func (c *UserDeleteController) Delete() {
	//userID := c.Input().Get("id")
	//fmt.Println("DELETE USER BASED ON ID:", userID)
	//id := c.Ctx.Input.Data()["id"]
	//fmt.Println("ID: ", id)

	body := string(c.Ctx.Input.RequestBody)
	strs := strings.Split(body, "=")
	userID := strs[1]

	if account.DeleteUser(userID) {
		log.Println("Delete user success for user ID:", userID)
		c.Ctx.ResponseWriter.Write([]byte{'s', 'u', 'c', 'c', 'e', 's', 's'})
	}else {
		c.ServeJSON()
	}
}
