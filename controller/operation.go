package controller

import (
	"ConfCenter_web_Admin/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type Operation struct {
	beego.Controller
}

func (o *Operation) GetOperations() {
	operations := &service.OperationsResult{}
	operations.Get()
	logs.Debug("this is operation data ------- %v", operations.Result)
	o.Data["operation_list"] = operations.Result
	o.Layout = "layout/layout.html"
	o.TplName = "operation/get.html"
	return
}
