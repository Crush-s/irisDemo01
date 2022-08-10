package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	app.Use(myMiddleware)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	app.Run(iris.Addr(":8787"))
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	// getHeader
	userAgent := ctx.GetHeader("User-Agent")
	if userAgent != "abv" {
		ctx.JSON("No authorized for ping")
		return
	}
	ctx.Next()
}
