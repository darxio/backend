package userapi

import (
	"backend/internal/database/user"

	"backend/internal/models"

	"log"

	"github.com/valyala/fasthttp"
)

func SignUp(ctx *fasthttp.RequestCtx) {
	u := &models.User{}
	u.UnmarshalJSON(ctx.PostBody())

	log.Println(u)
	code, cookie, message := user.SignUp(u.Username, u.Password)

	m := &models.Msg{}
	m.Message = message
	mJSON, _ := m.MarshalJSON()

	cook := &fasthttp.Cookie{}
	cook.SetKey("session")
	cook.SetValue(cookie)

	log.Println(code)
	switch code {
	case fasthttp.StatusCreated:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusCreated)
		ctx.SetBody(mJSON)
	case fasthttp.StatusBadRequest:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(mJSON)
	case fasthttp.StatusConflict:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusConflict)
		ctx.SetBody(mJSON)
	case fasthttp.StatusInternalServerError:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(mJSON)
	}

	return
}

func SignIn(ctx *fasthttp.RequestCtx) {
	u := &models.User{}
	u.UnmarshalJSON(ctx.PostBody())

	code, cookie, message := user.SignIn(u.Username, u.Password)

	m := &models.Msg{}
	m.Message = message
	mJSON, _ := m.MarshalJSON()

	cook := &fasthttp.Cookie{}
	cook.SetKey("session")
	cook.SetValue(cookie)

	switch code {
	case fasthttp.StatusOK:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(mJSON)
	case fasthttp.StatusBadRequest:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(mJSON)
	case fasthttp.StatusNotFound:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusConflict)
		ctx.SetBody(mJSON)
	case fasthttp.StatusInternalServerError:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(mJSON)
	}

	return
}

func SignOut(ctx *fasthttp.RequestCtx) {
	code, message := user.SignOut(string(ctx.Request.Header.Cookie("session")))

	m := &models.Msg{}
	m.Message = message
	mJSON, _ := m.MarshalJSON()

	cook := &fasthttp.Cookie{}
	cook.SetKey("session")
	cook.SetValue("")

	switch code {
	case fasthttp.StatusOK:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(mJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(mJSON)
	}

	return
}
