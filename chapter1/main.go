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

	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}

// 现在我们使用命令行启动服务
// 服务已经启动，我们使用浏览器访问下，我们可以看到浏览器页面，输出Hello World！
// 本节就到这，谢谢大家