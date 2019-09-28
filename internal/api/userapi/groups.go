package userapi

import (
	"log"

	"github.com/valyala/fasthttp"
)

func UserGroups(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.Method()) + (" ") + string(ctx.Path()))
}

func AddGroups(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.Method()) + (" ") + string(ctx.Path()))
}

func DeleteGroups(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.Method()) + (" ") + string(ctx.Path()))
}
