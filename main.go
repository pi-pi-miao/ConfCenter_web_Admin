package main

import (
	"github.com/astaxie/beego"
	"ConfCenter_web_Admin/initialization"
	_ "ConfCenter_web_Admin/route"
)

func main(){
	err := initialization.Init()
	if err != nil {
		return
	}
	beego.Run()
}
