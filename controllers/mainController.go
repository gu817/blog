package controllers

import (
	_"fmt"
	 "github.com/astaxie/beego"
	_"bengo_test/models"
)
 
type MainController struct {
	beego.Controller
}

//显示后台主界面
func (this *MainController) Main() { 
	this.TplName = "admin/index.html"
}
//显示系统信息
func (this *MainController) Welcome(){
	this.TplName="admin/welcome.html"
}