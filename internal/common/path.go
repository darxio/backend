package common

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

func NameOrID(ctx *fasthttp.RequestCtx) (string, int32) {
	slug := ctx.UserValue("name_or_id").(string)
	id, _ := strconv.ParseInt(slug, 10, 32)

	return slug, int32(id)
}
