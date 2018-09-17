package router

import (
	"github.com/astaxie/beego"
	"myweb/day14/Seckill/SecProxy/controller"
)

func init() {
	beego.Router("/SecKill", &controller.SecKillController{}, "*:SecKill")
	beego.Router("/FindSecKill", &controller.SecKillController{}, "*:FindSecKill")
}
