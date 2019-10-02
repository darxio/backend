package api

import (
	"log"

	"github.com/valyala/fasthttp"
)

func Ingredients_All(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
}

func Ingredients_About(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients About: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
}

func Ingredients_GroupAll(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients About: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
}

func User_AddExcludedIngredient(ctx *fasthttp.RequestCtx) {
	log.Println("User AddExcludedIngredient: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
}

func User_DeleteExcludedIngredient(ctx *fasthttp.RequestCtx) {
	log.Println("User DeleteExcludedIngredient: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
}