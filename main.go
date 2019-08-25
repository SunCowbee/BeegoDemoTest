package main

import (
	_ "BeegoDemoTest/routers"
	"github.com/astaxie/beego"
	_ "BeegoDemoTest/models"
)

func main() {
	beego.Run()
}

