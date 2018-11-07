package main

import (
	_ "bengo_test/routers"
	"github.com/astaxie/beego"
	_"bengo_test/models"
	_"github.com/astaxie/beego/orm" //引入orm
	_"github.com/go-sql-driver/mysql" // import your used driver
	 
)

func main() {
	 
	//引入session
	beego.BConfig.WebConfig.Session.SessionOn = true
	 

	beego.Run()
}

