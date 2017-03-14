package main

import (
	"github.com/astaxie/beego"
	"github.com/thoratou/cgi-jds/controllers"
)

func main() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/signin", &controllers.HomeController{}, "post:SignInQuery")
	beego.Router("/signup", &controllers.HomeController{}, "post:SignUpQuery")
	beego.Run()
}
