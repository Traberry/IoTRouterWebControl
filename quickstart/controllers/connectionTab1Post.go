package controllers

import (
	"github.com/astaxie/beego"
	"modifyfile"
	"netconfig"
	"os/exec"
	"time"
)

type Tab1Controller struct {
	beego.Controller
}

func (c *Tab1Controller) Post() {
	//set wireless name and password
	modifyfile.ChangeFileParameter(WIRELESS_AP_FN, "ssid", c.Input().Get("wirelessName"))
	modifyfile.ChangeFileParameter(WIRELESS_AP_FN, "wpa_passphrase", c.Input().Get("wirelessPasswd"))
	exec.Command("service",  "hostapd", "stop").Run()//root permission required
	time.Sleep(1000 * time.Millisecond)
	exec.Command("service",  "hostapd", "start").Run()

	//set new account and password
	if c.Input().Get("oldPassword") == beego.AppConfig.String("password") {
		beego.AppConfig.Set("username", c.Input().Get("newAccount"))
		beego.AppConfig.Set("password", c.Input().Get("newPassword"))
	}

	//firmware update
	netconfig.DeleteFile("/home/pi/test.zip")
	c.SaveToFile("newApp", "/home/pi/test.zip")
	netconfig.UnzipFile("/home/pi/test.zip", "/home/pi/")
	time.Sleep(900 * time.Millisecond)
	netconfig.Reboot()


	//set time zone
	/*timeZone := c.Input().Get("timeZone")
	switch timeZone {
	case "0":
		netconfig.SetTimeZone_1()
	case "12":
		netconfig.SetTimeZone_12()
	default:
		fmt.Println("this message should not print in theory")
	}*/

	c.Redirect("/connection", 301)
}

