package userapi

import (
	"backend/internal/database/user"
	"backend/internal/models"

	"log"

	"github.com/valyala/fasthttp"
)

func HealthCheck(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.Method()) + (" ") + string(ctx.Path()))
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func SignUp(ctx *fasthttp.RequestCtx) {
	log.Println("Sign Up: " + string(ctx.Method()) + " " + string(ctx.Path()) + " " + string(ctx.PostBody()))
	u := &models.User{}
	u.UnmarshalJSON(ctx.PostBody())
	code, cookie, message := user.SignUp(u.Username, u.Password)

	log.Println("Sign Up cookie: " + cookie)
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
	log.Println("Sign In: " + string(ctx.Method()) + " " + string(ctx.Path()) + " " + string(ctx.PostBody()))
	u := &models.User{}
	u.UnmarshalJSON(ctx.PostBody())

	code, cookie, message := user.SignIn(u.Username, u.Password)

	m := &models.Msg{}
	m.Message = message
	mJSON, _ := m.MarshalJSON()

	cook := &fasthttp.Cookie{}
	cook.SetKey("session")
	cook.SetValue(cookie)

	log.Println("Sign In cookie: " + cookie)

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
	log.Println("Sign Out: " + string(ctx.Method()) + " " + string(ctx.Path()) + " " + string(ctx.PostBody()))

	code, message := user.SignOut(string(ctx.Request.Header.Cookie("session")))

	log.Println(code)
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
