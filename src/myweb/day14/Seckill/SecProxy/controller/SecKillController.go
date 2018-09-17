package controller

import "github.com/astaxie/beego"

type SecKillController struct {
	beego.Controller
}

func (p *SecKillController) SecKill() {
	//
	p.Data["json"] = "SecKill"
	p.ServeJSON()
}

func (p *SecKillController) FindSecKill() {
	p.Data["json"] = "FindSecKill"
	p.ServeJSON()
}
