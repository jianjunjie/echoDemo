/*
上个视频演示了渲染html模版，如何使用模版
本节视频演示cookie和session
下面的代码是我从chapter1中拷贝的demo
*/
package main

// 引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
	"time"
	"fmt"
	// 如果没有这个两个包，使用go get 下载或者别的包管理下载也许
	"github.com/gorilla/sessions"
  	"github.com/labstack/echo-contrib/session"
)

func WriteCookie(c echo.Context)error{
	// 创建cookie实列对象
	cookie:=new(http.Cookie)
	// 我们给cookie设置一些值
	// 我们改变name与value看下
	cookie.Name="nameKey"
	cookie.Value="ValueData"
	cookie.Expires=time.Now().Add(2*time.Hour) // cookie的过期时间
	// 还有一些cookie的属性设置，本处不做赋值，有兴趣的看先cookie的参数

	// 设置cookie到echo.Context中
	c.SetCookie(cookie)
	return c.String(http.StatusOK,"write a cookie")
	// 我们启动服务看下
}

// 读取一个cookie的handler
func ReadCookie(c echo.Context)error{
	cookie,err:=c.Cookie("nameKey")
	if err !=nil{
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK,"read a cookie")
}

// 读取所有cookie的handler函数
func ReadAllCookie(c echo.Context)error{
	for _,cookie:=range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK,"read all the cookie")
}

// session handler
func SessionHnadler(c echo.Context)error{
	sess,_:=session.Get("session",c)
	sess.Options=&sessions.Options{
		Path: "/",
		MaxAge: 86400*7,
		HttpOnly: true,
	}

	sess.Values["foo"]="bar"
	sess.Save(c.Request(),c.Response())
	return c.String(http.StatusOK,"session handler")
}

func main(){
	// 创建echo实例
	e:=echo.New()

	//创建一个helloworld的路由
	e.GET("/",func(c echo.Context)error{
		return c.String(http.StatusOK,"Hello World!")
	})

	// 我们先创建一个写cookie到客户端
	// 我们创建路由与handler
	// 我们创建handler
	e.GET("/writeCookie",WriteCookie)

	// 下面我们读取cookie的操作
	// 创建ReadCookie函数
	// 我们重启服务测试看下效果
	e.GET("readCookie",ReadCookie)

	// 下面我们看下读取所有的cookie
	// 创建ReadAllCookie
	// 重启服务测试下，可以看到都输出了，第一个是因为刚才我们请求了readCookie
	e.GET("allCookies",ReadAllCookie)

	// 我们接下来看下session
	// 首先需要引入两个包，我拷贝下
	// 使用session，我们需要先配置下
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// 创建路由handler
	// 创建handler,重启服务
	e.GET("/session",SessionHnadler)
	// 我们可以看到前端已经有session信息，session的读取与cookie差不多，本处不做演示
	// 本次视频就到这，谢谢大家，喜欢的朋友，关注，点赞，谢谢

	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}