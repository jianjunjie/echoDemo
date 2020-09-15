/*
上一个视频演示了路由，本次视频将演示从request中获取请求参数
例如从formdata,query,path等，获取请求参数
*/

package main

// 引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
)

// Product 数据结构
// json tag将json请求参数绑定到product上
// form --formdata
// query-- query 参数
type Product struct{
	Name string `json:"name" form:"name" query:"name"`
	Price string `json:"price" form:"price" query:"price"`
}

// 绑定参数的handler
func MyHandler(c echo.Context) error{
	p:=new(Product)
	if err:=c.Bind(p); err!=nil {
		return c.String(http.StatusBadRequest,"bad param")
	}
	return c.JSON(http.StatusOK, p)
}

// 从formdata获取参数的handler
func MyHandler2(c echo.Context) error {
	name := c.FormValue("name") // 从formdata请求中获取数据
	return c.String(http.StatusOK, name)
}

// 从query获取参数的Handler
func MyHandler3(c echo.Context) error {
	name := c.QueryParam("name") // 从query请求中获取数据
	return c.String(http.StatusOK, name)
}

// 从path获取参数的Handler
func MyHandler4(c echo.Context) error {
	name := c.Param("name") // 从path请求中获取数据，直接使用c.Param函数获取参数
	return c.String(http.StatusOK, name)
}

// 主程序入口
func main(){

	// 创建实例
	e:=echo.New()

	// 先演示绑定的使用，使用echo.Context将请求参数根据tag绑定为Go类型的数据结构
	// 例如formdata请求，query参数，json请求等，会自动根据header请求类型自动判断
	// 定义数据结构，我们定义一个product数据结构
	// 实现一个handler函数MyHandler
	// 实现路由
	// 我们测试下
	// 这就是参数绑定，将请求参数，绑定为Go的数据类型
	e.GET("/products",MyHandler)

	// 我们也可以不使用绑定,直接从echo.Context中获取

	// formdata
	// 从formdata从获取参数
	// 实现handler MyHandler2
	// 实现路由，测试下
	e.GET("/products2",MyHandler2)

	// 在看下从query中获取参数
	// 测试下
	e.GET("/products3",MyHandler3)

	// 看下path请求参数，创建路由
	// 我们实现handler，测试下
	// 这就是从路径中获取参数
	e.GET("/products4/:name", MyHandler4)

	// echo的http请求参数，在设计restful风格的api时候，已经满足需求了
	// 请求数据的有效性判断，本视频不做演示，可以参考官方文档的实例
	// https://echo.labstack.com/guide/request
	// 结尾，介绍下仓库地址，本视频演示的代码仓库地址：
	// https://github.com/jianjunjie/echoDemo
	// 本次视频就到这，谢谢大家！喜欢的话素质二连，关注，点赞

	// 启动服务
	e.Logger.Fatal(e.Start(":1323"))
}