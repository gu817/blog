package routers

import (
	"bengo_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
   // beego.Router("/", &controllers.MainController{})
	//beego.Router("/list", &controllers.MainController{},"*:List")
	//beego.Router("/add", &controllers.MainController{},"post:Add")

	//admin登录控制器
	beego.Router("/admin/login",&controllers.LoginController{},"get:Login")
	beego.Router("/admin/lg",&controllers.LoginController{},"post:Lg") 
	//添加admin
	beego.Router("/admin/add",&controllers.LoginController{},"get:Add")

	//后台 主界面
	beego.Router("/admin/index",&controllers.MainController{},"get:Main")	
	//显示系统信息
	beego.Router("/admin/welome",&controllers.MainController{},"get:Welcome")
 }


