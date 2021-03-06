package api

import (
	"backend/internal/common"
	"backend/internal/database/ingredients"
	"backend/internal/models"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

/*

func Ingredients_All(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))

	ingredients_ := make(models.IngredientArr, 0, common.Limit)
	code, message := ingredients.All(&ingredients_)

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

func Ingredients_All(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))

	ingredients_ := make(models.IngredientArr, 0, common.Limit)
	code, message := ingredients.All(&ingredients_)

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

func Ingredients_About(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients About: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	ingredientName, ingredientID := common.NameOrID(ctx)

	var code int
	var message string
	i := &models.Ingredient{}
	if ingredientID != 0 {
		code, message = ingredients.About("", ingredientID, i)
	} else {
		code, message = ingredients.About(ingredientName, 0, i)
	}

	iJSON, _ := i.MarshalJSON()

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

func Ingredients_Search_Paginated(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients Search: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	ingredientName, _ := common.NameOrID(ctx)
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
	code, message = ingredients.Search(ingredientName, count, offset, &ings)

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

func Ingredients_Search(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients Search: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	ingredientName, _ := common.NameOrID(ctx)
	count := 10
	page := 0

	offset := count * page
	var code int
	var message string
	ings := models.IngredientArr{}
	code, message = ingredients.Search(ingredientName, count, offset, &ings)

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

func Ingredients_Top(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients Top: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
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
	code, message = ingredients.Top(count, offset, &ings)

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

func Ingredients_GroupAll(ctx *fasthttp.RequestCtx) {
	log.Println("Ingredients About: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	groupName, groupID := common.NameOrID(ctx)

	ingredients_ := make(models.IngredientArr, 0, common.Limit)
	var code int
	var message string

	if groupID != 0 {
		code, message = ingredients.GroupAll("", groupID, &ingredients_)
	} else {
		code, message = ingredients.GroupAll(groupName, 0, &ingredients_)
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
