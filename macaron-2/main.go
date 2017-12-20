package main

import "github.com/go-macaron/macaron"
import "log"

func main() {
	m := macaron.Classic()
	/*   ctx.data 多处理数据传递
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Num"] = 1

	},
		func(ctx *macaron.Context) {
			ctx.Data["Num"] = ctx.Data["Num"].(int) + 1

		},
		func(ctx *macaron.Context) string {
			return fmt.Sprintf("Num: %d", ctx.Data["Num"])

		})
	*/
	//http://127.0.0.1:4000/?uid=1000 数字是合法的,其他为非法,返回0
	//m.Get("/", func(ctx *macaron.Context) string {
	//	return fmt.Sprintf("Uid: %d", ctx.QueryInt64("uid"))
	//	return ctx.RemoteAddr()
	//	})

	//m.Get("/", next1, next2, next3)
	//安全的cookie 全局密钥
	/*	m.SetDefaultCookieSecret("123123123")
		m.Get("/set", func(ctx *macaron.Context) {
			ctx.SetSecureCookie("user", "value")

		})
		m.Get("/get", func(ctx *macaron.Context) string {
			val, _ := ctx.GetSecureCookie("user")
			return val
		})
		/*
			m.Get("/set", func(ctx *macaron.Context) {
				ctx.SetCookie("user", "value")

			})
			m.Get("/get", func(ctx *macaron.Context) string {
				return ctx.GetCookie("user")

			})
	*/
	//	m.Get("/set", func(ctx *macaron.Context) {
	//		ctx.SetSuperSecureCookie("123123", "user", "test")
	//	})
	//	m.Get("/get", func(ctx *macaron.Context) string {
	//		val, _ := ctx.GetSuperSecureCookie("123123", "user")
	//		return val
	//	})
	m.Get("/", func(l *log.Logger) {
		l.Println("this is log")
	})
	m.Run()

}

/*
func next1(ctx *macaron.Context) {
	fmt.Println("处理器1进入")
	ctx.Next()
	fmt.Println("处理器1退出")
}
func next2(ctx *macaron.Context) {
	fmt.Println("处理器2进入")
	ctx.Next()
	fmt.Println("处理器2退出")
}
func next3(ctx *macaron.Context) {
	fmt.Println("处理器3进入")
	ctx.Next()
	fmt.Println("处理器3退出")
}
*/
