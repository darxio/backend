package main

import (
	"log"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"backend/internal/api/groupapi"
	"backend/internal/api/userapi"
	mw "backend/internal/middleware"
)

func main() {
	var logFile *os.File
	if _, err := os.Stat("./logs/log.txt"); err != nil {
		os.Mkdir("./logs", 0777)
		logFile, _ = os.Create("./logs/log.txt")
		logFile.Chmod(0777)
	} else {
		logFile, err = os.OpenFile("./logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
	}
	defer logFile.Close()

	r := fasthttprouter.New()

	r.GET("/", userapi.HealthCheck)
	r.POST("/users", userapi.SignUp)
	r.POST("/session", userapi.SignIn)
	r.DELETE("/session", mw.Auth(userapi.SignOut))

	r.GET("/me/groups", userapi.UserGroups)
	r.POST("/me/groups/:name_or_id", userapi.AddGroups)
	r.DELETE("/me/groups/:name_or_id", userapi.DeleteGroups)

	r.GET("/groups", groupapi.All)
	r.GET("/groups/:name_or_id", groupapi.About)

	log.Println("Listening on http://localhost:8888...")
	// log.SetOutput(logFile)
	log.Fatal(fasthttp.ListenAndServe(":8888", r.Handler))
}
