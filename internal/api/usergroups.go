package api

import (
	"backend/internal/common"
	usergroups "backend/internal/database/usergroups"
	"backend/internal/models"
	"log"

	"github.com/valyala/fasthttp"
)

func User_AllGroups(ctx *fasthttp.RequestCtx) {
	log.Println("User AllGroups: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	var message string

	groups_ := make(models.GroupArr, 0, common.Limit)
	cookie := string(ctx.Request.Header.Cookie("session"))
	code, message := usergroups.All(cookie, &groups_)

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
	groupName, groupID := common.NameOrID(ctx)
	cookie := string(ctx.Request.Header.Cookie("session"))

	var code int
	var message string
	g := &models.Group{}

	if groupID != 0 {
		code, message = usergroups.About(cookie, "", groupID, g)
	} else {
		code, message = usergroups.About(cookie, groupName, 0, g)
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
	groupName, groupID := common.NameOrID(ctx)
	cookie := string(ctx.Request.Header.Cookie("session"))
	var code int
	var message string

	groups_ := make(models.GroupArr, 0, common.Limit)

	if groupID != 0 {
		code, message = usergroups.Add(cookie, "", groupID, &groups_)
	} else {
		code, message = usergroups.Add(cookie, groupName, 0, &groups_)
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
	groupName, groupID := common.NameOrID(ctx)
	cookie := string(ctx.Request.Header.Cookie("session"))
	var code int
	var message string

	groups_ := make(models.GroupArr, 0, common.Limit)

	if groupID != 0 {
		code, message = usergroups.Delete(cookie, "", groupID, &groups_)
	} else {
		code, message = usergroups.Delete(cookie, groupName, 0, &groups_)
	}

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
