package main

import (
	_ "influxproject/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

