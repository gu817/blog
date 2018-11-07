package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"bengo_test/models"
	 
)

type LoginController struct {
	beego.Controller
}

 
//显示登录页面
func (this *LoginController) Login() { 
 
	this.TplName = "admin/login.html"
}
//登录
func(this *LoginController) Lg(){
	var Us models.Beego_admin
	Us.A_user=this.GetString("user")
	Us.A_pw=this.GetString("pw") 
	admin,or:=models.Login_admin(Us)
 	
	 
	if or{
		this.SetSession("user",admin.A_user)
		this.SetSession("u_id",admin.Id)
		 //this.Data["json"] = "{\"msg\":\"登录成功\",\"status\":\"0\"}"  
		 this.Data["json"] = &models.Jsonstruct{0,"登录成功"}  
	}else{
		this.Data["json"] = &models.Jsonstruct{1,"登录失败"}  
	}
	this.ServeJSON()
}
/*
func(this *LoginController) Lg(){
	var Us models.Beego_admin
	Us.A_user=this.GetString("user")
	Us.A_pw=this.GetString("pw")
	admin,or:=models.Login_admin(Us)
	if or{
		this.SetSession("user",admin.A_user)
		 
		this.Redirect("/admin/main",200) 
	}else{
		 if err := json.Unmarshal(this.Ctx.Input.RequestBody, &Us); err == nil {
        	admin,or:=models.Login_admin(Us)
        this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	    } else {
	        this.Data["json"] = err.Error()
	    }
	    this.ServeJSON()
	}
} 
*/
//退出登录

//添加
func (this *LoginController) Add() {
	var Us models.Beego_admin
	Us.A_user="admin"
	Us.A_pw="admin"
	or:=models.Add_admin(Us)
	if or>0{
		fmt.Println("")
	}
	this.Ctx.WriteString("添加成功")
} 

