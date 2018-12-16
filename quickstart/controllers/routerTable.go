package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/vishvananda/netlink"
	"log"
	"strings"
)

type RouterTableController struct {
	beego.Controller
}

type routeTable []staticRouter

type staticRouter struct {
	DstAddr string `json:"dstAddr"`
	DstMask string `json:"dstMask"`
	Gateway string `json:"gateway"`
	Port string `json:"port"`
}

const allAddress = "0.0.0.0"

func (c *RouterTableController) Get() {
	var staticRoutes routeTable

	link, err := netlink.LinkByName(ETH_PORT)
	if err != nil {
		fmt.Println(err)
	}
	routerTable, err := netlink.RouteList(link, netlink.FAMILY_V4)
	for _, v := range routerTable {
		staticRoutes = append(staticRoutes, parseNetlink(v, ETH_PORT))
	}

	b, err := json.Marshal(staticRoutes)
	if err != nil {
		log.Println("json marshal static route table error:", err)
	}

	fmt.Fprint(c.Ctx.ResponseWriter, string(b))
}

func parseNetlink(route netlink.Route, port string) staticRouter {
	var sRouter staticRouter

	if route.Dst == nil {
		sRouter.DstAddr = allAddress
		sRouter.DstMask = allAddress
	}else {
		ipAndMask := route.Dst.String()
		res := strings.Split(ipAndMask, "/")
		sRouter.DstAddr = res[0]
		sRouter.DstMask = res[1]
	}

	if route.Gw == nil {
		sRouter.Gateway = allAddress
	}else {
		sRouter.Gateway = route.Gw.String()
	}

	sRouter.Port = port

	return sRouter
}
