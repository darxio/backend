package api

import (
	"backend/internal/common"
	"backend/internal/models"
	usergroups "backend/internal/database/usergroups"
	"log"

	"github.com/valyala/fasthttp"
)

func User_AllGroups(ctx *fasthttp.RequestCtx) {
	log.Println("User Groups: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	un := &models.Username{}
	un.UnmarshalJSON(ctx.PostBody())

	var message string

	groups_ := make(models.GroupArr, 0, common.Limit)
	code, message := usergroups.All(un.Username, &groups_)

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

func User_AboutGroup(ctx *fasthttp.RequestCtx) {
	log.Println("User AboutGroup: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	un := &models.Username{}
	un.UnmarshalJSON(ctx.PostBody())
	groupName, groupID := common.NameOrID(ctx)

	var code int
	var message string
	g := &models.Group{}

	if groupID != 0 {
		code, message = usergroups.About(un.Username, "", groupID, g)
	} else {
		code, message = usergroups.About(un.Username, groupName, 0, g)
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

func User_AddGroup(ctx *fasthttp.RequestCtx) {
	log.Println("User AddGroups: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	un := &models.Username{}
	un.UnmarshalJSON(ctx.PostBody())
	groupName, groupID := common.NameOrID(ctx)
	var code int
	var message string

	groups_ := make(models.GroupArr, 0, common.Limit)

	if groupID != 0 {
		code, message = usergroups.Add(un.Username, "", groupID, &groups_)
	} else {
		code, message = usergroups.Add(un.Username, groupName, 0, &groups_)
	}

	groupsJSON, _ := groups_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(groupsJSON)
	case fasthttp.StatusNotFound:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	case fasthttp.StatusConflict:
		ctx.SetStatusCode(fasthttp.StatusConflict)
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

func User_DeleteGroup(ctx *fasthttp.RequestCtx) {
	log.Println("User DeleteGroups: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	un := &models.Username{}
	un.UnmarshalJSON(ctx.PostBody())
	groupName, groupID := common.NameOrID(ctx)
	var code int
	var message string

	groups_ := make(models.GroupArr, 0, common.Limit)

	if groupID != 0 {
		code, message = usergroups.Delete(un.Username, "", groupID, &groups_)
	} else {
		code, message = usergroups.Delete(un.Username, groupName, 0, &groups_)
	}

	groupsJSON, _ := groups_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(groupsJSON)
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
