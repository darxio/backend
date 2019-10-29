package api

/*
import (
	"backend/internal/common"
	"backend/internal/database/useringredients"
	"backend/internal/models"
	"log"

	"github.com/valyala/fasthttp"
)


func User_AllExcludedIngredient(ctx *fasthttp.RequestCtx) {
	log.Println("User AllExcludedIngredient: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	var message string

	ingredients_ := make(models.IngredientArr, 0, common.Limit)
	cookie := string(ctx.Request.Header.Cookie("session"))
	code, message := useringredients.AllExcludedIngredients(cookie, &ingredients_)

	ingredientsJSON, _ := ingredients_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(ingredientsJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}
}

func User_AddExcludedIngredient(ctx *fasthttp.RequestCtx) {
	log.Println("User AddExcludedIngredient: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	ingredientName, ingredientID := common.NameOrID(ctx)
	cookie := string(ctx.Request.Header.Cookie("session"))
	var code int
	var message string

	ingredients_ := make(models.IngredientArr, 0, common.Limit)

	if ingredientID != 0 {
		code, message = useringredients.AddExcludedIngredient(cookie, "", ingredientID, &ingredients_)
	} else {
		code, message = useringredients.AddExcludedIngredient(cookie, ingredientName, 0, &ingredients_)
	}

	ingredientsJSON, _ := ingredients_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(ingredientsJSON)
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

func User_DeleteExcludedIngredient(ctx *fasthttp.RequestCtx) {
	log.Println("User DeleteExcludedIngredient: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
		ingredientName, ingredientID := common.NameOrID(ctx)
		cookie := string(ctx.Request.Header.Cookie("session"))
		var code int
		var message string

		ingredients_ := make(models.IngredientArr, 0, common.Limit)

		if ingredientID != 0 {
			code, message = useringredients.DeleteExcludedIngredient(cookie, "", ingredientID, &ingredients_)
		} else {
			code, message = useringredients.DeleteExcludedIngredient(cookie, ingredientName, 0, &ingredients_)
		}

		ingredientsJSON, _ := ingredients_.MarshalJSON()

		switch code {
		case fasthttp.StatusOK:
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.SetBody(ingredientsJSON)
		case fasthttp.StatusInternalServerError:
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			m := &models.Msg{}
			m.Message = message
			mJSON, _ := m.MarshalJSON()
			ctx.SetBody(mJSON)
		}
}
*/
