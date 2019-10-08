package main

import (
	"io"
	"log"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"backend/internal/api"
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

	// Authorization
	r.GET("/", api.StatusCheck)
	r.POST("/users", api.Users_SignUp)
	r.POST("/session", api.Session_SignIn)
	r.DELETE("/session", mw.Auth(api.Session_SignOut))

	// Personalization
	r.GET("/user/groups", mw.Auth(api.User_AllGroups))
	r.GET("/user/groups/:name_or_id", mw.Auth(api.User_AboutGroup))
	r.POST("/user/groups/:name_or_id/add", mw.Auth(api.User_AddGroup))
	r.DELETE("/user/groups/:name_or_id", mw.Auth(api.User_DeleteGroup))

	// Data
	r.GET("/groups", api.Groups_All)
	r.GET("/groups/:name_or_id", api.Groups_About)

	// Ingredients
	r.GET("/ingredients", api.Ingredients_All)
	r.GET("/ingredients/:name_or_id", api.Ingredients_About)
	r.GET("/ingredients/:name_or_id/groups/", api.Ingredients_GroupAll) // name or id of groups!

	// Excluded ingredients
	r.POST("/user/ingredients/:name_or_id", api.User_AddExcludedIngredient)
	r.DELETE("/user/ingredients/:name_or_id", api.User_DeleteExcludedIngredient)

	log.Println("Listening on http://localhost:8888...")
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.Fatal(fasthttp.ListenAndServe(":8888", r.Handler))
}
