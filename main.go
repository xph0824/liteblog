package main

import (
	"github.com/astaxie/beego"
	// 运行main方法之前 调用models包下的文件的init方法。先初始化数据库
	_ "github.com/liteblog/models"
	_ "github.com/liteblog/routers"
	"strings"
)

func main() {
	initTemplate()
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
}