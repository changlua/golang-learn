package main

import (
	iris "github.com/kataras/iris/v12"
)

func main()  {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
	})

	app.Run(iris.Addr(":8080"))
}
