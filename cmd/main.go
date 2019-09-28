package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"backend/internal/api/userapi"
	mw "backend/internal/middleware"
)

func main() {
	r := fasthttprouter.New()

	r.GET("/", userapi.HealthCheck)
	r.POST("/users", userapi.SignUp)
	r.POST("/session", userapi.SignIn)
	r.DELETE("/session", mw.Auth(userapi.SignOut))

	log.Println("Listening on http://localhost:8888...")
	log.Fatal(fasthttp.ListenAndServe(":8888", r.Handler))
}
