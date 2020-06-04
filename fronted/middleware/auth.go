package middleware

import "github.com/kataras/iris"

func AuthConProduct(ctx iris.Context) {
	uid := ctx.GetCookie("uid")
	if uid == "0" {
		ctx.Application().Logger().Debug("必须先登陆")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug(uid)
	ctx.Application().Logger().Debug("已经登陆")
	ctx.Next()
}
