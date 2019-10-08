package api

import (
	"backend/internal/database/groups"
	"backend/internal/models"
	"log"

	"backend/internal/common"

	"github.com/valyala/fasthttp"
)

func Groups_All(ctx *fasthttp.RequestCtx) {
	log.Println("Groups All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))

	groups_ := make(models.GroupArr, 0, common.Limit)
	code, message := groups.All(&groups_)

	groupsJSON, _ := groups_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(groupsJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}
}

func Groups_About(ctx *fasthttp.RequestCtx) {
	log.Println("Groups About: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	groupName, groupID := common.NameOrID(ctx)

	var code int
	var message string
	g := &models.Group{}
	if groupID != 0 {
		code, message = groups.About("", groupID, g)
	} else {
		code, message = groups.About(groupName, 0, g)
	}

	gJSON, _ := g.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(gJSON)
	case fasthttp.StatusNotFound:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}
}
