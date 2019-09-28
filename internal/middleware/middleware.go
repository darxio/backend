package middleware

import (
	"log"

	"github.com/valyala/fasthttp"
)

func Auth(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		cookie := ctx.Request.Header.Cookie("session")

		log.Println("Sign Out Middleware cookie: " + string(cookie))
		if cookie == nil {
			ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.SetBody([]byte("{\"message\":\"User not authorized.\"}"))
			return
		}

		next(ctx)
	}
}
