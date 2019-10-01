package api

import (
	"backend/internal/database/groups"
	"backend/internal/models"
	"log"

	"github.com/valyala/fasthttp"
	"backend/internal/common"
)


func Groups_All(ctx *fasthttp.RequestCtx) {
	log.Println("Groups All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))

	groups_ := make(models.GroupArr, 0, common.Limit)
	code := groups.All(&groups_)

	groupsJSON, _ := groups_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(groupsJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}

func Groups_About(ctx *fasthttp.RequestCtx) {
	log.Println("Groups About: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	groupName, groupID := common.NameOrID(ctx)

	var code int
	g := &models.Group{}
	if groupID != 0 {
		code = groups.About("", groupID, g)
	} else {
		code = groups.About(groupName, 0, g)
	}

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
