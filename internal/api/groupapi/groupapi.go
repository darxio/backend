package groupapi

import (
	"backend/internal/database/group"
	"backend/internal/models"
	"log"

	"github.com/valyala/fasthttp"
)

func About(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.Method()) + (" ") + string(ctx.Path()))
	groupName := ctx.UserValue("name").(string)

	g := &models.Group{}
	code := group.About(groupName, g)
	gJSON, _ := g.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(gJSON)
	case fasthttp.StatusNotFound:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}
