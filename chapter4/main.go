/*
上个视频延时了从request中获取请求参数，本次视频演示返回数据类型
下面是我提前实现的demo，从chapter1中拷贝的
*/
package main

// 引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
	"encoding/json"
)

// 从前面demo中拷贝的，只是使用了json和xml标签
type Product struct {
	Name  string `json:"name" xml:"name"`
  	Price string `json:"price" xml:"price"`
}

func main(){
	// 创建echo实例
	e:=echo.New()

	//创建一个helloworld的路由
	// 这个helloworld handler 使用String函数返回字符串，一会我们在启动服务器观看下
	e.GET("/string",func(c echo.Context)error{
		return c.String(http.StatusOK,"<h1>Hello World!</h1>")
	})

	// html
	// 下面我们看下返回简单的html，使用echo.Conext实例的HTML方法
	// 我们创建一个路由以及handler，然后与String方法对比下

	e.GET("/html",func(c echo.Context)error{
		// 我们给字符串加上html标签，看下效果如何
		// 我们启动下服务测试下
		return c.HTML(http.StatusOK,"<h1>Hello World!</h1>")
	})

	// 我们看到html生效了
	// 我们可以看到：
	// String 不会做任何修改字符串，原封不动展示，我们看到的结果也是没有变化的
	// HTML 会解析HTML标签
	// 我们给String函数加上html标签看下效果

	// json
	// 我们下面看下返回json，一共有三种方式
	// 我们先创建一个product的struct结构
	// 我们创建路由
	e.GET("/json",func(c echo.Context)error{
		p:=&Product{
			Name:"football",
			Price:"$120",
		}
		// 下面是返回json的函数
		return c.JSON(http.StatusOK,p)
	})

	// 还有json pretty和json blob，还有stream json 方法，我们也一起写出来，一会请求演示

	// JSONPretty
	e.GET("/prettyjson",func(c echo.Context)error{
		p:=&Product{
			Name:"football",
			Price:"$120",
		}
		return c.JSONPretty(http.StatusOK,p,"  ")
	})

	// stream json
	e.GET("/streamjson",func(c echo.Context)error{
		p:=&Product{
			Name:"football",
			Price:"$120",
		}
		// 设置header
		c.Response().Header().Set(echo.HeaderContentType,
		echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		// 我们引入json包
		return json.NewEncoder(c.Response()).Encode(p)
	})

	//blob
	e.GET("/josnblob",func(c echo.Context)error{
		p:=&Product{
			Name:"football",
			Price:"$120",
		}
		// 将结构体p json 编码成[]byte
		data,_:=json.Marshal(p)
		// 返回blob,函数变化来，我们使用空字符串填充下
		// 我演示下
		// 我们使用别的字符窜看下，函数修改，我还没研究，用法，应该是前缀，有兴趣可以到官网看下
		return c.JSONPBlob(http.StatusOK,"【】",data)
	})

	// xml与json类似，我们就演示简单的返回xml,streamxml,XMLBlob 等与json类似
	// 只是修改json修改为xml
	// 我们创建路由和handler
	e.GET("/xml",func(c echo.Context)error{
		p:=&Product{
			Name:"football",
			Price:"$120",
		}
		// 下面是返回xml的函数
		// 我们测试下,我们可以看到前端显示的是xml数据
		return c.XML(http.StatusOK,p)
	})

	// 返回还有file，attainmentfile，blob，等类型，我们在这就演示一个file，和blob
	// 我在当前路径下，放来一个public文件夹，里面放置来一张panda图片，我们看下
	// 我们使用file函数返回图片
	e.GET("/png",func(c echo.Context)error{
		return c.File("./public/index.html")
		// 我们看下效果,可以看到输出图片，我们添加一个html文件看下,我们可以看到效果
	})

	//blob的使用我们将官网的例子拷贝使用下
	e.GET("blob",func(c echo.Context)error{
		data := []byte(`0306703,0035866,NO_ACTION,06/19/2006
	  0086003,"0005866",UPDATED,06/19/2006`)
	return c.Blob(http.StatusOK, "text/csv", data)
	})

	// 还有NoConetent,我们也拷贝下
	e.GET("null",func(c echo.Context)error{
		return c.NoContent(http.StatusOK)
	})
	// 我们测试下null 和 blob
	// 我们可以看到效果
	// 本次视频就到这，谢谢大家，喜欢的朋友关注点赞

	// 开启服务
	e.Logger.Fatal(e.Start(":1323"))
}