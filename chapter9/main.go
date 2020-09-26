/*
上个视频演示了echo框架如何使用jwt
本次视频演示echo的日志，下面的代码是从chapter1中拷贝的代码
*/
package main

// 引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/labstack/echo/middleware"
)

func main(){
	// 创建echo实例
	e:=echo.New()

	// 我们也可以使用中间件，echo会把每个请求输出一个日志，我们看下
	// 先引入中间件，添加代码
	//我们重启下，看下效果
	// middleware.Logger会把请求，返回的日志统一成json输出
	// 还是很方便的
	e.Use(middleware.Logger())


	// 我们开发的时候可以设置e.Debug
	// 我们可以看到日志输出的信息，我们在开发的时候，可以考虑使用e.debug
	e.Debug=true

	//创建一个helloworld的路由
	e.GET("/",func(c echo.Context)error{
		// 我们输出日志看下，Debug有三个输出模式，我们看下提示
		// 我们引入log
		// 我们可以使用默认的输出Debug函数，也可以格式化输出Debugf，同时也可以
		// 按照自定义的json输出
		// info error warn 都是类似的
		e.Logger.Debugf("这是格式化输出%s","ddddddddddddddd")
		e.Logger.Debugj(log.JSON{"aaaa":"cccccccccccccccccccc"})
		e.Logger.Debug("aaaaaaaaaaaaaaaaaaaaa")
		return c.String(http.StatusOK,"Hello World!")
	})

	// 我们如果不使用e.Debug的话，我们可以使用e.Logger.SetLevel来设置
	// 我们看下
	e.Logger.SetLevel(log.DEBUG)
	// e.Logger.SetOutput()
	e.GET("/info",func(c echo.Context)error{
		// 我们把上面的注释后，看下效果
		e.Logger.Infof("这是格式化输出%s","info ddddddddddddddd")
		e.Logger.Infoj(log.JSON{"info aaaa":"cccccccccccccccccccc"})
		e.Logger.Info("info aaaaaaaaaaaaaaaaaaaaa")
		return c.String(http.StatusOK,"INFO PAGE!")
	})


	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}

// 三种方式根据自己喜欢的方式使用
// 日志输出的本视频不做演示，可以看下函数e.Logger.SetOutput()
// 有兴趣的朋友可以自己试下，将日志输出到文件
// 本次视频就到这，谢谢大家，喜欢的朋友关注点赞，谢谢