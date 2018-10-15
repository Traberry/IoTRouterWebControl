package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gpmgo/gopm/modules/log"
	"restapi"
)

type OrgTableController struct {
	beego.Controller
}

func (c *OrgTableController) Get() {
	org := restapi.GetOrganizations("10")
	b, err := json.Marshal(org)
	if err != nil {
		log.Warn("organization table json marshal error: %v", err.Error())
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}

