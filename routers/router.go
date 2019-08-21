package routers

import (
    "github.com/astaxie/beego"
    "liteblog/controllers"
)

func init() {

    beego.ErrorController(&controllers.ErrorController{})
    //采用注解路由 需要引入
    beego.Include(&controllers.IndexController{})
}
