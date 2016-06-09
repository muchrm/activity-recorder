package routers

import (
	"github.com/astaxie/beego"
	"influxproject/controllers"
)

func init() {
	/*
	beego.Router("/addaccel", &controllers.AccelController{})
	beego.Router("/addgyro", &controllers.GyroController{})
	beego.Router("/addmagnet", &controllers.MagnetController{})
	beego.Router("/addheartrate", &controllers.HeartRateController{})
	*/
	beego.Router("/getacc", &controllers.GetAccController{})
	beego.Router("/addarray", &controllers.DataControls{})
}
