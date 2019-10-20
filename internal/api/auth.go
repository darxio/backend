package api

import (
	"backend/internal/database/user"
	"backend/internal/models"

	"log"

	"github.com/valyala/fasthttp"
)

func StatusCheck(ctx *fasthttp.RequestCtx) {
	log.Println("Status Check: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func Users_SignUp(ctx *fasthttp.RequestCtx) {
	log.Println("Sign Up: " + string(ctx.Method()) + " " + string(ctx.Path()) + " " + string(ctx.PostBody()))
	u := models.User{}
	u.UnmarshalJSON(ctx.PostBody())
	code, cookie, message := user.SignUp(&u)

	cook := &fasthttp.Cookie{}
	cook.SetKey("session")
	cook.SetValue(cookie)

	switch code {
	case fasthttp.StatusCreated:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusCreated)
		uJSON, _ := u.MarshalJSON()
		ctx.SetBody(uJSON)
	case fasthttp.StatusBadRequest:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	case fasthttp.StatusConflict:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusConflict)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	case fasthttp.StatusInternalServerError:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}

	return
}

func Session_SignIn(ctx *fasthttp.RequestCtx) {
	log.Println("Sign In: " + string(ctx.Method()) + " " + string(ctx.Path()) + " " + string(ctx.PostBody()))
	u := models.User{}
	u.UnmarshalJSON(ctx.PostBody())

	code, cookie, message := user.SignIn(&u)

	cook := &fasthttp.Cookie{}
	cook.SetKey("session")
	cook.SetValue(cookie)

	switch code {
	case fasthttp.StatusOK:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusOK)
		uJSON, _ := u.MarshalJSON()
		ctx.SetBody(uJSON)
		ctx.SetBody(uJSON)
	case fasthttp.StatusBadRequest:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	case fasthttp.StatusNotFound:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	case fasthttp.StatusInternalServerError:
		ctx.Response.Header.SetCookie(cook)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}

	return
}

func Session_SignOut(ctx *fasthttp.RequestCtx) {
	log.Println("Sign Out: " + string(ctx.Method()) + " " + string(ctx.Path()) + " " + string(ctx.PostBody()))

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
