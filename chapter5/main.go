
/*
上一个视频演示来response返回的数据类型，字符串，html，json，xml,blob等
本节视频演示静态文件服务
下面是我从chapter1拷贝的demo
*/
package main

// 引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
)

func main(){
	// 创建echo实例
	e:=echo.New()

	//创建一个helloworld的路由
	e.GET("/",func(c echo.Context)error{
		return c.String(http.StatusOK,"Hello World!")
	})

	// 上面路由我们不做修改，先放置在那，我们创建一个文件夹assets，然后添加一个html文件和一个css文件
	// 先看我们在response 节使用的File函数
	// 我们看下
	e.GET("/index",func(c echo.Context)error{
		return c.File("assets/index.html")
	})

	// 我们使用static函数，测试下,
	// 我们在看下效果，请求static路由的时候
	e.Static("/static","assets")

	// 我们可以看到使用Static函数，可以将文件夹，作为静态路径。
	// 在做前后端分离开发的时候，我们如果不使用nignx 或 apache的时候，我们可以使用static
	// 函数渲染我们的静态页面，js文件没有演示，与css文件没什么差别
	// 本次视频就到这，谢谢大家，喜欢的朋友关注点赞
	// https://github.com/jianjunjie/echoDemo github 地址，欢迎star 和提issue，谢谢


	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}
