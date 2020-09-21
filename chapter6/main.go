/*
上一个视频我们演示了静态文件，如何使用echo当作静态文件服务器
本次视频我们演示渲染html模版
下面的代码是我从chapter1中拷贝的demo代码
*/
package main

// 引入依赖包
import (
	"net/http"
	"html/template"
	"github.com/labstack/echo"
	"io"
)

//我们先引入html模版包
// 创建模版struct结构
type Tempalte struct{
	templates *template.Template
}

//我们实现Template结构的render方法
func (t *Tempalte)Render(w io.Writer,
	name string,data interface{},c echo.Context)error{
		// ExecuteTemplate 是go语言标准包html/template相关函数，有兴趣的朋友可以
		// 到go语言官网查看下
		return t.templates.ExecuteTemplate(w,name,data)
}

func main(){
	// 创建echo实例
	e:=echo.New()

	// 我们创建两个模版html文件
	// w我们创建模版实列
	t:=&Tempalte{
		templates:template.Must(template.ParseGlob("public/view/*.html")),
	}
	// 将实列赋值给echo的render
	e.Renderer=t

	//创建一个helloworld的路由
	e.GET("/",func(c echo.Context)error{
		// 我们使用 Render方法渲染模版
		// 此处的index，与前端页面的index，后面的内容渲染到{{.}}位置
		return c.Render(http.StatusOK,"index","this is content in index page")
	})

	// 我们在渲染hello的
	e.GET("/hello",func(c echo.Context)error{
		// 我们使用 Render方法渲染模版
		// 与上面一样，此处渲染的是hello的索引的
		return c.Render(http.StatusOK,"hello","hello page")
	})
	// 我们看下结果
	// 我们可以看到渲染的结果
	// 本节视频就到这，喜欢的朋友，欢迎关注 点赞，谢谢大家

	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}
