
/*
上节视频我们演示了如何使用cookie和session
本次视频我们演示如何使用JWT
下面的代码是我从chapter1 中拷贝的demo
*/
package main

// 引入依赖包
import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strconv"
)
const (
	jwtKey="secretKey"
)
type User struct{
	// 我们一会使用json的请求数据，tag设置成json的
	Username string `json:"username"`
	Password string `json:"password"`
}
type JwtCustomClaims struct{
	Name string `json:"name"`
	ID int `json:"id"`
	// 我们先引入jwt包
	jwt.StandardClaims
}

func Login(c echo.Context)error{
	// 在进一步开发时，我们先创建请求的用户数据结构
	// 绑定用户
	u:=new(User)
	if err:=c.Bind(u);err != nil{
		return c.JSON(http.StatusOK,echo.Map{
			"errcode":401,
			"errmsg":"request error",
		})
	}
	// 绑定成功后，我们判断用户名和密码，本处，我写死
	// 实际情况，自己需要把密码加密后与数据库对比，还有判断用户存在情况
	if "pass"==u.Password && "name"==u.Username{
		// 我们使用自定义的jwt配置
		// 我们声明一个数据结构
		// 我们创建JwtCustomClaims实列
		claims:=&JwtCustomClaims{
			Name:u.Username,
			ID:12,// 此处的ID可以理解为用户ID，需要查表操作，这里写死了
			StandardClaims:jwt.StandardClaims{
				ExpiresAt:time.Now().Add(time.Hour*12).Unix(),
			},
		}

		// 我们创建token
		token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

		//生产token，并编码,声明常量 jwtKey
		t,err:=token.SignedString([]byte(jwtKey))
		if err !=nil {
			return err
		}

		return c.JSON(http.StatusOK,echo.Map{
			"errcode":0,
			"errmsg":"ok",
			"token":t,
		})
	}else{
		return c.JSON(http.StatusOK,echo.Map{
			"errcode":-1,
			"errmsg":"error pass or error name",
		})
	}
}
func GetInfo(c echo.Context)error{
	name:=c.Get("name").(string)
	id:=c.Get("uid").(int)
	return c.String(http.StatusOK,"name: "+name+", id:"+strconv.Itoa(id))
}

func main(){
	// 创建echo实例
	e:=echo.New()

	//创建一个helloworld的路由
	e.GET("/",func(c echo.Context)error{
		return c.String(http.StatusOK,"Hello World!")
	})

	// 上面路由我们先放置不管
	// 我们创建一个login的路由
	//下面创建handler函数Login
	e.POST("/login",Login)

	//下面我们校验token，使用中间件
	// 使用jwt中间件
	// 如果使用同一个e下的路由，判断的时候会校验所有路径
	// 我们使用group创建一个子路由

	r:=e.Group("/api")

	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &JwtCustomClaims{},
		SigningKey:[]byte(jwtKey),
	}))

	// 解析jwt中间路由
	r.Use(func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context)error{
			user:=c.Get("user").(*jwt.Token)// 类型断言
			claims:=user.Claims.(*JwtCustomClaims)
			c.Set("name",claims.Name)
			c.Set("uid",claims.ID)
			return next(c)
		}
	})

	// 我们创建一个getInfo的路由
	r.GET("/getInfo",GetInfo)

	// 我们测试下 
	// 本次视频就到这，谢谢大家，源码的话到仓库获取，github地址：
	// https://github.com/jianjunjie/echoDemo
	// 喜欢的朋友关注，点赞，谢谢大家，再见

	// 开启服务
	// e.Logger.Fatal()函数是日志相关函数，后面会有介绍，本次只是实现一个简单的demo
	e.Logger.Fatal(e.Start(":1323"))
}