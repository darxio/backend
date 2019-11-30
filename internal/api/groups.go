package api

import (
	"backend/internal/database/groups"
	"backend/internal/models"
	"log"
	"strconv"

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

func Groups_Ingredients(ctx *fasthttp.RequestCtx) {
	log.Println("Groups All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	groupID, _ := strconv.Atoi(ctx.UserValue("group_id").(string))
	count, _ := strconv.Atoi(ctx.UserValue("count").(string))
	page, _ := strconv.Atoi(ctx.UserValue("page").(string))

	if count > 50 {
		count = 50
	} else if count < 1 {
		count = 1
	}
	if page > 50 {
		page = 50
	} else if page < 0 {
		page = 0
	}

	offset := count * page
	var code int
	var message string
	ings := models.IngredientArr{}
	code, message = groups.Ingredients(groupID, count, offset, &ings)

	iJSON, _ := ings.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(iJSON)
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

func Groups_Ingredients_Search(ctx *fasthttp.RequestCtx) {
	log.Println("Groups All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	// groupID, _ := strconv.Atoi(ctx.UserValue("group_id").(string))
	groupID, _ := ctx.UserValue("group_id").(int)
	query := ctx.UserValue("query").(string)
	count, _ := strconv.Atoi(ctx.UserValue("count").(string))
	page, _ := strconv.Atoi(ctx.UserValue("page").(string))

	if count > 50 {
		count = 50
	} else if count < 1 {
		count = 1
	}
	if page > 50 {
		page = 50
	} else if page < 0 {
		page = 0
	}

	offset := count * page
	var code int
	var message string
	ings := models.IngredientArr{}
	code, message = groups.Search_Ing(groupID, query, count, offset, &ings)

	iJSON, _ := ings.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(iJSON)
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

func Groups_Search(ctx *fasthttp.RequestCtx) {
	log.Println("Groups Search: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	groupName, _ := common.NameOrID(ctx)

	var code int
	var message string
	grps := models.GroupArr{}
	code, message = groups.Search(groupName, &grps)

	iJSON, _ := grps.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(iJSON)
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
