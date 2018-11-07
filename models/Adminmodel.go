package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql" 
)
 
 

type Beego_admin struct{
	Id int
	A_user string `form:"user"`
	A_pw string	`form:pw`
}

//链接数据库
func init() {
  
    //  default 是别名    beego_test 是数据库
    orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beego_test?charset=utf8", 30)
	 
    // register model
    orm.RegisterModel(new(Beego_admin)) 
    // create table
    orm.RunSyncdb("default", false, true)
	 
}
//添加管理员
func Add_admin(Admin Beego_admin)(int64){
	o:=orm.NewOrm()
	
	id,err:=o.Insert(&Admin)
	if err==nil{
		return id
	}else{
		return 0
	}
}
//登录操作
func Login_admin(Admin Beego_admin)(a Beego_admin,or bool){ 
	o:=orm.NewOrm()
	//user := Beego_admin{A_user: Admin.A_user}
	//err:=o.Read(&Admin,"A_user,A_pw")
	err := o.QueryTable("Beego_admin").Filter("A_user", Admin.A_user).Filter("A_pw", Admin.A_pw).One(&a)
	fmt.Println(err,Admin)
	if err==nil{
		return a,true
	}else{
		return a,false
	}
}
