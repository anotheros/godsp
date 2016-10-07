package main

import (
	"bid"

	"github.com/bsm/openrtb"
	"github.com/kataras/iris"
)

type myjson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
	Age   int32
}

func main() {
	iris.Config.Charset = "UTF-8"
	//var req *openrtb.BidRequest
	jsonHandlerSimple := func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, myjson{Name: "iris", Age: 3})
	}

	jsonHandlerWithRender := func(ctx *iris.Context) {
		// you can also change the charset for a specific render action with RenderOptions
		ctx.Render("application/json", myjson{Name: "iris"}, iris.RenderOptions{"charset": "utf-8"})
	}

	userHandler := func(ctx *iris.Context) {
		u := new(User)
		if err := ctx.ReadJSON(u); err != nil {
			ctx.EmitError(iris.StatusInternalServerError)
			return
		}
		u.Age = 18
		ctx.JSON(iris.StatusCreated, u)
	}

	bidController := func(ctx *iris.Context) {
		var req *openrtb.BidRequest
		req = new(openrtb.BidRequest)
		if err := ctx.ReadJSON(req); err != nil {
			ctx.EmitError(iris.StatusInternalServerError)
			return
		}
		//var resp openrtb.BidResponse
		resp := bid.BidHandler(*req)
		ctx.JSON(iris.StatusOK, resp)
	}

	iris.Post("/bid", bidController)

	iris.Post("/users", userHandler)

	iris.Get("/", jsonHandlerSimple)

	iris.Get("/render", jsonHandlerWithRender)

	iris.Listen(":8080")
}
