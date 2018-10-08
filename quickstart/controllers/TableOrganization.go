package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
)

type OrgTableController struct {
	beego.Controller
}

func (c *OrgTableController) Get() {
	org := restapi.GetOrganizations("10")
	b, err := json.Marshal(org)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}

