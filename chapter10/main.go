/*
上个视频演示了echo框架的日志，三种方式设置日志
本次视频演示文件上传下载，下面的代码是我从chapter1中拷贝的demo
*/
package main

// 引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
	"os"
	"io"
)

func upload(c echo.Context)error{
	// 可以通过echo.Contxt实例的FormFile函数获取客户端上传的单个文件
	file,err:=c.FormFile("filename") //这里的filename要与前端对应上
	if err!=nil{
		return err
	}
	// 先打开文件源
	src,err:=file.Open()
	if err!=nil{
		return err
	}
	defer src.Close()

	// 下面创建保存路径文件
	// file.Filename 即上传文件的名字
	dst,err:=os.Create("upload/"+file.Filename)
	if err !=nil {
		return err
	}
	defer dst.Close()

	// 下面将源拷贝到目标文件
	if _,err=io.Copy(dst,src);err !=nil{

		return err
	}

	return c.String(http.StatusOK,"文件上传成功")

	// 我们创建upload文件夹，然后重启看下效果
	// 我们可以看到文件上传成功
}
func upload2(c echo.Context)error{
	// 与单个文件上传不一样的地址是读取文件
	form,err :=c.MultipartForm()
	if err!=nil{
		return err
	}

	// files要与前端一致
	files:= form.File["files"]

	//下面都是一样，只是在for循环中
	// 我们从upload中拷贝下
	for _,file:=range files {
		// 先打开文件源
		src,err:=file.Open()
		if err!=nil{
			return err
		}
		defer src.Close()

		// 下面创建保存路径文件
		// file.Filename 即上传文件的名字
		dst,err:=os.Create("upload2/"+file.Filename)
		if err !=nil {
			return err
		}
		defer dst.Close()

		// 下面将源拷贝到目标文件
		if _,err=io.Copy(dst,src);err !=nil{

			return err
		}
	}
	return c.String(http.StatusOK,"多文件上传成功。")

	// 我们创建文件夹upload2，重启服务看下效果
	// 上传文件路径不对，我们在看下
	// 我们可以看到文件上传成功
}

// 本节视频就到这，喜欢的朋友关注，点赞，谢谢大家

func main(){
	// 创建echo实例
	e:=echo.New()

	//文件下载我们在静态文件以及response中有演示，我们简单看下
	// 文件下载这些函数，我们前面都有提到，可以翻看之前的视频。
	// e.GET("/",func(c echo.Context)error{
	// 	return c.Attachment("acchemt.txt","123.txt")
	// })

	// 下面我们看下文件上传，我们先看上传单个文件
	// 我们先实现上传页面,我们重启看下页面
	e.GET("/",func(c echo.Context)error{
		return c.File("index.html")
	})
	// 报错是因为我们还没实现upload方法，接下来我们实现方法
	e.POST("/upload",upload)

	// 下面看先多文件上传。
	// 首先我们先实现一个html页面
	// 我们实现页面渲染路由
	e.GET("/files",func(c echo.Context)error{
		return c.File("index2.html")
	})
	// 我们实现upload2路由方法,upload2与upload 有很多相似方法
	// 我们部分拷贝
	e.POST("/upload2",upload2)

	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}