package route

import (
	"github.com/astaxie/beego"
	"ConfCenter_web_Admin/controller"
)

func init(){
	beego.Router("/list_service",&controller.Operation{},"*:GetOperations")   //查看接入网关服务
}
