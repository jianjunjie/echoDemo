/*
上一个视频演示了一个hello world的demo
本次视频演示对应第二篇文章的内容--路由
通过本次视频，我们可以了解到echo框架的路由使用以及路由的匹配原则
*/

package main

//引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
)

// 先创建一个hello的handler
func Hellofunc(c echo.Context)error{
	return c.String(http.StatusOK,"hello world!")
}

// 主函数
func main(){
	// 创建echo实例
	e:=echo.New()

	// 创建一个跟路由，调用Hellofunc handler
	// 此处使用GET方法，我们知道HTTP请求有POST，PUT，DELETE等方法，我们可以看下提示
	// 我们可以看到echo框架都给我们实现了这些方法
	e.GET("/",Hellofunc)

	// 下面我们看下路由匹配的原则，别的web后端框架的，匹配相似
	// 匹配原则
	// 静态 大于 参数 大于 匹配路由
	// 下面举例演示下
	// 下面三个路由，我们把id,放到new的前面，我们匹配new路由时，
	// 看是调用第二，还是第三个

	// 匹配路由 后面跟上*
	e.GET("/product/1/price/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "product 1 price all")
	})
	// 参数路由 id可变
	e.GET("/product/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "product get by id")
	})
	//静态路由 没有变量
	e.GET("/product/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "new product")
	})

	// 我们可以看到匹配的原则

	//结果：
	/*
		优先级顺序
		/product/new
		/product/:id
		/product/1/price/*
	*/

	// 下面我们介绍子路由，我么先创建两个简单的handler
	// 创建子路由，使用e.Group方法
	// 这就是子路由的内容，开发中，可根据不同内容，实现分组接口管理
	r := e.Group("/api")
	// 请求路径 /api/user
	r.GET("/user", UserHandler) 
	// 请求路径 /api/score
	r.GET("/score", ScoreHandler)
	//重启服务看下
	
	// 对于POST，DELETE，PUT等方法，此处只写个列子
	// 现对应的方法，创建handler
	// 实现对应的路由
	// 使用浏览访问下GET的路径
	// 使用postman 测试下POST路径
	e.POST("/products", createProduct)
	e.GET("/products", findProduct)
	e.PUT("/products", updateProduct)
	e.DELETE("/products", deleteProduct)


	//下面我们启动服务看下页面输出内容
	// 启动服务
	e.Logger.Fatal(e.Start(":1323"))

	// 本次视频介绍了
	// 1.路由实现 2.路由的匹配原则 3.子路由 4.路由的HTTP请求方法
	// 本次视频就到这，谢谢大家
	
}

func UserHandler(c echo.Context)error{
	return c.String(http.StatusOK,"UserHandler")
}

func ScoreHandler(c echo.Context)error{
	return c.String(http.StatusOK,"ScoreHandler")
}


func createProduct(c echo.Context) error {
	return c.String(http.StatusOK,"createProduct")
}
func findProduct(c echo.Context) error {
	return c.String(http.StatusOK,"findProduct")
}
func updateProduct(c echo.Context) error {
	return c.String(http.StatusOK,"updateProduct")
}
func deleteProduct(c echo.Context) error {
	return c.String(http.StatusOK,"deleteProduct")
}