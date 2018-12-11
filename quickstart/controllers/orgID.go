package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gpmgo/gopm/modules/log"
	"restapi"
)

type OrgIDController struct {
	beego.Controller
}

func (c *OrgIDController) Get() {
	var orgIDs []string
	org := restapi.GetOrganizations("10")
	for _, v := range org.Result {
		orgIDs = append(orgIDs, v.ID)
	}
	b, err := json.Marshal(orgIDs)
	if err != nil {
		log.Warn("organization IDs json marshal error: %v", err.Error())
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}
