package api

import (
	"backend/internal/database/products"
	"backend/internal/models"

	"log"
	"strconv"

	"backend/internal/contents-analyzer/analyzer"

	"github.com/valyala/fasthttp"
)
/* 
func Product_All(ctx *fasthttp.RequestCtx) {
	log.Println("Product All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))

	products_ := make(models.ProductExtendedArr, 0, common.Limit)
	code, message := products.All(&products_)

	productsJSON, _ := products_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(productsJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}
}
 */
func Product_GetOneBarcode(ctx *fasthttp.RequestCtx) {
	log.Println("Product GetOneBarcode: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	barcode, _ := strconv.ParseInt(ctx.UserValue("barcode").(string), 10, 64)

	pExt := &models.ProductExtended{}
	pShr := &models.ProductShrinked{}
	shrinked := false
	code, message := products.GetOneBarcode(barcode, pExt, pShr, &shrinked)

	var pJSON []byte
	if shrinked == false {
		var err error
		pExt.Ingredients, err = analyzer.Analyze(pExt.Contents)
		if err != nil {
			log.Println("Product_GetOneBarcode.Analyze: " + err.Error())
		}
		pJSON, _ = pExt.MarshalJSON()
	} else {
		pJSON, _ = pShr.MarshalJSON()
	}

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(pJSON)
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

func Product_GetManyByName(ctx *fasthttp.RequestCtx) {
	log.Println("Product GetManyByName: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	name, _ := ctx.UserValue("name").(string)

	pExt := models.ProductExtendedArr{}
	pShr := models.ProductShrinkedArr{}
	shrinked := false
	code, message := products.GetManyByName(name, &pExt, &pShr, &shrinked)

	var pJSON []byte
	if shrinked == false {
		for _, v := range pExt {
			v.Ingredients, _ = analyzer.Analyze(v.Contents)
		}
		pJSON, _ = pExt.MarshalJSON()
	} else {
		pJSON, _ = pShr.MarshalJSON()
	}

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(pJSON)
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
