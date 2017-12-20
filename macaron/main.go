package main

//导入包
import (
	"log"
	"net/http"

	"github.com/go-macaron/macaron"
)

func main() {
	//创建实例
	m := macaron.Classic()
	//注册路由和处理器, func() 匿名函数
	m.Get("/", myhandler)
	m.Get("/*", myhandler)
	//注册路由多次,函数多次调用.

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func myhandler(ctx *macaron.Context) string {
	return "the request path is:" + ctx.Req.RequestURI
}
