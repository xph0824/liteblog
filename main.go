package main

import (
	// 运行main方法之前 调用models包下的文件的init方法。先初始化数据库
	_ "./models"
	_ "liteblog/routers"
	"github.com/astaxie/beego"
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