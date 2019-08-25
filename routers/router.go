package routers

import (
	"BeegoDemoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:ShowGet;post:Post")

    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandlePost")

    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
   /*
    //给请求指定自定义方法 一个请求指定一个方法r
    beego.Router("/login",&controllers.LoginController{},"get:ShowLogin;post:PostFunc")
    //给多个请求制定一个方法4
    beego.Router("/index",&controllers.IndexController{},"get,post:HandleFunc")
    //给所有请求指定一个方法
	beego.Router("/index",&controllers.IndexController{},"*:HandleFunc")
	//当两种指定方法冲突的时候
	beego.Router("/index",&controllers.IndexController{},"*:HandleFunc;post:PostFunc")
	*/


}
