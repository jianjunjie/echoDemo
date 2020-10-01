/*
上个视频演示了如何上传文件与下载文件
本视频演示如何使用websocket，下面的代码是从chapter1中拷贝的demo
*/
package main

// 引入依赖包，我们引入websocket的包
import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
)

//声明upgrade
// 对于upgrader不是很熟悉的朋友，可以到mdn网站看下
// https://developer.mozilla.org/zh-CN/docs/Web/API/WebSockets_API/Writing_WebSocket_servers
// 有兴趣的朋友可以研究下
var (
	upgrader=websocket.Upgrader{}
)

func hello(c echo.Context)error{
	// 创建ws
	ws,err:=upgrader.Upgrade(c.Response(),c.Request(),nil)
	if err !=nil{
		return err
	}
	defer ws.Close()
	// 写个循环一直侦听客户端发送的消息
	// 同时往客户端发送消息
	for {
		err :=ws.WriteMessage(websocket.TextMessage,[]byte("hello client!"))
		if err != nil {
			c.Logger().Error(err)
		}
		// 读取
		_,msg,err:=ws.ReadMessage()
		if err !=nil {
			c.Logger().Error(err)
		}
		// 这里的msg是[]byte，我们转为string
		fmt.Println(string(msg))
	}
	// 服务端写完了，我们写前端
	// 本次视频就到这谢谢大家，喜欢的朋友关注，点赞，谢谢
}

func main(){
	// 创建echo实例
	e:=echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//创建今天页面服务路径
	e.Static("/","./public")
	// 我们创建一个hello的handler
	e.GET("/ws",hello)
	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}
