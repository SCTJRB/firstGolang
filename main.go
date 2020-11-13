package main

import (
	_ "firstProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//fmt.Println(beego.AppConfig.String("mysqluser"))
	beego.Run()
}

